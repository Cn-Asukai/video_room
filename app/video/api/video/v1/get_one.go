package v1

import (
	"VideoRoom/app/video/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetOneReq struct {
	g.Meta `path:"/" tags:"Video" method:"get" summary:"get a video"`
	Id     int `json:"id"`
}

type GetOneRes struct {
	g.Meta `mine:"application/json"`
	*entity.Videos
}
