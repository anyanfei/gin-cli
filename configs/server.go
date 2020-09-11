package configs

func SetServerConfig() (serverConfig map[string]string){
	serverConfig = make(map[string]string)
	serverConfig["HOST"] = "0.0.0.0"            //监听地址
	serverConfig["PORT"] = "8080"               //监听端口
	serverConfig["VIEWS_PATTERN"] = "views/*/*" //模板路径pattern
	serverConfig["ENV"] = "debug"               // 环境模式 release/debug/test
	return
}
