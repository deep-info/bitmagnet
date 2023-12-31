package crawler

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/anacrolix/dht/v2/krpc"
	"github.com/bitmagnet-io/bitmagnet/internal/dht/routing"
	"github.com/bitmagnet-io/bitmagnet/internal/dht/staging"
	"time"
)

func (c *crawler) crawlPeersForInfoHashes(ctx context.Context) {
	fullLogged := false
	for {
		if c.staging.Count() < c.maxStagingSize {
			fullLogged = false
			go (func() {
				err := c.routingTable.TryEachPeer(ctx, c.sampleInfoHashesFromLockedPeer)
				if err != nil {
					c.logger.Debugw("error crawling peers for info hashes", "err", err)
				}
			})()
		} else if !fullLogged {
			c.logger.Debug("staging is full, not crawling peers for info hashes")
			fullLogged = true
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(c.sampleInfoHashesInterval):
			continue
		}
	}
}

func (c *crawler) sampleInfoHashesFromLockedPeer(ctx context.Context, peer routing.PeerInfo) error {
	t := [20]byte{}
	if _, randErr := rand.Read(t[:]); randErr != nil {
		return fmt.Errorf("could not generate random bytes: %w", randErr)
	}
	res, resErr := c.dhtServer.Query(ctx, peer.Addr(), "sample_infohashes", krpc.MsgArgs{
		ID:     c.peerID,
		Target: t,
	})
	if resErr != nil {
		return resErr
	}
	if res.Reply.R == nil {
		return fmt.Errorf("sample_infohashes nil ret: %v", res.Reply)
	}
	peersToAdd := make([]krpc.NodeAddr, 0, len(res.Reply.R.Nodes))
	for _, n := range res.Reply.R.Nodes {
		peersToAdd = append(peersToAdd, n.Addr)
	}
	c.routingTable.ReceivePeers(peersToAdd...)
	if res.Reply.R.Samples != nil {
		hashesToStage := make([]staging.InfoHashWithPeer, 0, len(*res.Reply.R.Samples))
		for _, s := range *res.Reply.R.Samples {
			hashesToStage = append(hashesToStage, staging.InfoHashWithPeer{
				InfoHash: s,
				Peer:     peer.Addr(),
			})
		}
		c.staging.Stage(hashesToStage...)
		c.logger.Debugw("staged hashes", "nHashes", len(hashesToStage))
	}
	return nil
}
