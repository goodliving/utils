package utils

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shima-park/agollo"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

type RpcxInfo struct {
	AppName string
	ServiceAddr string
	ConsulAddr string
	RpcxBasePath string
}

const (
	FlagAppName = "app.name"
	FlagRpcxPort = "rpcx.port"
	FlagRpcxConsulAddr = "rpcx.consul.addr"
	FlagRpcxBasePath = "rpcx.base.path"
)

func GetRpcxInfo() RpcxInfo {
	appName := agollo.Get(FlagAppName)
	serviceAddr := fmt.Sprintf("%s:%s", GetHostIP(), agollo.Get(FlagRpcxPort))
	consulAddr := agollo.Get(FlagRpcxConsulAddr)
	basePath := agollo.Get(FlagRpcxBasePath)
	return RpcxInfo{
		AppName: appName,
		ServiceAddr: serviceAddr,
		ConsulAddr:  consulAddr,
		RpcxBasePath: basePath,
	}
}

func AddConsulRegistryPlugin(s *server.Server, basePath, addr, consulAddr string) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ConsulServers:  []string{consulAddr},
		BasePath:       basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}

	err := r.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("rpcx服务启动失败")
	}

	s.Plugins.Add(r)
	log.Info().Dict("rpcxInfo", zerolog.Dict().
		Str("basePath", basePath).Str("addr", addr).Str("consulAddr", consulAddr)).Msg("加载配置")

}

