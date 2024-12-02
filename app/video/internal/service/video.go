// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"VideoRoom/app/video/internal/model/do"
	"VideoRoom/app/video/internal/model/entity"
	"context"
)

type (
	IVideo interface {
		Create(ctx context.Context, in *do.Videos) (err error)
		Delete(ctx context.Context, in *do.Videos) (err error)
		GetOne(ctx context.Context, in *do.Videos) (out *entity.Videos, err error)
		GetList(ctx context.Context, page int, limit int, in *do.Videos) (videos []entity.Videos, count int, err error)
	}
)

var (
	localVideo IVideo
)

func Video() IVideo {
	if localVideo == nil {
		panic("implement not found for interface IVideo, forgot register?")
	}
	return localVideo
}

func RegisterVideo(i IVideo) {
	localVideo = i
}
