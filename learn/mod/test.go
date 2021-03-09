package main

import (
    "fmt"
    "test/dir1"
    "test/dir2"
    "test/dir3/dirsub"
    "github.com/kukutt/go_basic/test"
)

func main(){
    /* 本地目录包调用 import是目录名 可以不与包名字相同 */
    fmt.Println(p1.Debug())
    fmt.Println(pack2.Debug())
    fmt.Println(p3.Debug())

    /* 网络包调用 */
    utils.Printtest()
}
