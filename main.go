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
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	//初始化,加载上次
	LoadConfig()
	app := appconfig.App

	//fmt.Println( app.LasttimeExec)
	ticker := time.NewTicker(time.Second * 5)//time.Minute*30)//time.Second * 5)
	for t := range ticker.C {
		if t.Hour() == 10 { //0凌晨时段,触发事件
			//检查上传状态,如果还未上传就触发上传操作
			//fmt.Println("CheckUploadStatus")
			if app.CheckUploadStatus(t) {
				//fmt.Println("updateBills")
				updateBills(t)
				app.LasttimeExec = t //更新上次执行时间
			}
		}
		//fmt.Println(t)
	}


}

func updateBills(t time.Time) {

	t = t.AddDate(0, 0, -1)//设置时间为当前时间的前一天
//	fmt.Println("updateBills time:", t)
//	return

	bills := models.GetChargeListByDate(t)  //获取一天的数据
    url := "http://localhost:8083/upload"
    jsonStr, _ := json.Marshal(bills)

//    bill := bills[0]
//    fmt.Println("bill: ",bill)
//    url := "http://localhost:8083/uploadone"
//    jsonStr, _ := json.Marshal(bill)






	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//初始化系统,读取上次执行的最后状态
func LoadConfig() {
	//appconfig.App
}
