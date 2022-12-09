## gRPC示例

gRPC的简单示例项目，使用方式如下。

- 启动方式如下：
    ```shell
    make serv
    ``` 
    会同时启动http服务和grpc服务。

- 重新生成`go`文件

    ```shell
    make proto
    ```

- 在Docker中运行

    ```shell
    docker compose up
    ```

- 进入gRPC客户端 REPL模式，直观的进行调试

    ```shell
    make evans
    ```
