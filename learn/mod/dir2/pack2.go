package pack2

import "runtime"

func Debug() string{
    _, file, _, _ := runtime.Caller(1)
    return "[pack2.test:"+file+"]";
}

