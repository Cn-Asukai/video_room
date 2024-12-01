// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// VideoTag is the golang structure for table video_tag.
type VideoTag struct {
	Id      int `json:"id"      orm:"id"       description:""` //
	VideoId int `json:"videoId" orm:"video_id" description:""` //
	TagId   int `json:"tagId"   orm:"tag_id"   description:""` //
}
