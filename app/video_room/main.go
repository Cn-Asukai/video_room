package main

import (
	_ "video_room/app/video_room/internal/packed"

	_ "video_room/app/video_room/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"video_room/app/video_room/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
