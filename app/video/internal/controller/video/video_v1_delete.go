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

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Video().Delete(ctx, &do.Videos{
		Id: req.Id,
	})
	if err != nil {
		glog.Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	return &v1.DeleteRes{}, nil
}
