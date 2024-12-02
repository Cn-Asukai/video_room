package v1

import (
	"VideoRoom/app/video/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/list" method:"get" tags:"Video"`
	Page   int `json:"page"`
	Size   int `json:"size"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	List   []entity.Videos `json:"list"`
	Total  int             `json:"total"`
}
