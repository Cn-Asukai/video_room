package video

import (
	"VideoRoom/app/video/api/video/v1"
	"VideoRoom/app/video/internal/model/do"
	"VideoRoom/app/video/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Video().Create(ctx, &do.Videos{
		Name:     req.Name,
		Url:      req.Url,
		CoverUrl: req.CoverUrl,
	})
	if err != nil {
		glog.Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	return &v1.CreateRes{}, nil
}
