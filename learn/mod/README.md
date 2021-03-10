# go mod 实验

## 本地包

## 网络包

### 网络包添加
go get 下载指定包
```
# 最新版本下载
go get -u github.com/kukutt/go_basic/test
# 指定版本
go get github.com/kukutt/go_basic/test@a340a6b5119
```python
go mod tidy 整理全部包

## 下载依赖包到本地目录
1. 依赖包出现问题;
2. 想查看依赖包的代码;
3. ide工具下代码查看和搜索，一般包含子目录的源码才可以方便搜索;
可以使用一下命令复制和编译
```
go mod vendor              # 此时会创建vendor目录到当前目录
go build -mod=vendor xx.go # 使用vendor下包编译,测试下来发现默认此项
go build -mod=mod xx.go    # 使用mod管理的包编译
```
