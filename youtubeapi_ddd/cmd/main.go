package main

import (
	"fmt"
	"log"
	"os"
	"strconv" // strconvを追加

	"github.com/joho/godotenv"
	"github.com/nds-masashi/sr/domain/model/video"
	"github.com/nds-masashi/sr/infrastructure/yotube_resolver_impl"
	video_usecase "github.com/nds-masashi/sr/usecase/get_video_usecase"
)

func main() {
	apiKey, title, limit := loadEnvVariables()
	yt, err := yotube_resolver_impl.New(apiKey)
	if err != nil {
		log.Fatalf("YouTube client error: %v", err)
	}

	usecase := video_usecase.Usecase{Service: yt}

	videos, err := usecase.Execute(title, limit)
	if err != nil {
		log.Fatalf("Failed to fetch videos: %v", err)
	}

	printVideos(videos)
}

func loadEnvVariables() (string, string, int) {
	_ = godotenv.Load()
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY not found")
	}

	title := os.Getenv("YOUTUBE_SEARCH_TITLE")
	if title == "" {
		fmt.Println("検索タイトルを入力してください:")
		fmt.Scan(&title)
		if title == "" {
			log.Fatal("YOUTUBE_SERCH_TITLE not found")
		}
	}

	limitStr := os.Getenv("YOUTUBE_SEARCH_LIMIT")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Fatalf("Invalid YOUTUBE_SEARCH_LIMIT: %v", err)
	}

	return apiKey, title, limit
}

func printVideos(videos []video.Model) {
	fmt.Println("🎬 最新動画リスト:")
	for _, v := range videos {
		fmt.Printf("- %s [%s]\n", v.GetURL(), v.GetTitle())
	}
}
