# ToadOCRRPCClient

## 概述

​	ToadOCRRpcClient是ToadOCR PRC服务体系的客户端，其提供对ToadOCREngine与ToadOCRPreprocessor服务的调用方式，使得新服务可以快速接入ToadOCR PRC服务体系。

## 获取方式

```sh
 go get github.com/suvvm/ToadOCRRpcClient
```

## 使用方法

### EngineClient

- 构建客户端

```go
client := rpc.NewEngineClient(appID, appSecret, discoverUrl)
```

- 初始化与服务端的链接

```go
if err := client.InitEngineClient(); err != nil {
   log.Printf("init engine client fail!")
   return err
}
```

- 调用OCR服务

```go
preVal, err := client.Predict("snn", imageBytes)
if err != nil {
		log.Printf("call rpc fail")
		return err
}
```

### ProcessorClient

- 构建客户端

```go
client := rpc.NewProcessorClient(appID, appSecret, discoverUrl)
```

- 初始化与服务端的链接

```go
if err := client.InitProcessorClient(); err != nil {
		log.Printf("init processor client fail!")
		return err
	}
```

- 调用OCR服务

```go
preVal, err := client.Process("snn", imageBytes)
if err != nil {
		log.Printf("call rpc fail")
		return err
}
```

## FAQ

```sh
# github.com/coreos/etcd/clientv3/balancer/picker
../../../pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/picker/err.go:37:44: undefined: balancer.PickOptions
../../../pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/picker/roundrobin_balanced.go:55:54: undefined: balancer.PickOptions
# github.com/coreos/etcd/clientv3/balancer/resolver/endpoint
../../../pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:114:78: undefined: resolver.BuildOption
../../../pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:182:31: undefined: resolver.ResolveNowOption
```

​	该问题由依赖的etcdclientv3引发，grpc的更行导致了部分命名规则的改变与部分方法的废弃，etcd社区工作者并未认真解决该问题，多个issue都在扯皮中并建议依赖其项目的开发者使用grpc@v1.26.0来应付该问题，该问题预计在2021年6月～2022年6月之间修复。