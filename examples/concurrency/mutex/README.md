## 参考：
- https://tour.golang.org/concurrency/9
- https://gobyexample.com/mutexes
- https://godoc.org/sync

## 收获
- 运行测试时使用`-race`选项模拟并发测试:
    > You should test your server with -race option and then eliminate all the race conditions it throws. That way you can easier eliminate such errors before they occur.
    ```bash
    go run -race server.go
    ```

## 结论
- Build-in map类型是非线程/协和安全的，在并发环境下容易导致数据冲突，甚至是Runtime Fatal Error: concurrent map read and map write.

