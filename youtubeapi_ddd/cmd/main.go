package main

import (
	"fmt"
	"log"
	"os"

	"example.com/m/infrastructure/reoisitory_imple"
	"example.com/m/usecase/video"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY not found")
	}

	yt, err := reoisitory_imple.NewYouTubeClient(apiKey)
	if err != nil {
		log.Fatalf("YouTube client error: %v", err)
	}

	usecase := video_usecase.FetchVideosUseCase{Repo: yt}
	videos, err := usecase.Execute("SHOWROOM", 3)
	if err != nil {
		log.Fatalf("Failed to fetch videos: %v", err)
	}

	fmt.Println("🎬 SHOWROOM 最新動画リスト:")
	for _, v := range videos {
		fmt.Printf("- %s\n", v.URL)
	}
}
