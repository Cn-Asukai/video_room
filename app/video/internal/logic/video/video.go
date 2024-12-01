package video

import (
	"VideoRoom/app/video/internal/dao"
	"VideoRoom/app/video/internal/model/entity"
	"VideoRoom/app/video/internal/service"
	"context"
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
