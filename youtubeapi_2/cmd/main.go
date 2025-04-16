package main

import (
	"fmt"
	"log"
	"os"
	"strconv" // strconvを追加

	"github.com/joho/godotenv"
	"github.com/omikuu/sr/domain/output"
	"github.com/omikuu/sr/infrastructure/reoisitory_imple"
	video_usecase "github.com/omikuu/sr/usecase/video"
)

// 日本人による直近3日以内のApex Legendsに関しての動画を
// 人気top10を出力するスクリプトを作成してください。
func main() {
	apiKey, title, limit := loadEnvVariables()
	yt, err := reoisitory_imple.NewYouTubeClient(apiKey)
	if err != nil {
		log.Fatalf("YouTube client error: %v", err)
	}

	usecase := video_usecase.FetchVideosUseCase{Repo: yt}

	videos, err := usecase.Execute(title, limit)
	if err != nil {
		log.Fatalf("Failed to fetch videos: %v", err)
	}

	fmt.Println(videos)
	// 再生回数順にソート
	// トップ10出力

}

func loadEnvVariables() (string, string, int) {
	_ = godotenv.Load()
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY not found")
	}

	title := os.Getenv("YOUTUBE_SEARCH_TITLE")
	if title == "" {
		log.Fatal("YOUTUBE_SERCH_TITLE not found")
	}

	limitStr := os.Getenv("YOUTUBE_SEARCH_LIMIT")
	if limitStr == "" {
		log.Fatal("YOUTUBE_SEARCH_LIMIT not found")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Fatalf("Invalid YOUTUBE_SEARCH_LIMIT: %v", err)
	}

	return apiKey, title, limit
}

func printVideos(videos []output.Video) {
	fmt.Println("🎬 最新動画リスト:")
	for _, v := range videos {
		fmt.Printf("- %s\n", v.URL)
	}
}
