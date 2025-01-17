// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserProfile = "user_profiles"

// UserProfile mapped from table <user_profiles>
type UserProfile struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	FullName  string    `gorm:"column:full_name" json:"full_name"`
	AvatarURL string    `gorm:"column:avatar_url" json:"avatar_url"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
}

// TableName UserProfile's table name
func (*UserProfile) TableName() string {
	return TableNameUserProfile
}
