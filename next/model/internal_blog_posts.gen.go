// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInternalBlogPost = "internal_blog_posts"

// InternalBlogPost mapped from table <internal_blog_posts>
type InternalBlogPost struct {
	ID          string    `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Slug        string    `gorm:"column:slug;not null" json:"slug"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Summary     string    `gorm:"column:summary;not null" json:"summary"`
	Content     string    `gorm:"column:content;not null" json:"content"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	IsFeatured  bool      `gorm:"column:is_featured;not null" json:"is_featured"`
	Status      string    `gorm:"column:status;not null;default:draft" json:"status"`
	CoverImage  string    `gorm:"column:cover_image" json:"cover_image"`
	SeoData     string    `gorm:"column:seo_data" json:"seo_data"`
	JSONContent string    `gorm:"column:json_content;not null;default:{}" json:"json_content"`
}

// TableName InternalBlogPost's table name
func (*InternalBlogPost) TableName() string {
	return TableNameInternalBlogPost
}