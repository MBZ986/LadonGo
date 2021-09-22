package ladon

//Ladon Scanner for golang
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"github.com/MBZ986/LadonGo/dcom"
	"github.com/MBZ986/LadonGo/dic"
	"github.com/MBZ986/LadonGo/exp"
	"github.com/MBZ986/LadonGo/ftp"
	"github.com/MBZ986/LadonGo/http"
	"github.com/MBZ986/LadonGo/icmp"
	"github.com/MBZ986/LadonGo/mode"
	"github.com/MBZ986/LadonGo/mongodb"
	"github.com/MBZ986/LadonGo/mssql"
	"github.com/MBZ986/LadonGo/mysql"
	"github.com/MBZ986/LadonGo/nbt"
	//"github.com/MBZ986/LadonGo/oracle"
	"github.com/MBZ986/LadonGo/ping"
	"github.com/MBZ986/LadonGo/port"
	"github.com/MBZ986/LadonGo/rexec"
	"github.com/MBZ986/LadonGo/smb"
	"github.com/MBZ986/LadonGo/snmp"
	"github.com/MBZ986/LadonGo/ssh"
	"github.com/MBZ986/LadonGo/t3"
	//"github.com/k8gege/LadonGo/worker"
	//"github.com/k8gege/LadonGo/color" //Only Windows
	"github.com/MBZ986/LadonGo/vul"
	"github.com/MBZ986/LadonGo/winrm"
	//"github.com/k8gege/LadonGo/tcp"
	"github.com/MBZ986/LadonGo/redis"
	"github.com/MBZ986/LadonGo/routeros"
	"github.com/fatih/color"
	"log"
	"net"
	"os"
	"os/user"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func help() {
	//if runtime.GOOS=="windows" {fmt.Println("\nHelp:")
	//} else{fmt.Println("\033[32m\nHelp:\033[0m")}
	color.Green("\nHelp:")
	s := ""
	if runtime.GOOS != "windows" {
		s = "./"
	}
	//fmt.Println(s+"Ladon Help")
	fmt.Println(s + "Ladon FuncList")
	fmt.Println(s + "Ladon Detection")
	fmt.Println(s + "Ladon VulDetection")
	fmt.Println(s + "Ladon BruteFor")
	fmt.Println(s + "Ladon RemoteExec")
	fmt.Println(s + "Ladon Exploit")
	fmt.Println(s + "Ladon Example")

}

func FuncList() {
	//help()
	Detection()
	VulDetection()
	BruteFor()
	RemoteExec()
	Exploit()
	Example()
}

func Example() {
	//if runtime.GOOS=="windows" {fmt.Println("\nExample:")
	//} else{fmt.Println("\033[32m\nExample:\033[0m")}
	color.Green("\nExample:")
	s := ""
	if runtime.GOOS != "windows" {
		s = "./"
	}
	fmt.Println(s + "Ladon 192.168.1.8/24 MS17010")
	fmt.Println(s + "Ladon 192.168.1/c MS17010")
	fmt.Println(s + "Ladon 192.168/b MS17010")
	fmt.Println(s + "Ladon 192/a MS17010")
	fmt.Println(s + "Ladon 192.168.1-192.168.5 MS17010")
	fmt.Println(s + "Ladon http://192.168.1.8:8080 BasicAuthScan")
	fmt.Println(s + "Ladon ip.txt MS17010")
	fmt.Println(s + "Ladon url.txt HttpBanner")
	fmt.Println("")
}

func Detection() {
	//if runtime.GOOS=="windows" {fmt.Println("\nDetection:")
	//} else{fmt.Println("\033[33m\nDetection:\033[0m")}
	color.Blue("\nDetection:")
	fmt.Println("PingScan\t(Using system ping to detect Online hosts)")
	fmt.Println("IcmpScan\t(Using ICMP Protocol to detect Online hosts)")
	fmt.Println("SnmpScan\t(Using Snmp Protocol to detect Online hosts)")
	fmt.Println("HttpBanner\t(Using HTTP Protocol Scan Web Banner)")
	fmt.Println("HttpTitle\t(Using HTTP protocol Scan Web titles)")
	fmt.Println("T3Scan  \t(Using T3 Protocol Scan Weblogic hosts)")
	fmt.Println("PortScan\t(Scan hosts open ports using TCP protocol)")
	fmt.Println("TcpBanner\t(Scan hosts open ports using TCP protocol)")
	fmt.Println("OxidScan \t(Using dcom Protocol enumeration network interfaces)")
	fmt.Println("NbtInfo\t(Scan hosts open ports using NBT protocol)")
}

