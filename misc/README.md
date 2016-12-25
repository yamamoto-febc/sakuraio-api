# go-swaggerでのコード再生成

- さくらのIoT Platform APIドキュメントから[定義ファイル](https://api.sakura.io/v1/docs/api.yml)を取得

```bash
$ curl -o api.yml https://api.sakura.io/v1/docs/api.yml
```

- 必要に応じてapi.ymlを修正

- go-swaggerでコード生成

```bash
$ cd gen
$ docker run --rm -it -v /path/to/your/project/sakura-iot:/go/src/github.com/yamamoto-febc/akuraio-api -w /go/src/github.com/yamamoto-febc/sakuraio-api/gen quay.io/goswagger/swagger generate client -f ../misc/api.yml -A sakuraIoT
```

## REMARK

### api.ymlの修正箇所

- `/services`のGET/PUTをtypeごとに再定義
- `DELETE`系のメソッドにレスポンス`204`を追加

### 生成されたコードの修正箇所

現在go-swaggerでのコード生成後、以下の部分を手作業で修正しています。

- gen/models/incomming_webhook_service_update.go
- gen/models/outgoing_webhook_service_update.go

```
    // 修正前
	// secret
	Secret string `json:"secret,omitempty"`
```

```
    // 修正後
	// secret
	Secret string `json:"secret"`
```
