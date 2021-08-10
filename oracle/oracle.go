package oracle

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"database/sql"
	"github.com/MBZ986/LadonGo/dic"
	"github.com/MBZ986/LadonGo/port"
	_ "github.com/godror/godror"
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"strings"
)

func OracleAuth(host, port, user, pass string) (bool) {
	db, err := sql.Open("godror", user+"/"+pass+"@"+host+":"+port+"/orcl")
	//if err != nil {
	//panic(err)
	//return false
	//}
	if err == nil {
		//defer db.Close()
		err = db.Ping()
		if err == nil {
			return true
		}
	}

	return false
}

func OracleScan2(ScanType string, Target string, result mode.Result) {
	if port.PortCheck(Target, 1521, result) {
	Loop:
		for _, u := range dic.UserDic() {
			for _, p := range dic.PassDic() {
				//fmt.Println("Check... " + Target + " " + u + " " + p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "1521", "user": u, "pass": p}
				result.Push(datamap)
				res := OracleAuth(Target, "1521", u, p)
				if res {
					//logger.PrintIsok2(ScanType, Target, "1521", u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}
			}
		}
	}
}

func OracleScan(ScanType string, Target string, result mode.Result) {
	if port.PortCheck(Target, 1521, result) {
		if dic.UserPassIsExist() {
		Loop:
			for _, up := range dic.UserPassDic() {
				s := strings.Split(up, " ")
				u := s[0]
				p := s[1]
				//fmt.Println("Check... " + Target + " " + u + " " + p)
				datamap := map[string]string{"flag": "checking", "target": Target, "port": "1521", "user": u, "pass": p}
				result.Push(datamap)
				res := OracleAuth(Target, "1521", u, p)
				if res {
					//logger.PrintIsok2(ScanType, Target, "1521", u, p)
					datamap["flag"] = "found"
					result.Push(datamap)
					break Loop
				}

			}
		} else {
			OracleScan2(ScanType, Target, result)
		}
	}
}
