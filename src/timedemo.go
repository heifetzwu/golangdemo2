package main

import (
	"fmt"
	"time"
)

func timedemo() {
	//unix time
	fmt.Println("unix: ", time.Now().Unix())

	//取到 nano second
	fmt.Println("unix: ", time.Now().UnixNano())

	//format成 正常格式化後的時間
	fmt.Println("datetime: ", time.Now().Format("2006/03/02 15:04:05"))

	//after
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("hello world")
	})

	//這邊 Sleep 5s 是為了讓上面的 AfterFunc 會執行，不然就像前面章節有講到的 只要 main thread 結束，任何 sub thread 都會跟著一起結束
	time.Sleep(5 * time.Second)
	fmt.Println("end")
}
