// Youtube の情報を取得するパッケージ
package yotube_resolver_impl

import (
	"context"
	"fmt"

	"github.com/nds-masashi/sr/domain/model/video"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type ServiceImpl struct {
	service *youtube.Service
}

func New(apiKey string) (*ServiceImpl, error) {
	ctx := context.Background()
	srv, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &ServiceImpl{service: srv}, nil
}

func (c *ServiceImpl) GetVideoUrl(title string, limit int) ([]video.Model, error) {
	call := c.service.Search.List([]string{"id", "snippet"}).
		Q(title).
		Type("video").
		Order("date").
		MaxResults(int64(limit))

	res, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videos []video.Model
	for _, item := range res.Items {
		videos = append(videos, video.New(
			fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId),
			item.Snippet.Title,
		))
	}
	return videos, nil
}
