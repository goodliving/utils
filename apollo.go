package utils

import (
	"fmt"
	"github.com/shima-park/agollo"
	"os"
)

// SetupApollo 默认配置apollo
func SetupApollo() {
	err := agollo.InitWithDefaultConfigFile(
		agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))), // 打印日志信息
		agollo.AutoFetchOnCacheMiss(), // 在配置未找到时，去apollo的带缓存的获取配置接口，获取配置
		agollo.FailTolerantOnBackupExists(), // 在连接apollo失败时，如果在配置的目录下存在.agollo备份配置，会读取备份在服务器无法连接的情况下
	)

	if err != nil {
		panic(err)
	}

	// 如果想监听并同步服务器配置变化，启动apollo长轮训
	// 返回一个期间发生错误的error channel,按照需要去处理
	errorCh := agollo.Start()

	// 监听apollo配置更改事件
	// 返回namespace和其变化前后的配置,以及可能出现的error
	watchCh := agollo.Watch()

	//stop := make(chan bool)
	//appNSCh := agollo.WatchNamespace("application", stop)
	go func() {
		for {
			select {
			case err := <-errorCh:
				fmt.Println("Error:", err)
			case resp := <-watchCh:
				fmt.Println("Watch Apollo:", resp)
				//case resp := <-appNSCh:
				//	fmt.Println("Watch Namespace", "application", resp)
				//case <-time.After(time.Second):
				//	fmt.Println("timeout:", agollo.Get("timeout"))
			}
		}
	}()
}