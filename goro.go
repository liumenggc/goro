//goro 创建两个goroutine, 以并发的方式分别显示大写和小写的英文字母。
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	//分配一个逻辑处理器给调度器使用。
	runtime.GOMAXPROCS(1)

	//wg用来等待程序完成。
	//计数2,表示等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutine")

	go func() {
		//函数退出时，试调用Done()，通知工作已经完成
		defer wg.Done()

		//
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		//函数退出时，试调用Done()，通知工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++  {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	//等待Goroutine结束
	fmt.Println("Waiting a Finsh")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}