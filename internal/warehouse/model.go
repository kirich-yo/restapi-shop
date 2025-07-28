package warehouse

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	Address string `json:"address"`
}
