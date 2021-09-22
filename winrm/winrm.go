package winrm

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"github.com/MBZ986/LadonGo/dic"
	"github.com/MBZ986/LadonGo/mode"
	"github.com/masterzen/winrm"
	"os"

	"github.com/MBZ986/LadonGo/port"
	"strings"
)

var help = func() {
	fmt.Println("Winrm Shell by k8gege")
	fmt.Println("====================================================")
	fmt.Println("winrmcmd host port user pass cmd")
}

func WinrmAuth(host, user, pass string, port int) (result bool, err error) {
	result = false
	endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, user, pass)
	if err != nil {
		//panic(err)
	}
	res, err := client.Run("echo ISOK", os.Stdout, os.Stderr)
	if res == 0 {
		result = true
	}
	return result, err
}

func WinrmScan2(ScanType string, Target string, result mode.Result) {
Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			//fmt.Println("Check... " + Target + " " + u + " " + p)
			datamap := map[string]string{"flag": "checking", "target": Target, "port": "5985", "user": u, "pass": p}
			result.Push(datamap)
			res, err := WinrmAuth(Target, u, p, 5985)
			if res == true && err == nil {
				//logger.PrintIsok2(ScanType, Target, "5985", u, p)
				datamap["flag"] = "found"
				result.Push(datamap)
				break Loop
			}
		}
	}

}

func WinrmScan(ScanType string, Target string, result mode.Result) {
	if port.PortCheck(Target, 5985, result) {
		if dic.UserPassIsExist() {
		Loop:
			for _, up := range dic.UserPassDic() {
				s := strings.Split(up, " ")
				u := s[0]
				p := s[1]
				//fmt.Println("Check... "+Target+" "+u+" "+p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "5985", "user": u, "pass": p}
				result.Push(datamap)
				res, err := WinrmAuth(Target, u, p, 5985)
				if res == true && err == nil {
					//logger.PrintIsok2(ScanType,Target,"5985",u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
			}
		} else {
			WinrmScan2(ScanType, Target, result)
		}
	}
}

// func main() {
// args := os.Args
// if len(args) < 5 || args == nil {
// help()
// return
// }
// host := args[1]
// port,err := strconv.Atoi(args[2])
// user := args[3]
// pass := args[4]
// cmd := args[5]

// //WinrmAuth(host,user,pass,port)

// //endpoint := winrm.NewEndpoint("192.168.1.116", 5985, false, false, nil, nil, nil, 0)
// //client, err := winrm.NewClient(endpoint, "k8gege", "k8test520!@#")
// endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)
// client, err := winrm.NewClient(endpoint, user, pass)
// if err != nil {
// panic(err)
// }
// client.Run(cmd, os.Stdout, os.Stderr)
// }
