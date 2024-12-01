package video

import (
	"VideoRoom/app/video/internal/model/entity"
	"VideoRoom/app/video/internal/service"
	"context"

	"VideoRoom/app/video/api/video/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Video().Create(ctx, &entity.Videos{
		Id:        0,
		CreatedAt: nil,
		UpdatedAt: nil,
		DeletedAt: nil,
		Name:      req.Name,
		Url:       req.Url,
		CoverUrl:  req.CoverUrl,
		State:     0,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
