package main

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

var HttpServer *gin.Engine

func main() {
	//使用最多的逻辑处理核心
	runtime.GOMAXPROCS(runtime.NumCPU())


}
