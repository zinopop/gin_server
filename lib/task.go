package lib

import (
	"fmt"
	"time"
)

var testChan = make(chan int)
var buildChan = make(chan int)

func init() {
	go func() {
	LOOP:
		for {
			select {
			case res, ok := <-testChan:
				if ok {
					fmt.Println("chan number:", res)
					continue LOOP
				}
			case _, ok := <-buildChan:
				if ok {
					consumeRebuild()
					continue LOOP
				}
			}
		}
	}()
	fmt.Println("task异步协程启动成功....")
}

func RebuildTask() {
	buildChan <- 1
}

func consumeRebuild() {
	fmt.Println("start")
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("end")
	//如果失败启动协程重新消费阻塞后面的队列
	go func() {
		buildChan <- 1
	}()
	//command := "/home/admin/gin_server_v1/build.sh"
	//cmd := exec.Command("/bin/bash", "-c", command)
	//bytes, err := cmd.Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//resp := string(bytes)
	//fmt.Println(resp)
}

func TestTask() {
	testChan <- 1
}
