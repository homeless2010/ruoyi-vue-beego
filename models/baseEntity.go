package models

import "time"

type BaseEntity struct {
	CreateBy   string    `orm:"column(create_by)"`
	CreateTime time.Time `orm:"column(create_time)"`
	UpdateBy   string    `orm:"column(update_by);auto_now_add"`
	UpdateTime time.Time `orm:"column(update_time)"`
}
