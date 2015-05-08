package tasker
import (
    "time"
    "github.com/hefju/CBChareClient/models"
    "github.com/hefju/CBChareClient/myconfig"
    "encoding/json"
    "net/http"
    "bytes"
    "io/ioutil"
    "github.com/donnie4w/go-logger/logger"
    "github.com/hefju/CBChareClient/jutool"
)


var ChareTask ChareUpload//充电上传任务

//
type ChareUpload struct {
    LasttimeExec time.Time //上次执行的时间
}



func (uploader *ChareUpload)UploadBills(t time.Time){

    t = t.AddDate(0, 0, -1) //设置时间为当前时间的前一天
    logger.Info("触发上传任务,上传数据时间:",t.Format("2006-01-02 15:04:05"))

    bills := models.GetChargeListByDate(t) //获取一天的数据
    url :=myconfig.UploadAddr  // "http://localhost:8083/upload" // "http://localhost:8083/upload"  192.168.1.200
    jsonStr, _ := json.Marshal(bills)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        //panic(err)//发送数据失败, 程序不要死啊. panic就完啦
        defer func() {
            if r := recover(); r != nil { //r的信息有什么用?还不如直接输出err
                logger.Error("发送失败,", err)//  fmt.Println("发送失败,", r, err)
            }
        }()
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))
    logger.Info("上传成功,",t.Format("2006-01-02 15:04:05"), string(body))
}

//检测上传状态, 如果上次更新日期跟今次的更新日期一样就不需要上传了, 如果更新日期不相符就返回true, 执行更新
func (uploader *ChareUpload) CheckUploadStatus(t time.Time) bool {
    if uploader.LasttimeExec.Year() == t.Year() && uploader.LasttimeExec.Month() == t.Month() && uploader.LasttimeExec.Day() == t.Day() {
        return false
    } else {
        return true
    }
}

//更新上次执行时间
func (uploader *ChareUpload)SetLastExecuteTime(t time.Time) {
    uploader.LasttimeExec=t
}


type EmailSend struct {
    LasttimeExec time.Time //上次执行的时间
}

func (sender *EmailSend)SendEmail() {
    jutool.SendEmail("充电信息提示","设备准备就绪.")
}


//检测上传状态, 如果上次更新日期跟今次的更新日期一样就不需要上传了, 如果更新日期不相符就返回true, 执行更新
func (sender *EmailSend) CheckUploadStatus(t time.Time) bool {
    if sender.LasttimeExec.Year() == t.Year() && sender.LasttimeExec.Month() == t.Month() && sender.LasttimeExec.Day() == t.Day() {
        return false
    } else {
        return true
    }
}

//更新上次执行时间
func (sender *EmailSend)SetLastExecuteTime(t time.Time) {
    sender.LasttimeExec=t
}