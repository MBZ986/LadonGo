package mssql

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"database/sql"
	"fmt"
	"github.com/MBZ986/LadonGo/dic"
	"github.com/MBZ986/LadonGo/mode"
	"github.com/MBZ986/LadonGo/port"
	_ "github.com/denisenkom/go-mssqldb"
	"strings"
)

func MssqlAuth(ip, port, user, pass string) (result bool, err error) {
	result = false
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;encrypt=disable", ip, user, pass, port)
	db, err := sql.Open("mssql", connString)
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result = true
		}
	}

	return result, err
}

func MssqlScan2(ScanType string, Target string, result mode.Result) {

Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			//fmt.Println("Check... " + Target + " " + u + " " + p)
			datamap := map[string]string{"flag": "checking", "target": Target, "port": "1433", "user": u, "pass": p}
			result.Push(datamap)
			res, err := MssqlAuth(Target, "1433", u, p)
			if res == true && err == nil {
				//logger.PrintIsok2(ScanType, Target, "1433", u, p)
				datamap["flag"] = "found"
				result.Push(datamap)
				break Loop
			}
		}
	}

}

func MssqlScan(ScanType string, Target string, result mode.Result) {
	if port.PortCheck(Target, 1433, result) {
		if dic.UserPassIsExist() {
		Loop:
			for _, up := range dic.UserPassDic() {
				s := strings.Split(up, " ")
				u := s[0]
				p := s[1]
				//fmt.Println("Check... "+Target+" "+u+" "+p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "1433", "user": u, "pass": p}
				result.Push(datamap)
				res, err := MssqlAuth(Target, "1433", u, p)
				if res == true && err == nil {
					//logger.PrintIsok2(ScanType,Target,"1433",u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
			}
		} else {
			MssqlScan2(ScanType, Target, result)
		}
	}
}
