# README
## windows安裝方式
* 當前go版本為1.18.1
* 下載protoc-21.10-win64.zip [release](https://github.com/protocolbuffers/protobuf/releases)
* protoc-21.10-win64.zip解壓縮後，裡面會有bin與include資料夾
* 將protoc-21.10-win64改名為protoc3，放置C槽，並將路徑加入環境變數中(在搜尋打環境變數)
* 其他語言可直接使用protoc來轉換，go要再另外下載protoc-gen-go
```sh
go get -u github.com/golang/protobuf/protoc-gen-go
```
* 下載完後，GOPATH/bin底下會有protoc-gen-go
* GOPATH/bin也要在環境變數中，才可供全域使用
* 建立.proto file，注意要預設go_package，也就是.pb.go產生位置
```go
option go_package = "./test";
```
* 使用go，protoc指令需要多增加plugin參數
```sh
# .pb.go產生於預設的test/
protoc.exe --plugin=protoc-gen-go.exe --go_out=. simple.proto

# .pb.go產生於go/，要先自行建立go/
protoc.exe --plugin=protoc-gen-go.exe --go_out=go simple.proto
```
* 以上，即可確定protoc安裝成功
* 參考
  * [golang在windows下安装和使用protobuf](https://studygolang.com/articles/8804)
  * [protoc-gen-go: unable to determine Go import path for "simple.proto"](https://stackoverflow.com/questions/70586511/protoc-gen-go-unable-to-determine-go-import-path-for-simple-proto)