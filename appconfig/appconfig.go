package appconfig

import "time"

var App AppConfig

func init() {
	// App.LasttimeExec=time.Now()
}

type AppConfig struct {
	LasttimeExec time.Time //上次执行的时间
}

func (app AppConfig) CheckUploadStatus(t time.Time) bool {
	if app.LasttimeExec.Year() == t.Year() && app.LasttimeExec.Month() == t.Month() && app.LasttimeExec.Day() == t.Day() {
		return false
	} else {
		return true
	}
}
