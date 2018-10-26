# Golang单元测试Demo

## go test说明
- 在需要测试的包中编写功能函数/模块，以及对应的测试案例，测试案例放在`_test.go`后缀的文件中
- 执行`go test -v -cover=true PACKAGE [-run TEST_CASE_NAME]`进行测试，如本例：    
    ```bash
    # 执行leetcode包下的所有测试用例
    go test -v -cover=true ./leetcode
    # 只执行leetcode包下的TestSuccessTwoSum这个用例
    go test -v -cover=true ./leetcode -run TestSuccessTwoSum
    ```



