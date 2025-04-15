
# how to create

create .gitignore

create go.mod
```sh
go mod init github.com/omikuu/sr' to initialize a v0 or v1 module
```

download module
```sh
go get google.golang.org/api/youtube/v3
go get github.com/joho/godotenv
```

add main method

goplsがないとコード補完や、参照機能が動作しない
```sh
go install golang.org/x/tools/gopls@latest
```

# reference

## YouTube Data API v3 を使って YouTube 動画を検索する
https://qiita.com/koki_develop/items/4cd7de3898dae2c33f20

1. Google Cloud Console にログイン
2. プロジェクトを新規作成
3. 「API とサービス」→「ライブラリ」→YouTube Data API v3 を検索して有効化
4. 「認証情報」タブで APIキー を作成

## Go + YouTube Data API v3 で特定のチャンネルの動画データ一覧を取得する
https://qiita.com/tokuchi765/items/31c3766c3b7963dc6e64

## YouTube API Reference
https://qiita.com/tokuchi765/items/31c3766c3b7963dc6e64#:~:text=YouTube%20API%20Reference

# other

`curl 'https://www.googleapis.com/youtube/v3/search?key=APIキーを入れる&type=video&part=snippet&q=dog'`

