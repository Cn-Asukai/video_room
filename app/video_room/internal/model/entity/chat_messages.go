// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatMessages is the golang structure for table chat_messages.
type ChatMessages struct {
	Id        int         `json:"id"        orm:"id"         description:""`      //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`      //
	Sender    int         `json:"sender"    orm:"sender"     description:"发展者id"` // 发展者id
	Receiver  int         `json:"receiver"  orm:"receiver"   description:"接收者id"` // 接收者id
	Seen      int         `json:"seen"      orm:"seen"       description:"是否已读"`  // 是否已读
}
