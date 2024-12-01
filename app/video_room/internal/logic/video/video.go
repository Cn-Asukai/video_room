package video

import (
	"context"
	"video_room/app/video_room/internal/dao"
	"video_room/app/video_room/internal/model/entity"
	"video_room/app/video_room/internal/service"
)

type sVideo struct {
}

func init() {
	service.RegisterVideo(New())
}

func New() *sVideo {
	return &sVideo{}
}

func (v *sVideo) Create(ctx context.Context, in *entity.Videos) (err error) {
	_, err = dao.Videos.Ctx(ctx).Insert(in)
	if err != nil {
		return err
	}
	return nil
}
