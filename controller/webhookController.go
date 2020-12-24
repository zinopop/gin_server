package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"os/exec"
	"syscall"
)

type Cmd struct {
	Path  string    //运行命令的路径，绝对路径或者相对路径
	Args  []string  // 命令参数
	Env   []string  //进程环境，如果环境为空，则使用当前进程的环境
	Dir   string    //指定command的工作目录，如果dir为空，则comman在调用进程所在当前目录中运行
	Stdin io.Reader //标准输入，如果stdin是nil的话，进程从null device中读取（os.DevNull），stdin也可以时一个文件，否则的话则在运行过程中再开一个goroutine去

	Stdout       io.Writer //标准输出
	Stderr       io.Writer //错误输出，如果这两个（Stdout和Stderr）为空的话，则command运行时将响应的文件描述符连接到os.DevNull
	ExtraFiles   []*os.File
	SysProcAttr  *syscall.SysProcAttr
	Process      *os.Process      //Process是底层进程，只启动一次
	ProcessState *os.ProcessState //ProcessState包含一个退出进程的信息，当进程调用Wait或者Run时便会产生该信息．
}

func Rebuild(c *gin.Context) {
	command := "/home/admin/gin_server_v1/build.sh"
	CmdBash(command) //重定向

}

func CmdBash(commandName string) {
	command := exec.Command(commandName) //初始化Cmd
	err := command.Start()               //运行脚本
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println("Process PID:", command.Process.Pid)
	err = command.Wait() //等待执行完成
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println("ProcessState PID:", command.ProcessState.Pid())
}
