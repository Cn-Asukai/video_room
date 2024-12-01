package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta   `path:"/" tags:"Video" method:"Post" summary:"create a new video"`
	Name     string `json:"name" v:"required"`
	Url      string `json:"url" v:"required"`
	CoverUrl string `json:"cover_url" v:"required"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
}
