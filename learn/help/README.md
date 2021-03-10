# 天下代码一半抄,该如何查找资料

## go.dev网站
[go.dev](https://go.dev/)

该网站可以方便查找到想要的库/包。

当然，github也是个不错的选择。

## go doc
查找某包下的所有`成员变量`和`成员方法`，使用方法如下
```
go doc -all github.com/kukutt/go_basic/test 
```

## %v %
如果想知道某个变量类型，可以使用%T%v，代码如下
```
xlog.Debugf("type:%T", i)
xlog.Debugf("value:%v", i)
xlog.Debugf("value+:%+v", i)
xlog.Debugf("value#:%#v", i)
```
打印信息
```
2021/03/11 00:19:28 [DBUG] type:utils.Power [test.go:17]
2021/03/11 00:19:28 [DBUG] value:{10 178 NewMan} [test.go:18]
2021/03/11 00:19:28 [DBUG] value+:{Age:10 High:178 Name:NewMan} [test.go:19]
2021/03/11 00:19:28 [DBUG] value#:utils.Power{Age:10, High:178, Name:"NewMan"} [test.go:20]
```
疑问，如何知道`utils.Power`属于哪个包呢?
