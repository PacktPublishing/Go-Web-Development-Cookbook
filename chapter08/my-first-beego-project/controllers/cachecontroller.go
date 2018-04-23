package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type CacheController struct {
	beego.Controller
}

var beegoCache cache.Cache
var err error

func init() {
	beegoCache, err = cache.NewCache("memory", `{"interval":60}`)
	beegoCache.Put("foo", "bar", 100000*time.Second)
}

func (this *CacheController) GetFromCache() {
	foo := beegoCache.Get("foo")
	this.Ctx.WriteString("Hello " + fmt.Sprintf("%v", foo))
}
