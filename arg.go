package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

var parser = argparse.NewParser(
	"sfwg",
	"segfault 上的 openvpn 管理程序",
)

var isHelp = parser.Flag("h", "help", &argparse.Options{
	Required: false,
	Help:     "显示帮助信息",
})

var isShowStatus = parser.Flag("s", "status", &argparse.Options{
	Required: false,
	Help:     "查看状态",
})

var isOvpnDown = parser.Flag("d", "down", &argparse.Options{
	Required: false,
	Help:     "关闭连接",
})

var configFile = parser.String("", "use", &argparse.Options{
	Required: false,
	Help:     "指定配置文件",
	Default:  "openvpn.conf",
})

var isView = parser.Flag("", "view", &argparse.Options{
	Required: false,
	Help:     "查看配置文件内容",
})

var isOvpnUp = parser.Flag("", "up", &argparse.Options{
	Required: false,
	Help:     "开启连接",
})

var upWithUsername = parser.String("u", "user", &argparse.Options{
	Required: false,
	Help:     "指定用户名",
})

var upWithPassword = parser.String("p", "password", &argparse.Options{
	Required: false,
	Help:     "指定密码",
})

var upWithRoutes = parser.StringList("r", "route", &argparse.Options{
	Required: false,
	Help:     "指定路由",
})

func doParse() {
	parser.DisableHelp()

	if err := parser.Parse(os.Args); err != nil {
		fmt.Println(parser.Usage(err))
	}

	if len(os.Args) == 1 {
		fmt.Println(parser.Usage(nil))
		return
	}

	if *isHelp {
		fmt.Println(parser.Usage(nil))
		doOvpnHelp()
	}

	if *isShowStatus {
		doOvpnStatus()
	}

	if *isOvpnDown {
		doOvpnDown()
	}

	if *isView {
		doView(*configFile)
	}

	if *isOvpnUp {
		doOvpnUp(*configFile, *upWithUsername, *upWithPassword, *upWithRoutes)
	}
}
