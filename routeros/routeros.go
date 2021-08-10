package routeros
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"strings"
	"github.com/go-routeros/routeros"
	"github.com/MBZ986/LadonGo/port"
	"github.com/MBZ986/LadonGo/dic"
)

func RouterOSAuth(ip string, port string, user string, pass string) ( result bool,err error) {
	result = false
    _, err = routeros.Dial(ip+":"+port,user,pass)
    if err == nil {
	result = true
    }

	//defer c.Close()

	return result,err
}

func RouterOSScan2(ScanType string,Target string,result mode.Result) {
	Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			//fmt.Println("Check... "+Target+" "+u+" "+p)
			datamap := map[string]string{"flag": "checking", "target": Target, "port": "8728", "user": u, "pass": p}
			result.Push(datamap)
			res,err := RouterOSAuth(Target, "8728", u, p)
			if res==true && err==nil {
				//logger.PrintIsok2(ScanType,Target,"8728",u, p)
				datamap["flag"] = "found"
				result.Push(datamap)
				break Loop
			}
		}
	}
}

func RouterOSScan(ScanType string,Target string,result mode.Result) {
	if port.PortCheck(Target,8728,result) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "8728", "user": u, "pass": p}
				result.Push(datamap)
				res,err := RouterOSAuth(Target, "8728", u, p)
				if res==true && err==nil {
					//logger.PrintIsok2(ScanType,Target,"8728",u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
				
			}
		} else {
			RouterOSScan2(ScanType,Target,result)
		}
	}
}


