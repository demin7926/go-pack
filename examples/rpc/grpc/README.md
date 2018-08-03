# a gRPC demo.

## 打开此Demo的正确姿势
1. 定义.proto文件    
2. 用`protoc`工具编译生成对应的go类    
    ```bash
    protoc --go_out=plugins=grpc:. greeter.proto
    ```
    关于`protoc`工具及`protoc-gen-go`的安装与使用，参考[gRPC Quick Start](https://grpc.io/docs/quickstart/go.html#install-grpc)和[protobuf-go-generated](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
3. 编写server端代码        
4. 编写client端代码    
5. 编译/运行server    
    ```bash
    > cd server
    > go build -o server
    > ./server
    ```
6. 编译/运行client    
    ```bash
    > cd client
    > go build -o client
    > ./client
    ```


