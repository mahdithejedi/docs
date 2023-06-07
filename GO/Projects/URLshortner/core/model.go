package core

import (
	"time"

	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	URL       string
	Hash      string    `gorm:index:Hash_index,unique`
	CreatedAt time.Time `gorm:"<-create"`
}
