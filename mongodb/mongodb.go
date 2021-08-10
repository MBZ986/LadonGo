package mgo
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
	"strings"
	"github.com/MBZ986/LadonGo/port"
	"github.com/MBZ986/LadonGo/dic"
	"github.com/MBZ986/LadonGo/logger"
)

func MongoAuth(ip string, port string, username string, password string) ( result bool,err error) {
	session, err := mgo.DialWithTimeout("mongodb://"+username+":"+password+"@"+ip+":"+port+"/"+"admin", time.Second*3)
	if err == nil && session.Ping() == nil {
		defer session.Close()
		if err == nil && session.Run("serverStatus", nil) == nil {
			result = true
		}
	}
	return result,err
}

func MongoUnAuth(ip string, port string) (result bool,err error) {
	session, err := mgo.Dial(ip + ":" + port)
	if err == nil && session.Run("serverStatus", nil) == nil {
		result = true
	}
	return result,err
}

func MongoScan2(ScanType string,Target string,result mode.Result) {
	Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			//fmt.Println("Check... "+Target+" "+u+" "+p)
			datamap := map[string]string{"flag": "checking", "target": Target, "port": "27017", "user": u, "pass": p}
			result.Push(datamap)
			res,err := MongoAuth(Target, "27017", u, p)
			if res==true && err==nil {
				//logger.PrintIsok(ScanType,Target,u, p)
				datamap["flag"] = "found"
				result.Push(datamap)
				break Loop
			}
		}
	}

}

func MongoScan(ScanType string,Target string,result mode.Result) {
	if port.PortCheck(Target,27017,result) {
		//if dic.PwdIsExist()==false {
		fmt.Println("Check... "+Target)
		res1,_ := MongoUnAuth(Target, "27017")
		if res1 {
			logger.PrintIsok0(ScanType,Target,"27017")
		}
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				//fmt.Println("Check... "+Target+" "+u+" "+p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "27017", "user": u, "pass": p}
				result.Push(datamap)
				res,err := MongoAuth(Target, "27017", u, p)
				if res==true && err==nil {
					//logger.PrintIsok(ScanType,Target,u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
				
			}
		} else {
			MongoScan2(ScanType,Target,result)
		}
	}
}