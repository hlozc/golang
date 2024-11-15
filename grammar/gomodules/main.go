package main

/*
* @ TEST 1 Go Modules
* 1. 是 Go 的依赖解决方案
* 2. 可以通过 Go Modules 的方式创建一个项目，但是建议和 GOPATH 分开

* @ TEST 2 go mod 命令
go mod init 生成 go.mod 文件
go mod download 下载 go.mod 文件中指定的所有依赖
go mod tidy 整理现有的依赖
go mod graph 查看现有的依赖结构
go mod edit 编辑 go.mod 文件
go mod vendor 到处项目所有依赖到 vendor 目录
go mod verify 检验一个模块是否有被篡改过
go mod why 查看为什么需要依赖某模块

* @ TEST 3 环境变量详细
1. GO111MODULE 用来作为 Go Modules 的开关。auto 表示【只要项目包含 go.mod 文件的话启动 Go Modules】
     on 表示启动，off 表示关闭
   go env -w GO111MODULE=on
2. GOPROXY 就是 Go 模块的代理，作用就是用来让 Go 在后续拉去模块版本的时候直接通过镜像站点来快速拉取
3. GOSUMDB 用来检验下载过来的时候没有经过篡改
4. GONOPROXY/GONOSUMDB/GOPRIVATE 表示哪些不需要代理/校验/私有
*/
