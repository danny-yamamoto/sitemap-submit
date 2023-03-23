# sitemap-submit

## Google Search Console API v1
| スコープ | -- |
| ---- | ---- |
| https://www.googleapis.com/auth/webmasters | 確認済みサイトの検索コンソールデータを表示および管理する |
| https://www.googleapis.com/auth/webmasters.readonly | 確認済みサイトの検索コンソールデータを表示する |


## Building the Environment
- Google Search Console Labrary を追加する。
```
vscode ➜ /workspaces/sitemap-submit (main) $ sudo chmod -R 777 /go/
vscode ➜ /workspaces/sitemap-submit (main) $ go get golang.org/x/oauth2/google
go: downloading golang.org/x/oauth2 v0.6.0
go: downloading cloud.google.com/go/compute/metadata v0.2.0
go: downloading google.golang.org/appengine v1.6.7
go: downloading golang.org/x/net v0.8.0
go: downloading github.com/golang/protobuf v1.5.2
go: downloading google.golang.org/protobuf v1.28.0
go: added cloud.google.com/go/compute/metadata v0.2.0
go: added github.com/golang/protobuf v1.5.2
go: added golang.org/x/net v0.8.0
go: added golang.org/x/oauth2 v0.6.0
go: added google.golang.org/appengine v1.6.7
go: added google.golang.org/protobuf v1.28.0
vscode ➜ /workspaces/sitemap-submit (main) $
vscode ➜ /workspaces/sitemap-submit (main) $ go get google.golang.org/api/webmasters/v3
go: downloading google.golang.org/grpc v1.53.0
go: downloading github.com/google/uuid v1.3.0
go: downloading github.com/googleapis/gax-go/v2 v2.7.0
go: downloading go.opencensus.io v0.24.0
go: downloading cloud.google.com/go/compute/metadata v0.2.3
go: downloading google.golang.org/genproto v0.0.0-20230223222841-637eb2293923
go: downloading github.com/googleapis/enterprise-certificate-proxy v0.2.3
go: downloading cloud.google.com/go/compute v1.18.0
go: downloading github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e
go: downloading golang.org/x/text v0.8.0
go: downloading golang.org/x/sys v0.6.0
go: added cloud.google.com/go/compute v1.18.0
go: upgraded cloud.google.com/go/compute/metadata v0.2.0 => v0.2.3
go: added github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e
go: added github.com/google/uuid v1.3.0
go: added github.com/googleapis/enterprise-certificate-proxy v0.2.3
go: added github.com/googleapis/gax-go/v2 v2.7.0
go: added go.opencensus.io v0.24.0
go: added google.golang.org/api v0.111.0
go: added google.golang.org/genproto v0.0.0-20230223222841-637eb2293923
go: added google.golang.org/grpc v1.53.0
go: upgraded google.golang.org/protobuf v1.28.0 => v1.28.1
vscode ➜ /workspaces/sitemap-submit (main) $
```

## Set environment variables
- `credentials.json` は、Google Cloud Console で json を作成する。
```
vscode ➜ /workspaces/sitemap-submit (main) $ export SITE_URL="sc-domain:trial-net.co.jp"
vscode ➜ /workspaces/sitemap-submit (main) $ export FEEDPATH="https://www.trial-net.co.jp/sitemap/sitemap-0.xml"
vscode ➜ /workspaces/sitemap-submit (main) $ export KEY_JSON=`cat credentials.json`
vscode ➜ /workspaces/sitemap-submit (main) $ go run main.go 
```

- 新しい Terminal から実行する。
```
vscode ➜ /workspaces/sitemap-submit (main) $ curl localhost:8080
```

## 参考
- [Google Search Console API](https://developers.google.com/webmaster-tools/v1/sitemaps/submit?hl=ja)
- [Client の作成方法](https://dev.classmethod.jp/articles/bigquery-api-go-client-try/)
- [Func NewService](https://pkg.go.dev/google.golang.org/api@v0.111.0/webmasters/v3#NewService)
- [Search Console 画面を開かずに API を操作する](https://propen.dream-target.jp/blog/google-search-console-api)
