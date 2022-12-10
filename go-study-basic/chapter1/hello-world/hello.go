package main

import "fmt"

// go mod tidy ソースを解析し、必要なサードパーティーの依存関係を解決する
// go build
// go.modがあると、ファイルを指定しなくてもそのディレクトリに置かれた.goファイルを解析してコンパイルしてビルドしてくれる。
// ./hello-worldで実行可能
func main() {
  fmt.Println("goのmoduleはこんな感じ")
}