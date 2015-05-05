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

//type Args struct {
//    Who string
//}
//type Reply struct {
//    Message string
//}

func main() {

	//初始化,加载上次
	LoadConfig()
	app := appconfig.App

	//fmt.Println( app.LasttimeExec)
	ticker := time.NewTicker(time.Second * 3) //time.Minute*30)
	for t := range ticker.C {
		if t.Hour() == 17 { //0凌晨时段,触发事件
			//检查上传状态,如果还未上传就触发上传操作
			fmt.Println("CheckUploadStatus")
			if app.CheckUploadStatus(t) {
				fmt.Println("updateBills")
				updateBills(t)
				app.LasttimeExec = t
			}
		}
		fmt.Println(t)
	}

	//fmt.Println("from server:")

	//    client, err := net.Dial("tcp",myconfig.RemoteIpt+":8081")
	//    if err != nil {
	//        log.Fatal("dialing:", err)
	//    }
	//
	//    args := &Args{"caonima"}
	//    reply := new(Reply)
	//    c := jsonrpc.NewClient(client)
	//    err = c.Call("Hello.Say", args, reply)
	//    if err != nil {
	//        log.Fatal("mygod:", err)
	//    }
	//    fmt.Println("from server:", reply)
}

func updateBills(t time.Time) {

	t = t.AddDate(0, 0, -1)
	fmt.Println("updateBills time:", t)
	return

	bills := models.GetChargeListByDate(t)
	bill := bills[0]

	url := "http://localhost:8083/uploadone"
	jsonStr, _ := json.Marshal(bill)

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
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//初始化系统,读取上次执行的最后状态
func LoadConfig() {
	//appconfig.App
}
