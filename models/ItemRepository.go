// models/ItemRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemRepository struct {
	Db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{Db: db}
}

func (r *ItemRepository) GetItems(c *gin.Context) {
	var items []Item
	r.Db.Find(&items)
	c.JSON(200, items)
}

func (r *ItemRepository) PostItem(c *gin.Context) {
	var newItem Item
	c.BindJSON(&newItem)
	r.Db.Create(&newItem)
	c.JSON(200, newItem)
}

