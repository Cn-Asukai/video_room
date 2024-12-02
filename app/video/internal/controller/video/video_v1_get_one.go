package video

import (
	"VideoRoom/app/video/internal/model/do"
	"VideoRoom/app/video/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"VideoRoom/app/video/api/video/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	out, err := service.Video().GetOne(ctx, &do.Videos{Id: req.Id})
	if err != nil {
		glog.Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	res = &v1.GetOneRes{
		Videos: out,
	}
	return res, nil
}
