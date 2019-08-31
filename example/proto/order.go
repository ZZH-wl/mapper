package order

import (
	"time"
)

type OldOrder struct {
	Id          int64     `json:"id" xorm:"pk autoincr"`
	CreatedOn   time.Time `json:"createdOn" xorm:"created"`
	ModifiedOn  time.Time `json:"modifiedOn" xorm:"updated"`
	DeletedOn   time.Time `json:"deletedOn" xorm:"deleted"`
	StoreId     int64     `json:"storeId" xorm:"index"`
	UserUid     string    `json:"userUid" xorm:"index" `
	OrderNo     string    `json:"orderNo"   xorm:"notnull VARCHAR(60) unique"`
	PayNo       string    `json:"payNo"   xorm:"VARCHAR(60)"`
	LogisticsNo string    `json:"logisticsNo" xorm:"VARCHAR(60)" `
	OrderStatus int64     `json:"orderStatus"   xorm:"default(1)" `
	PayTime     time.Time `json:"payTime" `
	OriginFee   float64   `json:"originFee" xorm:"default(0) NUMERIC(19,5)"`
	PayFee      float64   `json:"payFee" xorm:"notnull NUMERIC(19,5)" `
	UserMem     string    `json:"userMem" `
	IsStore     bool      `json:"isStore" xorm:"notnull default(true)"`
}
