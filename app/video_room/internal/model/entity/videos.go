// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Videos is the golang structure for table videos.
type Videos struct {
	Id        int         `json:"id"        orm:"id"         description:""`             //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`             //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`             //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`             //
	Name      string      `json:"name"      orm:"name"       description:""`             //
	Url       string      `json:"url"       orm:"url"        description:""`             //
	CoverUrl  string      `json:"coverUrl"  orm:"cover_url"  description:""`             //
	State     int         `json:"state"     orm:"state"      description:"0:审核中 1:审核完成"` // 0:审核中 1:审核完成
}
