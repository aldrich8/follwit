package models

import (
	"net/http"
	"time"
)

// Boilerplate is a model for boilerplate table
type Boilerplate struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	CreateAt    *time.Time `json:"create_at,omitempty,string"`
	UpdateAt    *time.Time `json:"update_at,omitempty,string"`
}

func (b Boilerplate) Bind(r *http.Request) error {
	//TODO implement me
	return nil
}

// TableName is Database TableName of this model
func (Boilerplate) TableName() string {
	return "boilerplate"
}
