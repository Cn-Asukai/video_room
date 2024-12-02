package video

import (
	"VideoRoom/app/video/internal/dao"
	"VideoRoom/app/video/internal/model/do"
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

func (v *sVideo) Create(ctx context.Context, in *do.Videos) (err error) {
	_, err = dao.Videos.Ctx(ctx).Insert(in)
	if err != nil {
		return err
	}

	return nil
}

func (v *sVideo) Delete(ctx context.Context, in *do.Videos) (err error) {
	_, err = dao.Videos.Ctx(ctx).Where("id", in.Id).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (v *sVideo) GetOne(ctx context.Context, in *do.Videos) (out *entity.Videos, err error) {
	out = &entity.Videos{}
	err = dao.Videos.Ctx(ctx).Where("id", in.Id).Scan(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (v *sVideo) GetList(ctx context.Context, page int, limit int, in *do.Videos) (videos []entity.Videos, count int, err error) {
	videos = make([]entity.Videos, 0)
	err = dao.Videos.Ctx(ctx).Where(in).
		OrderDesc("id").
		Limit(page, limit).
		ScanAndCount(&videos, &count, false)

	if err != nil {
		return nil, 0, err
	}
	return videos, count, nil
}
