package p3

import "runtime"

func Debug() string{
    _, file, _, _ := runtime.Caller(1)
    return "[p3.test:"+file+"]";
}

