# playground : go-grpc-server

- gRPC에 대해 학습하며 gRPC 기반의 golang server를 만들어봅니다

## tech spec

- back-end
  - gRPC (HTTP/2)
  - golang 1.18
  - sqlboiler
  - ...

- Infra
  - CI/CD based on github-action
  - ...


## directory structure

``` 
protos : protobuf
gen : generated protobuf from protos/protobuf
server : server
  ⊢ handler
  ⊢ db
  ㄴ ...
config : config
```