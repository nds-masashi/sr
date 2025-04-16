package video

import (
	"github.com/omikuu/sr/domain/video_info"
)

type Repository interface {
	GetVideoUrl(title string, limit int) ([]video_info.VideoInfo, error)
}
