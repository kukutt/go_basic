package main

import (
	"github.com/swxctx/xlog"
    "github.com/kukutt/go_basic/test"
)

func main() {
	xlog.Default.Level = xlog.DebugLevel
    
    var i utils.Power = utils.Power{Age: 10, High: 178, Name: "NewMan"}
    //var i utils.Power = utils.Power{10, 178, "NewMan"}

	xlog.Debugf("type: %T", xlog.ErrorLevel)
	xlog.Debugf("value#: %#v", xlog.ErrorLevel)

    xlog.Debugf("type:%T", i)
    xlog.Debugf("value:%v", i)
    xlog.Debugf("value+:%+v", i)
    xlog.Debugf("value#:%#v", i)

    var interf interface{} = i
    xlog.Debugf("%v", interf)
}
