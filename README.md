# Protobuf
[![编译状态](https://github.ruijc.com:20443/api/badges/dronestock/protobuf/status.svg)](https://github.ruijc.com:20443/dronestock/protobuf)
[![Golang质量](https://goreportcard.com/badge/github.com/dronestock/protobuf)](https://goreportcard.com/report/github.com/dronestock/protobuf)
![版本](https://img.shields.io/github/go-mod/go-version/dronestock/protobuf)
![仓库大小](https://img.shields.io/github/repo-size/dronestock/protobuf)
![最后提交](https://img.shields.io/github/last-commit/dronestock/protobuf)
![授权协议](https://img.shields.io/github/license/dronestock/protobuf)
![语言个数](https://img.shields.io/github/languages/count/dronestock/protobuf)
![最佳语言](https://img.shields.io/github/languages/top/dronestock/protobuf)
![星星个数](https://img.shields.io/github/stars/dronestock/protobuf?style=social)

`Drone`持续集成`Protobuf`插件，功能有

- 支持绝大部分开发语言（包括：`Go`、`Java`、`Swift`、`Python`、`Javascript等等`）
- 使用简单，只需要简单的配置（可以做到零配置，默认生成`Go`代码）就能使用本插件
- 增加部分语言的扩展支持（比如`Go`语言增加了标签注入）

## 包含插件

- [ckaznocha/protoc-gen-lint](https://github.com/ckaznocha/protoc-gen-lint)
- [danielvladco/go-proto-gql](https://github.com/danielvladco/go-proto-gql)
- [dart-lang/protobuf](https://github.com/dart-lang/protobuf)
- [envoyproxy/protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)
- [mwitkow/go-proto-validators](https://github.com/mwitkow/go-proto-validators)
- [gogo/protobuf](https://github.com/gogo/protobuf)
- [golang/protobuf](https://github.com/golang/protobuf)
- [google/protobuf](https://github.com/google/protobuf)
- [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [grpc/grpc](https://github.com/grpc/grpc)
- [grpc/grpc-go](https://github.com/grpc/grpc-go)
- [grpc/grpc-java](https://github.com/grpc/grpc-java)
- [grpc/grpc-swift](https://github.com/grpc/grpc-swift)
- [grpc/grpc-web](https://github.com/grpc/grpc-web)
- [improbable-eng/ts-protoc-gen](https://github.com/improbable-eng/ts-protoc-gen)
- [protobuf-c/protobuf-c](https://github.com/protobuf-c/protobuf-c)
- [pseudomuto/protoc-gen-doc](https://github.com/pseudomuto/protoc-gen-doc)
- [stepancheg/grpc-rust](https://github.com/stepancheg/grpc-rust)
- [stepancheg/rust-protobuf](https://github.com/stepancheg/rust-protobuf)
- [TheThingsIndustries/protoc-gen-fieldmask](https://github.com/TheThingsIndustries/protoc-gen-fieldmask)
- [TheThingsIndustries/protoc-gen-gogottn](https://github.com/TheThingsIndustries/protoc-gen-gogottn)

## 支持语言

- C
- C#
- C++
- Dart / Flutter
- Go / Gogo
- Java / JavaNano (Android)
- JavaScript
- Objective-C
- PHP
- Python
- Ruby
- Rust
- Swift
- Typescript

## 使用

```yaml
steps:
  - name: 编译
    image: dronestock/protobuf
    settings:
      targets:
        - lang: go
          output: $${GO}
          opt: module=github.com/storezhang/transfer
        - lang: java
          output: $${JAVA}
```

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢
[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png)](https://www.jetbrains.com/?from=dronestock/protobuf)
