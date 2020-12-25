package lib

import (
	"fmt"
	"os/exec"
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
					go consumeRebuild()
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
	command := "/home/admin/gin_server_v1/build.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	resp := string(bytes)
	fmt.Println(resp)
}

func TestTask() {
	testChan <- 1
}
