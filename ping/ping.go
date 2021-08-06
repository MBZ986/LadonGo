package ping

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"bytes"
	"github.com/MBZ986/LadonGo/mode"
	"os/exec"
	//"log"
	"fmt"
	"runtime"
	"strings"
)

//Support User
func CmdPing(host string) (result bool, err error) {
	sysType := runtime.GOOS
	if sysType == "linux" {
		cmd := exec.Command("/bin/sh", "-c", "ping -c 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		//err := cmd.Run()
		//if err != nil {
		//log.Fatal(err)
		//}
		//fmt.Println(out.String())
		if strings.Contains(out.String(), "ttl=") {
			//fmt.Println("ISOK")
			result = true
		}
	} else if sysType == "windows" {
		cmd := exec.Command("cmd", "/c", "ping -a -n 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		//err := cmd.Run()
		//if err != nil {
		//log.Fatal(err)
		//}
		//fmt.Println(out.String())
		if strings.Contains(out.String(), "TTL=") {
			//fmt.Println("ISOK")
			result = true
		}
	}
	return result, err
}

func PingOK(host string) (result bool) {
	sysType := runtime.GOOS
	if sysType == "linux" {
		cmd := exec.Command("/bin/sh", "-c", "ping -c 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		if strings.Contains(out.String(), "ttl=") {
			//fmt.Println("ISOK")
			result = true
		}
	} else if sysType == "windows" {
		cmd := exec.Command("cmd", "/c", "ping -n 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		if strings.Contains(out.String(), "TTL=") {
			//fmt.Println("ISOK")
			result = true
		}
	}
	return result
}

func PingName(host string, res mode.Result) (result bool, err error) {
	sysType := runtime.GOOS
	if sysType == "linux" {
		cmd := exec.Command("/bin/sh", "-c", "ping -c 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		//err := cmd.Run()
		//if err != nil {
		//log.Fatal(err)
		//}
		//fmt.Println(out.String())
		if strings.Contains(out.String(), "ttl=") {
			fmt.Println("PING: " + host)
			result = true
			res.Push(host)
		}
	} else if sysType == "windows" {
		cmd := exec.Command("cmd", "/c", "ping -a -n 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		//err := cmd.Run()
		//if err != nil {
		//log.Fatal(err)
		//}

		if strings.Contains(out.String(), "TTL=") && strings.Contains(out.String(), "["+host+"]") {
			//fmt.Println("PING: " + host + "\t" + str.GetBetween(out.String(), "Pinging", " ["+host+"]"))
			result = true
			res.Push(host)
		} else if strings.Contains(out.String(), "TTL=") {
			//fmt.Println("PING: " + host)
			res.Push(host)
		}
	}
	return result, err
}
