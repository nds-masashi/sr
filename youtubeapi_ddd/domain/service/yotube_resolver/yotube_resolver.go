// Youtube の情報を取得するInterfaceパッケージ
package yotube_resolver

import "github.com/nds-masashi/sr/domain/model/video"

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./$GOPACKAGE.mock.go

type Service interface {
	GetVideoUrl(title string, limit int) ([]video.Model, error)
}
