package libs

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"time"
)

func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
//日志等级转换函数
func Log_Level(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
func Create_File(log_path string)(bool)  {
	err := os.MkdirAll(log_path,0771)
	if err == nil{
		return true
	}else {
		return false
	}
	return false
}

func  NewLog()(FileLogs1  *logs.BeeLogger){

	//ConsoleLogs := logs.NewLogger(1000)
	//ConsoleLogs.SetLogger("console")
	FileLogs := logs.NewLogger(1000)
	y := time.Now().Format("2006")
	m := time.Now().Format("01")
	d := time.Now().Format("02")
	path := beego.AppConfig.String("log_path")
	log_path := path + "/" + y + "/" + m + "/"
	logs_file := d + "_log.txt"
	logs_dir := log_path + logs_file
	if !PathExists("log_path"){
		if !Create_File(log_path){
			beego.Info("创建文件夹:log_path = %v 失败",log_path)
		}
	}
	//log_level := beego.AppConfig.String("log_level")
	config := make(map[string]interface{})
	config["filename"] = logs_dir
	//config["level"] = Log_Level(log_level)
	config["level"] = logs.LevelDebug
	configStr,err := json.Marshal(config)
	if err != nil{
		beego.Info("json转化失败")
		return
	}
	FileLogs.SetLogger(logs.AdapterFile,string(configStr))
	FileLogs.EnableFuncCallDepth(true)
	return FileLogs
}