func VulDetection() {
	//if runtime.GOOS=="windows" {fmt.Println("\nVulDetection:")
	//} else{fmt.Println("\033[33m\nVulDetection:\033[0m")}
	color.Yellow("\nVulDetection:")
	fmt.Println("MS17010 \t(Using SMB Protocol to detect MS17010 hosts)")
	fmt.Println("SmbGhost\t(Using SMB Protocol to detect SmbGhost hosts)")
	fmt.Println("CVE-2021-21972\t(Check VMware vCenter 6.5 6.7 7.0 Rce Vul)")
	fmt.Println("CVE-2021-26855\t(Check CVE-2021-26855 Microsoft Exchange SSRF)")
}

func BruteFor() {
	//if runtime.GOOS=="windows" {fmt.Println("\nBruteForce:")
	//} else{fmt.Println("\033[35m\nBruteForce:\033[0m")}
	color.Red("\nBruteForce:")
	fmt.Println("SmbScan \t(Using SMB Protocol to Brute-For 445 Port)")
	fmt.Println("SshScan \t(Using SSH Protocol to Brute-For 22 Port)")
	fmt.Println("FtpScan \t(Using FTP Protocol to Brute-For 21 Port)")
	fmt.Println("401Scan \t(Using HTTP BasicAuth to Brute-For web Port)")
	fmt.Println("MysqlScan \t(Using Mysql Protocol to Brute-For 3306 Port)")
	fmt.Println("MssqlScan \t(Using Mssql Protocol to Brute-For 1433 Port)")
	fmt.Println("OracleScan \t(Using Oracle Protocol to Brute-For 1521 Port)")
	fmt.Println("MongodbScan \t(Using Mongodb Protocol to Brute-For 27017 Port)")
	fmt.Println("WinrmScan \t(Using Winrm Protocol to Brute-For 5985 Port)")
	fmt.Println("SqlplusScan \t(Using Oracle Sqlplus Brute-For 1521 Port)")
	fmt.Println("RouterOSScan \t(Using RouterOS API Brute-For 8728 Port)")
}

func RemoteExec() {
	//if runtime.GOOS=="windows" {fmt.Println("\nRemoteExec:")
	//} else{fmt.Println("\033[35m\nRemoteExec:\033[0m")}
	color.Magenta("\nRemoteExec:")
	fmt.Println("SshCmd   \t(SSH Remote command execution Default 22 Port)")
	fmt.Println("WinrmCmd \t(Winrm Remote command execution Default 5985 Port)")
	fmt.Println("PhpShell \t(PHP Shell Remote command execution Default 80 Port)")
}

func Exploit() {
	//if runtime.GOOS=="windows" {fmt.Println("\nExploit:")
	//} else{fmt.Println("\033[35m\nExploit:\033[0m")}
	color.Magenta("\nExploit:")
	fmt.Println("PhpStudyDoor\t(PhpStudy 2016 & 2018 BackDoor Exploit)")
	fmt.Println("CVE-2018-14847\t(Export RouterOS Password 6.29 to 6.42)")
}

func Noping() {
	s := ""
	if runtime.GOOS != "windows" {
		s = "./"
	}
	fmt.Println("\nNoping Example:")
	fmt.Println(s + "Ladon noping 192.168.1.8 PortScan")
	fmt.Println(s + "Ladon noping ip.txt PortScan")
}

var isicmp bool
var ver = "3.8"

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func GetUser() {
	u, err := user.Current()
	if err != nil {
		//fmt.Println(err)
		return
	}
	if isicmp {
		color.Magenta("User: " + u.Username + " IsAdmin")
	} else {
		fmt.Println("User: " + u.Username + " IsUser")
	}

}

var debugLog *log.Logger
var scanports string

