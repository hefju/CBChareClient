package appconfig

import (
    //"time"
    "github.com/donnie4w/go-logger/logger"
    "os"
)

//var App AppConfig

func init() {
    //需要从数据库或者文件读取上次时间, 为了保证程序关闭了又重新打开时候, 不会继续执行上传
	// App.LasttimeExec=time.Now()
}

func LoadConfig() {
    //logger.SetConsole(false)//默认是输出到控制台的, 所以logger.SetConsole(true) 写不写都无所谓
    os.MkdirAll("./log", 0777)                  //创建log文件夹, 用来存放日志
    logger.SetRollingDaily("./log", "test.log") //如果没有log文件夹, 需要新增文件夹
    logger.SetLevel(logger.DEBUG)
}



//type AppConfig struct {
//    LasttimeExec time.Time //上次执行的时间
//}
//
////检测上传状态, 如果上次更新日期跟今次的更新日期一样就不需要上传了, 如果更新日期不相符就返回true, 执行更新
//func (app AppConfig) CheckUploadStatus(t time.Time) bool {
//	if app.LasttimeExec.Year() == t.Year() && app.LasttimeExec.Month() == t.Month() && app.LasttimeExec.Day() == t.Day() {
//		return false
//	} else {
//		return true
//	}
//}
