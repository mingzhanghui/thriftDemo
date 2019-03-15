# thriftDemo
https://studygolang.com/articles/9607

安装 Thrift 的 Golang 库有两种方案：

直接通过 go get 命令安装，缺点是因为不可抗拒的网络因素大部分人可能会失败：$ go get git.apache.org/thrift.git/lib/go/thrift
通过源码安装：
在 $GOPATH 的 src 目录下创建多层级目录：git.apache.org/thrift.git/lib/go
从 github 上下载 thrift 0.10.0 的源码 ，解压后进入 thrift-0.10.0/lib/go 目录下，将 thrift 目录 copy 到刚创建的 $GOPATH/src/git.apache.org/thrift.git/lib/go 目录下
在任意目录下执行 $ go install git.apache.org/thrift.git/lib/go/thrift 就完成了 golang 的 thrift 库的安装
安装 Thrift 的 IDL 编译工具

windows 平台下安装：
直接下载：thrift complier 下载地址，下载完成后改名为：thrift.exe 并将其放入到系统环境变量下即可使用

Linux 平台下安装：
从 github 上下载 thrift 0.10.0 的源码，解压后进入：thrift-0.10.0/compiler/cpp 目录执行如下命令完成编译后
$ mkdir cmake-build
$ cd cmake-build
$ cmake ..
$ make

$ sudo cp bin/thrift  /usr/local/bin/

验证是否安装成功：

$ thrift -version，如果打印出来：Thrift version 0.10.0 表明 complier 安装成功


