package appconfig

import "time"

var App AppConfig

func init() {
    //需要从数据库或者文件读取上次时间, 为了保证程序关闭了又重新打开时候, 不会继续执行上传
	// App.LasttimeExec=time.Now()
}

type AppConfig struct {
	LasttimeExec time.Time //上次执行的时间
}

//检测上传状态, 如果上次更新日期跟今次的更新日期一样就不需要上传了, 如果更新日期不相符就返回true, 执行更新
func (app AppConfig) CheckUploadStatus(t time.Time) bool {
	if app.LasttimeExec.Year() == t.Year() && app.LasttimeExec.Month() == t.Month() && app.LasttimeExec.Day() == t.Day() {
		return false
	} else {
		return true
	}
}
