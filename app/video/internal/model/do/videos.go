// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Videos is the golang structure of table videos for DAO operations like Where/Data.
type Videos struct {
	g.Meta    `orm:"table:videos, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Name      interface{} //
	Url       interface{} //
	CoverUrl  interface{} //
	State     interface{} //
}