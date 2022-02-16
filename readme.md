# dapr demo

> 原生grpc client 负载均衡调用

> 通过 dapr grpc proxy 特性

> 多grpc服务调用

## server
sh server.sh

## client
sh client.sh

## dapr version

```shell
CLI version: 1.6.0 
Runtime version: 1.6.0
```

## 服务发现
> https://docs.dapr.io/reference/components-reference/supported-name-resolution
1. 默认 mdns
2. consul(待测试)

## 负载均衡

## 引用
1. https://docs.dapr.io/operations/configuration/preview-features
2. https://docs.dapr.io/developing-applications/building-blocks/service-invocation/howto-invoke-services-grpc