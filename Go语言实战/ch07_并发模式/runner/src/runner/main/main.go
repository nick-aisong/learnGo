// 这个示例程序演示如何使用通道来监视
// 程序运行的时间，以在程序运行时间过长
// 时如何终止程序
package main

import (
	"log"
	"os"
	"runner"
	"time"
)

// timeout 规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

// main 是程序的入口
func main() {
	log.Println("Starting work.")

	// 为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

// createTask 返回一个根据 id
// 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

// G:\GitHub\learnGo\Go语言实战\ch07_并发模式\runner\src\runner\main>set GOPATH=G:\GitHub\learnGo\Go语言实战\ch07_并发模式\runner
// G:\GitHub\learnGo\Go语言实战\ch07_并发模式\runner\src\runner\main>go run main.go
// 2019/10/07 18:41:42 Starting work.
// 2019/10/07 18:41:42 Processor - Task #0.
// 2019/10/07 18:41:42 Processor - Task #1.
// 2019/10/07 18:41:43 Processor - Task #2.
// 2019/10/07 18:41:45 Terminating due to timeout.
// exit status 1
