// Youtube のVideo情報を検索するパッケージ
package get_video_usecase

import (
	"github.com/nds-masashi/sr/domain/model/video"
	"github.com/nds-masashi/sr/domain/service/yotube_resolver"
)

type Usecase struct {
	Service yotube_resolver.Service
}

func (u *Usecase) Execute(title string, limit int) ([]video.Model, error) {
	return u.Service.GetVideoUrl(title, limit)
}
