package models

type Favorite struct {
	ID          uint `json:"id" form:"id" gorm:"primarykey"`
	UserID      uint `json:"user_id" form:"user_id"`
	WarehouseID uint `json:"warehouse_id" form:"warehouse_id"`
	Warehouse   *Warehouse
	User        *User `json:"-"`
}
