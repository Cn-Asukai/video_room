// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoComments is the golang structure for table video_comments.
type VideoComments struct {
	Id          int         `json:"id"          orm:"id"          description:""`      //
	VideoId     int         `json:"videoId"     orm:"video_id"    description:""`      //
	Content     string      `json:"content"     orm:"content"     description:""`      //
	Commentator int         `json:"commentator" orm:"commentator" description:"评论者id"` // 评论者id
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`      //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`      //
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:""`      //
}