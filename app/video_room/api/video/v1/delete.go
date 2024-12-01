package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteReq struct {
	g.Meta `path:"/" method:"Delete"`
	Id     int `json:"id"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
}
