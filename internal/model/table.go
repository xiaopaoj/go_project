package model

import "time"

type TTable struct {
	ID   		uint64
	Name 		string
	CreateTime 	time.Time
	Utime 	string
}

func (TTable) TableName() string {
	return "t_table"
}