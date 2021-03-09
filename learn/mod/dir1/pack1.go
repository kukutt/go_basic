package p1

import "runtime"

func Debug() string{
    _, file, _, _ := runtime.Caller(1)
    return "[p1.test:"+file+"]";
}

