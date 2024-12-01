// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Rooms is the golang structure for table rooms.
type Rooms struct {
	Id             int         `json:"id"             orm:"id"               description:""`          //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`          //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`          //
	Creator        int         `json:"creator"        orm:"creator"          description:"创建者id"`     // 创建者id
	Owner          int         `json:"owner"          orm:"owner"            description:"所有人id"`     // 所有人id
	CurrentVideo   int         `json:"currentVideo"   orm:"current_video"    description:"当前播放的视频id"` // 当前播放的视频id
	Name           string      `json:"name"           orm:"name"             description:"房间名"`       // 房间名
	UserCount      int         `json:"userCount"      orm:"user_count"       description:"当前房间人数"`    // 当前房间人数
	UserCountLimit int         `json:"userCountLimit" orm:"user_count_limit" description:"房间人数上限"`    // 房间人数上限
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:""`          //
}