func Run(result mode.Result, param ...string) {
	os.Args = param
	fmt.Println(os.Args)
	color.Yellow("LadonGo " + ver + " by k8gege")
	fmt.Println("Arch: " + runtime.GOARCH + " OS: " + runtime.GOOS)
	if icmp.IcmpOK("localhost") {
		isicmp = true
	}
	GetUser()
	fmt.Println("Pid: ", os.Getpid(), "Process:", path.Base(os.Args[0]))
	scanports = ""
	ParLen := len(os.Args)
	if ParLen == 1 {
		help()
		os.Exit(0)
	}

	// if ParLen>1 {
	// SecPar := os.Args[1]
	// fmt.Println("Load "+SecPar)
	// }

	if ParLen == 2 {
		SecPar := strings.ToUpper(os.Args[1])
		fmt.Println("Load " + SecPar)
		if SecPar == "HELP" || SecPar == "/HELP" || SecPar == "-H" || SecPar == "-H" {
			help()
			os.Exit(0)
		}
		if SecPar == "HELPLIST" || SecPar == "FUNCLIST" {
			FuncList()
			os.Exit(0)
		}
		if SecPar == "NOPING" {
			Noping()
			os.Exit(0)
		}
		if SecPar == "BRUTEFOR" || SecPar == "BRUTE" || SecPar == "BRUTEFORCE" || SecPar == "BRUTE-FORCE" {
			BruteFor()
			os.Exit(0)
		}
		if SecPar == "DETECTION" || SecPar == "DETECT" || SecPar == "DISCOVER" {
			Detection()
			os.Exit(0)
		}
		if SecPar == "VULDETECTION" || SecPar == "VULNERABLE" || SecPar == "POC" || SecPar == "VUL" {
			VulDetection()
			os.Exit(0)
		}
		if SecPar == "REMOTEEXEC" {
			RemoteExec()
			os.Exit(0)
		}
		if SecPar == "EXPLOIT" || SecPar == "EXP" || SecPar == "RCE" {
			Exploit()
			os.Exit(0)
		}
		if SecPar == "EXAMPLE" || SecPar == "USAGE" {
			Example()
			os.Exit(0)
		}
		if SecPar == "WINRMCMD" || SecPar == "WINRMEXEC" || SecPar == "SSHSHELL" {
			rexec.WinrmHelp()
			os.Exit(0)
		}
		if SecPar == "SSHCMD" || SecPar == "SSHEXEC" || SecPar == "SSHSHELL" {
			ssh.SshHelp()
			os.Exit(0)
		}
		if SecPar == "PHPSTUDYDOOR" || SecPar == "PHPSTUDYBACKDOOR" || SecPar == "PHPSTUDYRCE" || SecPar == "PHPSTUDYEXP" {
			exp.PhpStudyDoorHelp()
			os.Exit(0)
		}
		if SecPar == "CVE-2018-14847" {
			exp.Cve2018_14847Help()
			os.Exit(0)
		}
		if SecPar == "PHPSHELL" || SecPar == "PHPWEBSHELL" {
			rexec.PhpShellHelp()
			os.Exit(0)
		}
		fmt.Println(SecPar, "Moudle Not Found")
		os.Exit(0)
	}

	EndPar := os.Args[ParLen-1]
	Target := os.Args[ParLen-2]

	if ParLen == 5 {
		SecPar := strings.ToUpper(os.Args[1])
		if SecPar == "PHPSHELL" || SecPar == "PHPWEBSHELL" {
			fmt.Println("Load " + SecPar)
			rexec.PhpShellExec(os.Args[2], os.Args[3], os.Args[4])
			os.Exit(0)
		}
	} else if ParLen == 4 {
		ThirdPar := strings.ToUpper(os.Args[2])
		if ThirdPar == "PORTSCAN" || ThirdPar == "SCANPORT" || ThirdPar == "TCPSCAN" || ThirdPar == "TCPBANNER" || ThirdPar == "PORTSCANBNNER" || ThirdPar == "SCANPORTBANNER" {
			fmt.Println("Load " + ThirdPar)
			EndPar = ThirdPar
			Target = os.Args[1]
			scanports = strings.ToUpper(os.Args[3])
			fmt.Println("port " + scanports)
		}

		SecPar := strings.ToUpper(os.Args[1])
		fmt.Println("Load " + SecPar)
		if SecPar == "PHPSTUDYDOOR" || SecPar == "PHPSTUDYBACKDOOR" || SecPar == "PHPSTUDYRCE" || SecPar == "PHPSTUDYEXP" {
			exp.PhpStudyDoorExp(os.Args[2], os.Args[3])
			os.Exit(0)
		}

		if SecPar == "CVE-2018-14847" {
			exp.Cve2018_14847Exp(os.Args[2], os.Args[3])
			os.Exit(0)
		}
	} else if ParLen > 4 {
		SecPar := strings.ToUpper(os.Args[1])
		fmt.Println("Load " + SecPar)
		if SecPar == "WINRMCMD" || SecPar == "WINRMEXEC" || SecPar == "WINRMSHELL" {
			rexec.WinrmCmd(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
			os.Exit(0)
		}
		if SecPar == "SSHCMD" || SecPar == "SSHEXEC" || SecPar == "SSHSHELL" {
			ssh.ExecCmd(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
			os.Exit(0)
		}

	} else if ParLen > 3 {
		fmt.Println("Load ===333")
		SecPar := strings.ToUpper(os.Args[1])
		fmt.Println("Load " + SecPar)
	}

	//EndPar := os.Args[ParLen-1]
	//Target := os.Args[ParLen-2]
	fmt.Println("Targe: " + Target)
	//log.Println("Start...")
	fmt.Println("\nScanStart: " + time.Now().Format("2006-01-02 03:04:05"))
	if ParLen == 3 {
		fmt.Println("Load " + EndPar)
		ScanType := strings.ToUpper(EndPar)
		if ScanType == "NBTINFO" {
			nbt.Info(Target, result)
			fmt.Println(" Finished: " + time.Now().Format("2006-01-02 03:04:05"))
			//os.Exit(0)
		}
	}

	ScanType := strings.ToUpper(EndPar)
	if strings.Contains(Target, "/c") || strings.Contains(Target, "/C") {
		CScan(ScanType, Target, result)
	} else if strings.Contains(Target, "/b") || strings.Contains(Target, "/B") {
		BScan(ScanType, Target, result)
	} else if strings.Contains(Target, "/a") || strings.Contains(Target, "/A") {
		AScan(ScanType, Target, result)
	} else if strings.Contains(Target, "http:") || strings.Contains(Target, "https:") {
		LadonUrlScan(ScanType, Target, result)
	} else if strings.ToUpper(Target) == strings.ToUpper("ip.txt") {
		for _, ip := range dic.TxtRead(Target) {
			LadonScan(ScanType, ip, result)
		}
	} else if strings.ToUpper(Target) == strings.ToUpper("ip24.txt") {
		for _, ip := range dic.TxtRead(Target) {
			fmt.Println("\nC_Segment: " + ip)
			fmt.Println("=============================================")
			CScan(ScanType, ip, result)
		}
	} else if strings.ToUpper(Target) == strings.ToUpper("ip16.txt") {
		for _, ip := range dic.TxtRead(Target) {
			fmt.Println("\nB_Segment: " + ip)
			fmt.Println("=============================================")
			BScan(ScanType, ip, result)
		}
	} else if strings.ToUpper(Target) == strings.ToUpper("url.txt") {
		for _, ip := range dic.TxtRead(Target) {
			LadonUrlScan(ScanType, ip, result)
		}
	} else if strings.ToUpper(Target) == strings.ToUpper("host.txt") {
		//for _, ip := range dic.TxtRead(Target) {
		//LadonUrlScan(ScanType,ip)
		//}
	} else if strings.ToUpper(Target) == strings.ToUpper("domain.txt") {
		//for _, ip := range dic.TxtRead(Target) {
		//LadonUrlScan(ScanType,ip)
		//}
	} else if strings.Contains(Target, "-") && strings.Contains(Target, ".") {
		CRange := strings.Split(Target, "-")
		CIP := strings.Split(CRange[0], ".")
		IPC := CIP[0] + "." + CIP[1]
		SIP := strings.Split(CRange[0], ".")[2]
		EIP := strings.Split(CRange[1], ".")[2]
		ips, err := strconv.Atoi(SIP)
		ipe, err := strconv.Atoi(EIP)
		if err != nil {
		}
		for i := ips; i <= ipe; i++ {
			ip := fmt.Sprintf("%s.%d", IPC, i)

			fmt.Println("\nC_Segment: " + ip)
			fmt.Println("=============================================")
			CScan(ScanType, ip, result)

		}
	} else if strings.Contains(Target, "/") {
		if Target != "" {
			ip, ipNet, err := net.ParseCIDR(Target)
			if err != nil {
				fmt.Println(Target + " invalid CIDR")
				return
			}
			var wg sync.WaitGroup
			for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
				wg.Add(1)
				result.Add()
				go func(ip string) {
					defer wg.Done()
					defer result.Done()
					LadonScan(ScanType, ip, result)
				}(ip.String())
			}
			wg.Wait()
		}
	} else {
		LadonScan(ScanType, Target, result)
	}

	//log.Println("Finished")
	fmt.Println(" Finished: " + time.Now().Format("2006-01-02 03:04:05"))
	result.CloseOut()
}
func CEnd() {
	fmt.Println("CFinished: " + time.Now().Format("2006-01-02 03:04:05"))
}
func End() {
	fmt.Println(" Finished: " + time.Now().Format("2006-01-02 03:04:05"))
	os.Exit(0)
}
func CScan(ScanType string, Target string, result mode.Result) {
	ip := strings.Replace(Target, "/c", "", -1)
	ip = strings.Replace(ip, "/C", "", -1)
	ips := strings.Split(ip, ".")
	ip = ips[0] + "." + ips[1] + "." + ips[2]
	var wg sync.WaitGroup
	for i := 1; i < 256; i++ {
		ip := fmt.Sprintf("%s.%d", ip, i)
		wg.Add(1)
		result.Add()
		go func(ip string) {
			defer result.Done()
			defer wg.Done()
			//fmt.Println("c: "+ip)
			LadonScan(ScanType, ip, result)
		}(ip)
	}
	wg.Wait()
	//fmt.Println()
	//spew.Dump(process)
	CEnd()
}
func BScan(ScanType string, Target string, result mode.Result) {
	ip := strings.Replace(Target, "/b", "", -1)
	ip = strings.Replace(ip, "/B", "", -1)
	ips := strings.Split(ip, ".")
	ip = ips[0] + "." + ips[1]
	for i := 0; i < 256; i++ {
		ip := fmt.Sprintf("%s.%d", ip, i)
		fmt.Println("\nC_Segment: " + ip)
		fmt.Println("=============================================")
		CScan(ScanType, ip, result)
	}
}
func AScan(ScanType string, Target string, result mode.Result) {
	ip := strings.Replace(Target, "/a", "", -1)
	ip = strings.Replace(ip, "/A", "", -1)
	ips := strings.Split(ip, ".")
	ip = ips[0]
	for i := 0; i < 256; i++ {
		ip := fmt.Sprintf("%s.%d", ip, i)
		BScan(ScanType, ip, result)
	}
}

func LadonScan(ScanType string, Target string, result mode.Result) {
	if ScanType == "GETEXFQND" || ScanType == "FINDEXCHANGE" {
		//vul.GetExFQND(Target)
	} else if ScanType == "CVE-2021-26855" {
		vul.CheckCVE_2021_26855(Target, result)
	} else if ScanType == "CVE-2021-21972" {
		vul.CheckCVE_2021_21972(Target, result)
	} else if ScanType == "PINGSCAN" || ScanType == "PING" {
		ping.PingName(Target, result)
	} else if ScanType == "ICMPSCAN" || ScanType == "ICMP" {
		//icmp.Icmp(Target,debugLog)
		icmp.Online(Target, result)
	} else if ScanType == "SNMPSCAN" || ScanType == "SNMP" {
		snmp.GetInfo(Target, result)
	} else if ScanType == "ONLINEPC" {
		if isicmp {
			icmp.Icmp(Target, debugLog)
		} else if ping.PingOK(Target) {
			ping.PingName(Target, result)
		}
		snmp.SnmpOK(Target)
	} else if ScanType == "PORTSCAN" || ScanType == "SCANPORT" || ScanType == "TCPSCAN" {
		//Only Scan Port
		/*if isicmp {
		if icmp.IcmpOK(Target) {
			if strings.Contains(scanports, ",") {
				port.ScanPorts(Target,scanports)
			} else {
				port.ScanPort(Target)
			}
		}
		}else if ping.PingOK(Target) {
			if strings.Contains(scanports, ",") {
				port.ScanPorts(Target,scanports)
			} else {
				port.ScanPort(Target)
			}
		}*/
		//Scan Port & Banner
		if isicmp {
			if icmp.IcmpOK(Target) {
				if strings.Contains(scanports, ",") {
					port.ScanPortBanners(Target, scanports, result)
				} else if strings.Contains(scanports, "-") {
					port.ScanPortBannerRange(Target, scanports, result)
				} else if scanports != "" {
					port.ScanPortBannerSingle(Target, scanports, result)
				} else {
					port.ScanPortBanner(Target, result)
				}
			}
		} else if ping.PingOK(Target) {
			if strings.Contains(scanports, ",") {
				port.ScanPortBanners(Target, scanports, result)
			} else if strings.Contains(scanports, "-") {
				port.ScanPortBannerRange(Target, scanports, result)
			} else if scanports != "" {
				port.ScanPortBannerSingle(Target, scanports, result)
			} else {
				port.ScanPortBanner(Target, result)
			}
		}
	} else if ScanType == "TCPBANNER" || ScanType == "PORTSCANBNNER" || ScanType == "SCANPORTBANNER" {
		if isicmp {
			if icmp.IcmpOK(Target) {
				if strings.Contains(scanports, ",") {
					port.ScanPortBanners(Target, scanports, result)
				} else {
					port.ScanPortBanner(Target, result)
				}
			}
		} else if ping.PingOK(Target) {
			if strings.Contains(scanports, ",") {
				port.ScanPortBanners(Target, scanports, result)
			} else {
				port.ScanPortBanner(Target, result)
			}
		}
	} else if ScanType == "HTTPBANNER" || ScanType == "WEBBANNER" {
		http.HttpBanner(Target, result)
	} else if ScanType == "HTTPTITLE" || ScanType == "WEBTITLE" {
		http.ScanTitle(Target, result)
	} else if ScanType == "T3SCAN" || ScanType == "WEBLOGICSCAN" {
		t3.T3version(Target, result)
	} else if ScanType == "OXIDSCAN" || ScanType == "ETHSCAN" {
		dcom.OxidInfo(Target, result)
	} else if ScanType == "MS17010" {
		smb.MS17010(Target, 3, result)
	} else if ScanType == "SMBSCAN" {
		smb.SmbScan(ScanType, Target, result)
	} else if ScanType == "NBTINFO" {
		//nbt.Info(ScanType,Target)
		//nbt.Info(Target)
		//nbt.Info()
	} else if ScanType == "FTPSCAN" {
		ftp.FtpScan(ScanType, Target, result)
	} else if ScanType == "SMBGHOST" || ScanType == "CVE-2020-0796" {
		smb.SmbGhost(Target, 445, result)
	} else if ScanType == "SSHSCAN" {
		ssh.SshScan(ScanType, Target, result)
	} else if ScanType == "MYSQLSCAN" {
		mysql.MysqlScan(ScanType, Target, result)
	} else if ScanType == "MSSQLSCAN" {
		mssql.MssqlScan(ScanType, Target, result)
	} else if ScanType == "MONGODBSCAN" {
		mgo.MongoScan(ScanType, Target, result)
	} else if ScanType == "ORACLESCAN" {
		//oracle.OracleScan(ScanType, Target,result)
	} else if ScanType == "SQLPLUSSCAN" {
		//oracle.SqlPlusScan(ScanType, Target,result)
	} else if ScanType == "WINRMSCAN" {
		winrm.WinrmScan(ScanType, Target, result)
	} else if ScanType == "REDISSCAN" {
		redis.RedisNullScan(ScanType, Target)
	} else if ScanType == "ROUTEROSSCAN" {
		routeros.RouterOSScan(ScanType, Target, result)
	} else if ScanType == "CVE-2018-14847" {
		exp.Cve2018_14847Exp(Target, "8291")
	} else if ScanType == "HTTPBASICSCAN" || ScanType == "BASICAUTHSCAN" || ScanType == "401SCAN" {
		http.BasicAuthScan(ScanType, "http://"+Target, result)
	} else {
		fmt.Println(ScanType, "Moudle Not Found")
		os.Exit(0)
	}
}
func LadonUrlScan(ScanType string, Target string, result mode.Result) {
	if ScanType == "CVE-2021-21972" {
		vul.CheckCVE_2021_21972(Target, result)
	} else if ScanType == "HTTPBANNER" || ScanType == "WEBBANNER" {
		http.HttpBanner(Target, result)
	} else if ScanType == "HTTPTITLE" || ScanType == "WEBTITLE" {
		http.ScanTitle(Target, result)
	} else if ScanType == "HTTPBASICSCAN" || ScanType == "BASICAUTHSCAN" || ScanType == "401SCAN" {
		http.BasicAuthScan(ScanType, Target, result)
	} else {
		fmt.Println(ScanType, "Moudle Not Found")
		os.Exit(0)
	}
}
