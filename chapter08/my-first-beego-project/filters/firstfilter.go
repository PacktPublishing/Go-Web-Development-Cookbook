package filters

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/context"
)

var LogManager = func(ctx *context.Context) {
	fmt.Println("IP :: " + ctx.Request.RemoteAddr + ", Time :: " + time.Now().Format(time.RFC850))
}
