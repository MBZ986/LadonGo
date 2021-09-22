package smb

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/MBZ986/LadonGo/dic"
	"github.com/MBZ986/LadonGo/logger"
	"github.com/MBZ986/LadonGo/mode"
	"github.com/MBZ986/LadonGo/port"
	"github.com/stacktitan/smb/smb"
	"strings"
)

//Not Support 2003
func SmbAuth(ip string, port string, username string, password string) (result bool, err error) {
	result = false

	options := smb.Options{
		Host:        ip,
		Port:        445,
		User:        username,
		Password:    password,
		Domain:      "",
		Workstation: "",
	}

	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			result = true
		}
	}
	return result, err
}

func SmbScan2(ScanType string, Target string, result mode.Result) {

Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			//fmt.Println("Check... " + Target + " " + u + " " + p)
			datamap := map[string]string{"flag": "checking", "target": Target, "user": u, "port": p}
			result.Push(datamap)
			res, err := SmbAuth(Target, "445", u, p)
			if res == true && err == nil {
				logger.PrintIsok(ScanType, Target, u, p)
				datamap["flag"] = "ok"
				result.Push(datamap)
				break Loop
			}
		}
	}

}

func SmbScan(ScanType string, Target string, result mode.Result) {
	if port.PortCheck(Target, 445, result) {
		if dic.UserPassIsExist() {
		Loop:
			for _, up := range dic.UserPassDic() {
				s := strings.Split(up, " ")
				u := s[0]
				p := s[1]
				//fmt.Println("Check... " + Target + " " + u + " " + p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "445", "user": u, "pass": p}
				result.Push(datamap)
				res, err := SmbAuth(Target, "445", u, p)
				if res == true && err == nil {
					//logger.PrintIsok(ScanType, Target, u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
			}
		} else {
			SmbScan2(ScanType, Target, result)
		}
	}
}
