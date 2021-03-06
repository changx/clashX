package main // import "github.com/changx/clashX/ClashX"

import (
	"C"

	"github.com/changx/clash/config"
	"github.com/changx/clash/constant"
	"github.com/changx/clash/hub"
	"github.com/changx/clash/proxy"
	"github.com/changx/clash/tunnel"
)
import (
	"os"
	"os/signal"
	"syscall"
)

//export run
func run() *C.char {
	tunnel.Instance().Run()
	proxy.Instance().Run()
	hub.Run()

	config.Init()
	err := config.Instance().Parse()
	if err != nil {
		return C.CString(err.Error())

	}

	return C.CString("success")
}

//export updateAllConfig
func updateAllConfig() *C.char {
	err := config.Instance().Parse()
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString("")
}

//export setConfigHomeDir
func setConfigHomeDir(root string) {
	constant.SetHomeDir(root)
}
func main() {
	run()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
