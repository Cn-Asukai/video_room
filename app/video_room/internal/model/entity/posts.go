// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure for table posts.
type Posts struct {
	Id        int         `json:"id"        orm:"id"         description:""`      //
	Publisher int         `json:"publisher" orm:"publisher"  description:"发布者id"` // 发布者id
	Context   string      `json:"context"   orm:"context"    description:""`      //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`      //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`      //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`      //
}
