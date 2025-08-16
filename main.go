package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	doParse()
}

func isExist(fn string) bool {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return false
	}
	return true
}

func doOvpnHelp() {
	runCmd([]string{"curl", "sf/ovpn"})
}

func doOvpnStatus() {
	runCmd([]string{"curl", "sf/ovpn/show"})
}

func doOvpnDown() {
	runCmd([]string{"curl", "sf/ovpn/down"})
}

func doView(fn string) {
	if !isExist(fn) {
		fmt.Println("配置文件不存在")
		return
	}

	content, err := os.ReadFile(fn)
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return
	}

	fmt.Println(string(content))
}

func doOvpnUp(fn string, user string, pass string, routes []string) {
	if !isExist(fn) {
		fmt.Println("配置文件不存在")
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败", err)
	}

	command := []string{
		"curl", "sf/ovpn/up",
		"-d", fmt.Sprintf("config=%s", filepath.Join(cwd, fn)),
	}

	if user != "" && pass != "" {
		command = append(command, []string{
			"-d", fmt.Sprintf("user=%s", user),
			"-d", fmt.Sprintf("pass=%s", pass),
		}...)
	}

	if user == "" && pass != "" {
		command = append(command, []string{
			"-d", fmt.Sprintf("keypass=%s", pass),
		}...)
	}

	if len(routes) > 0 {
		for _, route := range routes {
			command = append(command, []string{
				"-d", fmt.Sprintf("route=%s", route),
			}...)
		}
	}

	runCmd(command)
}
