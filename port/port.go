package port

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"bufio"
	"fmt"
	"github.com/sas/secserver/app/models/asset-scan/mode"
	"github.com/MBZ986/LadonGo/tcp"
	"log"
	"net"
	"os"
	"sync"
	"time"
	//"io/ioutil"
	"strconv"
	"strings"
)

func IsBanner(address string) (string, error) {

	conn, err := net.DialTimeout("tcp", address, time.Second*10)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	tcpconn := conn.(*net.TCPConn)
	tcpconn.SetReadDeadline(time.Now().Add(time.Second * 5))
	reader := bufio.NewReader(conn)
	return reader.ReadString('\n')
}

func CheckPort(ip net.IP, port int) {
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if conn != nil {
		fmt.Println(tcpAddr.IP, tcpAddr.Port, "Open")
		conn.Close()
	}
	if err != nil {
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
		//	fmt.Println(err)
	}
}

func TxtWrite(text string) {
	f, err := os.Create("port.log")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(text))
		//_,err=f.Write(text)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func PortCheck(host string, port int, res mode.Result) (result bool) {
	result = false
	ip := net.ParseIP(host)
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if conn != nil {
		//fmt.Println(tcpAddr.IP, tcpAddr.Port, "Open")

		datamap := map[string]string{"host": host, "port": strconv.Itoa(port)}
		res.Push(datamap)
		//TxtWrite(tcpAddr.IP.String()+"\t"+tcpAddr.Port.String()+"\tOpen")
		//TxtWrite(host+"\t"+strconv.Itoa(port)+" Open")
		conn.Close()
		result = true
	}
	if err != nil {
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
		//	fmt.Println(err)
	}
	return result
}

func PortIsOpen(ip net.IP, port int) (result bool, err error) {
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if conn != nil {
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Open")
		conn.Close()
		result = true
	}
	if err != nil {
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
		//	fmt.Println(err)
	}
	return result, err
}

type Workdist struct {
	Host string
}

const (
	taskload = 255
	tasknum  = 255
)

var wg sync.WaitGroup

func TaskPort(ip string, debugLog *log.Logger) {
	tasks := make(chan Workdist, taskload)
	wg.Add(tasknum)

	for gr := 1; gr <= tasknum; gr++ {
		go workerPort(tasks, debugLog)
	}

	for i := 1; i < 256; i++ {
		host := fmt.Sprintf("%s.%d", ip, i)
		task := Workdist{
			Host: host,
		}
		tasks <- task
	}
	close(tasks)
	wg.Wait()
}

func workerPort(tasks chan Workdist, debugLog *log.Logger) {
	defer wg.Done()
	task, ok := <-tasks
	if !ok {
		return
	}
	host := task.Host

	//Default
	ScanPort(host)

}

var DefaultPorts = []int{21, 22, 23, 25, 80, 443, 8080, 110, 135, 139, 445, 389, 489, 587, 1433, 1434, 1521, 1522, 1723, 2121, 3000, 3306, 3389, 4899, 5631, 5632, 5800, 5900, 7071, 43958, 65500, 4444, 8888, 6789, 4848, 5985, 5986, 8081, 8089, 8443, 10000, 6379, 7001, 7002}

func ScanPort(host string) {
	var wg sync.WaitGroup
	for _, p := range DefaultPorts {
		wg.Add(1)
		//CheckPort(net.ParseIP(host),p)
		tcp.PortCheck(host, p)
		defer wg.Done()
	}
	wg.Wait()
}

func ScanPortBanner(host string,result mode.Result) {
	for _, p := range DefaultPorts {
		tcp.TcpBanner(host, strconv.Itoa(p),result)
	}
}

func ScanPorts(host, ports string) {
	var wg sync.WaitGroup
	for _, port := range strings.Split(ports, ",") {
		wg.Add(1)
		//CheckPort(net.ParseIP(host),p)
		p, err := strconv.Atoi(port)
		if err != nil {
		}
		tcp.PortCheck(host, p)
		defer wg.Done()
	}
	wg.Wait()
}

func ScanPortBanners(host, ports string,result mode.Result) {
	for _, port := range strings.Split(ports, ",") {
		//p, err := strconv.Atoi(port)
		//if err !=nil{
		//}
		//tcp.GetBanner(host,p)
		tcp.TcpBanner(host, port,result)
	}
}

func ScanPortBannerSingle(host, port string,result mode.Result) {
	tcp.TcpBanner(host, port,result)
}

func ScanPortBannerRange(host, ports string,result mode.Result) {
	port := strings.Split(ports, "-")
	p1, _ := strconv.Atoi(port[0])
	p2, _ := strconv.Atoi(port[1])

	for i := p1; i <= p2; i++ {
		tcp.TcpBanner(host, strconv.Itoa(i),result)
	}

}
