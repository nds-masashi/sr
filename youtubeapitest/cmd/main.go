package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
)

// YouTube Data APIを使用してSHOWROOMの最新動画を取得するプログラム
// このプログラムは、YouTube Data APIを使用してSHOWROOMの最新動画を取得し、
// タイトルとURLを表示します。
func main() {
	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// APIキーの取得
	apikey := os.Getenv("YOUTUBE_API_KEY")
	if apikey == "" {
		log.Fatal("API key not found")
	}

	// YouTube Data APIのクライアントを作成
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apikey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	// 検索リクエストの作成
	searchCall := youtubeService.Search.List([]string{"snippet"}).
		Q("SHOWROOM").
		Type("video").
		Order("date").
		MaxResults(100)

	// 検索リクエストの実行
	response, err := searchCall.Do()
	if err != nil {
		log.Fatalf("Error executing search request: %v", err)
	}

	// 検索結果の表示
	fmt.Println("検索結果の表示")
	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			fmt.Printf("Title: %s\n", item.Snippet.Title)
			fmt.Printf("URL: https://www.youtube.com/watch?v=%s\n", item.Id.VideoId)
			fmt.Println()
		}
	}

	// 検索結果がない場合の処理
	if len(response.Items) == 0 {
		fmt.Println("No results found.")
	}

	// エラーハンドリング
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
