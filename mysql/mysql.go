package mysql
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"github.com/MBZ986/LadonGo/port"
	"github.com/MBZ986/LadonGo/dic"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"strings"
)

func MysqlAuth(ip string, port string, user string, pass string) ( result bool,err error) {
	result = false
    db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+":"+port+")/mysql?charset=utf8")
    if err != nil {
    }
	if db.Ping()==nil {
		result = true
	}
	return result,err
}

func MysqlScan2(ScanType string,Target string, result mode.Result) {

	Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			fmt.Println("Check... "+Target+" "+u+" "+p)
			datamap := map[string]string{"flag": "checking", "target": Target, "port": "3306", "user": u, "pass": p}
			result.Push(datamap)
			res,err := MysqlAuth(Target, "3306", u, p)
			if res==true && err==nil {
				//logger.PrintIsok2(ScanType,Target,"3306",u, p)
				datamap["flag"] = "found"
				result.Push(datamap)
				break Loop
			}
		}
	}

}

func MysqlScan(ScanType string,Target string,result mode.Result) {
	if port.PortCheck(Target,3306,result) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				//fmt.Println("Check... "+Target+" "+u+" "+p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "3306", "user": u, "pass": p}
				result.Push(datamap)
				res,err := MysqlAuth(Target, "3306", u, p)
				if res==true && err==nil {
					//logger.PrintIsok2(ScanType,Target,"3306",u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
				
			}
		} else {
			MysqlScan2(ScanType,Target,result)
		}
	}
}