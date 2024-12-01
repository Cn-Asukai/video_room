// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoComments is the golang structure of table video_comments for DAO operations like Where/Data.
type VideoComments struct {
	g.Meta      `orm:"table:video_comments, do:true"`
	Id          interface{} //
	VideoId     interface{} //
	Content     interface{} //
	Commentator interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}
