package main

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"fmt"
	"github.com/fastmall/manager-gateway/route"
	"github.com/gin-contrib/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fastmall/manager-gateway/dubbo"
	"github.com/gin-gonic/gin"
)

var (
	BuildTime  string
	GoVersion  string
	GitMessage string
)

func startHttp() {
	r := gin.Default()
	pprof.Register(r)
	route.SetCustomerRoute(r)
	route.SetGoodsRoute(r)
	route.SetCartRoute(r)
	route.SetOrderRoute(r)
	err := r.Run(":80")
	if err != nil {
		logger.Fatal(err)
	}
}

func main() {
	dubbo.StartDubboConsumer()
	go startHttp()
	initSignal()
}

func init() {
	msg := fmt.Sprintf("BuildAt: %s\nBuildBy: %s\nGit：%s", BuildTime, GoVersion, GitMessage)
	fmt.Println(msg)
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.Sleep(time.Second * 2)
			time.AfterFunc(2*time.Second, func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})
			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}
