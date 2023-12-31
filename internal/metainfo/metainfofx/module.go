package metainfofx

import (
	"github.com/bitmagnet-io/bitmagnet/internal/boilerplate/config/configfx"
	"github.com/bitmagnet-io/bitmagnet/internal/metainfo/metainforequester"
	"go.uber.org/fx"
)

func New() fx.Option {
	return fx.Module(
		"metainfo",
		configfx.NewConfigModule[metainforequester.Config]("metainfo_requester", metainforequester.NewDefaultConfig()),
		fx.Provide(
			metainforequester.New,
		),
	)
}
