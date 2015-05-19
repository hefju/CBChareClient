package main

import (
	//    "log"
	//    "net"
	//    "net/rpc/jsonrpc"
	//    "github.com/hefju/CBChareClient/myconfig"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hefju/CBChareClient/appconfig"
	"github.com/hefju/CBChareClient/models"
	"github.com/hefju/CBChareClient/tasker"

	"io/ioutil"
	"net/http"
	"time"
   // "github.com/donnie4w/go-logger/logger"
	//"log"
    "github.com/hefju/CBChareClient/myconfig"
    "github.com/donnie4w/go-logger/logger"
    "github.com/hefju/CBChareClient/jutool"
)

func main() {

	//系统初始化,设置日志类,加载系统数据
	appconfig.LoadConfig()

	task := tasker.ChareTask //充电上传任务
    email:=tasker.EmailSend{}

	//定时执行
	ticker := time.NewTicker(time.Minute*10) //time.Minute*30)//time.Second * 5)
	for t := range ticker.C {
        switch t.Hour() {
            case 4: //凌晨4点时,触发事件
            if task.CheckUploadStatus(t) {
                // fmt.Println("CheckUploadStatus")
                task.UploadBills(t)
                task.SetLastExecuteTime(t)//更新上次执行时间
                //task.LasttimeExec = t
            }
            case 21://21  晚上9点钟,发送准备信息. 如果没有收到就要检查了
            if email.CheckUploadStatus(t){
                email.SendEmail()
                email.SetLastExecuteTime(t)
            }
            default:
            sentStatus(t)
           // email.LasttimeExec

           // fmt.Println("ticker17")
        }
//		if t.Hour() == 0 { //0凌晨时段,触发事件
//			//检查上传状态,如果还未上传就触发上传操作
//            //fmt.Println("ticker17")
//			if task.CheckUploadStatus(t) {
//               // fmt.Println("CheckUploadStatus")
//                //task.UploadBills(t)
//                task.SetLastExecuteTime(t)//更新上次执行时间
//				//task.LasttimeExec = t
//			}
//		}
//		fmt.Println(t)
	}

}

//发送状态报告
func sentStatus(t time.Time) {
    url :=myconfig.ReportAddr  // "http://localhost:8083/upload" // "http://localhost:8083/upload"  192.168.1.200
    report:=models.StatusReport{From:myconfig.MyName,FromTime:time.Now().String(),Title:"状态报告",Content:"I'm still alive"}
    jsonStr, _ := json.Marshal(report)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        //panic(err)//发送数据失败, 程序不要死啊. panic就完啦
        defer func() {
            if r := recover(); r != nil { //r的信息有什么用?还不如直接输出err
                logger.Error("状态报告失败,", err)//  fmt.Println("发送失败,", r, err)
               // jutool.SendEmail("状态报告失败","报告时间:"+t.Format("2006-01-02 15:04:05")+" "+err.Error())//上传失败也发个email通知我
            }
        }()
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))
    logger.Info("状态报成功,",t.Format("2006-01-02 15:04:05"), string(body))
}

func updateBills(t time.Time) {

	t = t.AddDate(0, 0, -1) //设置时间为当前时间的前一天

	bills := models.GetChargeListByDate(t) //获取一天的数据
	url := "http://localhost:8083/upload"
	jsonStr, _ := json.Marshal(bills)

	//    bill := bills[0]
	//    fmt.Println("bill: ",bill)
	//    url := "http://localhost:8083/uploadone"
	//    jsonStr, _ := json.Marshal(bill)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		//panic(err)
		//发送数据失败, 程序不要死啊.
		defer func() {
			if r := recover(); r != nil { //r的信息有什么用?还不如直接输出err
				// fmt.Println("Recovered in f, 发送失败", err)
				fmt.Println("发送失败,", r, err)
				// return
			}
		}()
		// fmt.Println("发送失败,",err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
