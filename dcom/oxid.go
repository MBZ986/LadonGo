package dcom

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"net"
	"strings"
	"time"
)

func OxidInfo(host string, result mode.Result) ([]string, error) {
	timeout := 3000 * time.Millisecond
	d := net.Dialer{Timeout: timeout}
	tcpcon, err := d.Dial("tcp", host+":135")
	if err != nil {
		return nil, err
	}
	//defer tcpcon.Close()
	sendData := "\x05\x00\x0b\x03\x10\x00\x00\x00\x48\x00\x00\x00\x01\x00\x00\x00\xb8\x10\xb8\x10\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x01\x00\xc4\xfe\xfc\x99\x60\x52\x1b\x10\xbb\xcb\x00\xaa\x00\x21\x34\x7a\x00\x00\x00\x00\x04\x5d\x88\x8a\xeb\x1c\xc9\x11\x9f\xe8\x08\x00\x2b\x10\x48\x60\x02\x00\x00\x00"
	n, err := tcpcon.Write([]byte(sendData))
	if err != nil {
		return nil, err
	}
	recvData := make([]byte, 4096)
	readTimeout := 3 * time.Second
	err = tcpcon.SetReadDeadline(time.Now().Add(readTimeout))
	n, err = tcpcon.Read(recvData)
	if err != nil {
		return nil, err
	}
	sendData2 := "\x05\x00\x00\x03\x10\x00\x00\x00\x18\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x05\x00"
	n, err = tcpcon.Write([]byte(sendData2))
	if err != nil {
		return nil, err
	}
	err = tcpcon.SetReadDeadline(time.Now().Add(readTimeout))
	n, err = tcpcon.Read(recvData)
	if err != nil {
		return nil, err
	}
	recvStr := string(recvData[:n])
	if len(recvStr) > 42 {
		recvStr_v2 := recvStr[42:]
		packet_v2_end := strings.Index(recvStr_v2, "\x09\x00\xff\xff\x00\x00")
		packet_v2 := recvStr_v2[:packet_v2_end]
		hostname_list := strings.Split(packet_v2, "\x00\x00")
		var datastr string
		if len(hostname_list) > 1 {
			for _, value := range hostname_list {
				if strings.Trim(value, " ") != "" {
					datastr += replaceSpace(value) + "\t"
					//fmt.Println(replace)
				}
			}
			//fmt.Println(datastr)
			result.Push(datastr)
			return hostname_list, nil
		}
	}
	return nil, nil

}
func replaceSpace(str string) string{
	runes := []rune(str)
	newstr := ""
	for _, s :=range runes{
		//fmt.Println(string(s))
		//fmt.Println(s)
		if s!=0{
			newstr+=string(s)
		}
	}
	return newstr
}
//Result
// WIN-788
// 192.168.1.30
// 192.168.30.10

// DESKTOP-K8gege
// 192.168.1.3
//192.168.1.3
// 2001:0:3876:fb58:2443:a003:4004:2992
