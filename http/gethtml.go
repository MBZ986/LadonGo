package http

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"bytes"
	"github.com/MBZ986/LadonGo/mode"
	"github.com/k8gege/LadonGo/str"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetHtml(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		//panic(err)
		return ""
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			//panic(err)
			return ""
		}
	}
	return result.String()
}

func GetTitle(html string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	html = strings.Replace(html, "\n", "", -1)
	title := strings.Trim(str.GetBetween(html, "<title>", "</title>"), " ")
	return title
}

func ScanTitle(host string, result mode.Result) {
	if strings.Contains(host, ":") {
		title := GetTitle(GetHtml(host))
		if title != "" {
			//fmt.Println(host+"\t"+title)
			result.Push(map[string]string{host: title})
		}
	} else {
		url := "http://" + host
		title := GetTitle(GetHtml(url))
		if title != "" {
			//fmt.Println(url+"\t"+title)
			result.Push(map[string]string{host: title})
		}
		url = "https://" + host
		title = GetTitle(GetHtml(url))
		if title != "" {
			//fmt.Println(url+"\t"+title)
			result.Push(map[string]string{host: title})
		}
	}
}

// func main() {
// fmt.Println(GetTitle(GetHtml("http://k8gege.org")))
// }
