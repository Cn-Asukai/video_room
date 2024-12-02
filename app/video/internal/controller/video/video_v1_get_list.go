package video

import (
	"VideoRoom/app/video/api/video/v1"
	"VideoRoom/app/video/internal/model/do"
	"VideoRoom/app/video/internal/service"
	"context"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	videos, total, err := service.Video().GetList(ctx, req.Page, req.Size, &do.Videos{})

	return &v1.GetListRes{
		List:  videos,
		Total: total,
	}, nil
}
