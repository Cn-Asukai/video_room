// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"VideoRoom/app/video/internal/model/entity"
	"context"
)

type (
	IVideo interface {
		Create(ctx context.Context, in *entity.Videos) (err error)
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
