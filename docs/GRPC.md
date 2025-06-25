# info

- proto 文件中只有定义了 service 块，使用 protoc-gen-go-grpc 才会生成 grpc.go 代码，否则只会生成 pb.go 代码

# 安装

- protoc

```shell
cd /tmp/
wget https://github.com/protocolbuffers/protobuf/releases/download/v29.1/protoc-29.1-linux-x86_64.zip
mkdir protobuf-29.1/
unzip protoc-29.1-linux-x86 64.zip -d protobuf-29.1/
cd protobuf-29.1/
sudo cp a include/* /usr/local/include
sudo cp bin/protoc /usr/local/bin/
protoc --version #查看 protoc 版本，成功输出版本号，说明安装成功libprotoc 29.1
```

- 其它插件

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.2
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
go install github.com/grpc-ecosystem/grpc gateway/v2/protoc gen grpc gateway@v2.24.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.24.0
go install github.com/onexstack/protoc-gen-defaults@v0.0.2
```

# protocol buffers v3 语法

- 李文周博客：https://www.liwenzhou.com/posts/Go/Protobuf3-language-guide-zh/
- 关于 package 和 go_package：https://blog.csdn.net/zhangh571354026/article/details/123852629?share_token=768dd728-1d65-4426-a372-9f14e57f54cb
- optional,oneof,FieldMask：https://www.liwenzhou.com/posts/Go/oneof-wrappers-field_mask/

- import protoc 的路径相对于生成代码时 `--proto_path=` 的路径

# vscode 插件相关

配置文件：
防止导入 protoc 飘红

```json
  "protoc": {
    "options": [
      // 指定 protoc 文件地址
      "--proto_path=miniblog/pkg/api/apiserver/v1"
    ]
  }
```
