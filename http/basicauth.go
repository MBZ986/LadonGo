package http

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/MBZ986/LadonGo/dic"
	"github.com/MBZ986/LadonGo/logger"
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"net/http"
	"strings"
)

func IsBasicAuthURL(url string, res mode.Result) (result bool) {
	result = false
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("", "")
	resp, err := client.Do(req)
	if err == nil {
		resp.Body.Close()
		if resp.StatusCode == 401 {
			//fmt.Println(url+" IS 401URL")
			datamap := map[string]string{"url": url}
			res.Push(datamap)
			result = true
		}
	}
	return result
}

func BasicAuth(ScanType, url, user, pass string) (result bool, err error) {
	result = false
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err == nil {
		resp.Body.Close()
		if resp.StatusCode != 401 {
			//logger.PrintIsok(ScanType,url,user, pass)
			return true, err
		} else {
			//fmt.Println("err")
		}
	}
	return result, err
}

func BasicAuthScan2(ScanType string, Target string, result mode.Result) {
	if IsBasicAuthURL(Target, result) {
	Loop:
		for _, u := range dic.UserDic() {
			for _, p := range dic.PassDic() {
				//fmt.Println("Check... " + Target + " " + u + " " + p)
				datamap := map[string]string{"flag": "checking", "target": Target, "user": u, "pass": p}
				result.Push(datamap)
				res, err := BasicAuth(ScanType, Target, u, p)
				if res == true && err == nil {
					logger.PrintIsok(ScanType, Target, u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
			}
		}
	}
}

func BasicAuthScan(ScanType string, Target string, result mode.Result) {
	if IsBasicAuthURL(Target, result) {
		if dic.UserPassIsExist() {
		Loop:
			for _, up := range dic.UserPassDic() {
				s := strings.Split(up, " ")
				u := s[0]
				p := s[1]
				//fmt.Println("Check... " + Target + " " + u + " " + p)
				datamap := map[string]string{"flag": "checking", "target": Target, "user": u, "pass": p}
				result.Push(datamap)
				res, err := BasicAuth(ScanType, Target, u, p)
				if res == true && err == nil {
					//logger.PrintIsok(ScanType, Target, u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
			}

		} else {
			BasicAuthScan2(ScanType, Target, result)
		}
	}
}
