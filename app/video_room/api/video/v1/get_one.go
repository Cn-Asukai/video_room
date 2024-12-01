package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"video_room/app/video_room/internal/model/entity"
)

type GetOneReq struct {
	g.Meta `path:"/" tags:"Video" method:"get" summary:"get a video"`
	Id     int `json:"id"`
}

type GetOneRes struct {
	g.Meta `mine:"application/json"`
	Video  entity.Videos `json:"video"`
}
