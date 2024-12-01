package video

import (
	"context"
	"video_room/app/video_room/internal/model/entity"
	"video_room/app/video_room/internal/service"

	"video_room/app/video_room/api/video/v1"
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
	return &v1.CreateRes{}, nil
}
