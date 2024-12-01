// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Rooms is the golang structure of table rooms for DAO operations like Where/Data.
type Rooms struct {
	g.Meta         `orm:"table:rooms, do:true"`
	Id             interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	Creator        interface{} // 创建者id
	Owner          interface{} // 所有人id
	CurrentVideo   interface{} // 当前播放的视频id
	Name           interface{} // 房间名
	UserCount      interface{} // 当前房间人数
	UserCountLimit interface{} // 房间人数上限
	DeletedAt      *gtime.Time //
}
