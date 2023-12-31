// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameContentCollection = "content_collections"

// ContentCollection mapped from table <content_collections>
type ContentCollection struct {
	Type           string         `gorm:"column:type;primaryKey;<-:create" json:"type"`
	Source         string         `gorm:"column:source;primaryKey" json:"source"`
	ID             string         `gorm:"column:id;primaryKey;<-:create" json:"id"`
	Name           string         `gorm:"column:name;not null" json:"name"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null;<-:create" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null" json:"updatedAt"`
	MetadataSource MetadataSource `gorm:"foreignKey:Source" json:"metadata_source"`
}

// TableName ContentCollection's table name
func (*ContentCollection) TableName() string {
	return TableNameContentCollection
}
