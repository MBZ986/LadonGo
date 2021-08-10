package http

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"net/http"
	//"os"
	"strings"
)

func IsUrl(url string) (result string) {
	if !strings.Contains(url, "http") {
		url := "http://" + url
		return url
	}

	return url
}
func HttpBanner(url string, res mode.Result) (result bool, err error) {

	url2 := IsUrl(url)
	response, err := http.Head(url2)
	if err != nil {
		//fmt.Println(err.Error())
		//os.Exit(2)
		return false, err
	}

	//fmt.Println(response)
	//fmt.Println(response.Status)
	for k, v := range response.Header {
		//fmt.Println(k+":", v)
		if k == "Server" {
			//fmt.Println(url2, v)
			//data := map[string]interface{}{url2: v}
			res.Push(fmt.Sprintf("HttpBanner:%s\t%s",url2,v))
		}
	}

	return result, err

}

// func main() {

// HttpBanner("http://www.baidu.com")
// HttpBanner("192.168.1.1")

// }
