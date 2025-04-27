// Youtube の情報を取得するInterfaceパッケージ
package yotube_resolver

import "github.com/nds-masashi/sr/domain/model/video"

//go:generate mockgen -source=yotube_resolver.go -destination=yotube_resolver_mock.go -package=yotube_resolver

type Service interface {
	GetVideoUrl(title string, limit int) ([]video.Model, error)
}
