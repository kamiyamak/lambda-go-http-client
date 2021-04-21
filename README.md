# lambda-go-http-client

# 紹介
環境変数に設定したURLに対してGETリクエストをするだけのプログラム<br>
AWS lambdaで利用する際は環境変数REQUEST_URLにリクエストするURLを指定します

## 使いかた
lambdaの設定でハンドラーを"httpget"に設定した場合

パッケージインストール
```
$ go get -u github.com/aws/aws-lambda-go/lambda
```
build
```
$ GOOS=linux GOARCH=amd64 go build -o httpget
```
zipに圧縮
```
$ zip httpget.zip ./httpget
```
zipをlambdaにアップロード