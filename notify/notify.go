package notify

import (
	"net/http"
	"fmt"
)

/**
记录同步信息
通知监听的服务
 */


func NotifyBoss(url string){

	//生成client 参数为默认
	client := &http.Client{}

	//提交请求
	reqest, err := http.NewRequest("POST", url, nil)

	if err != nil {
		fmt.Println(err)
	}
	//处理返回结果
	client.Do(reqest)
}