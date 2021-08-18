package dic
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"bufio"
	"strings"
	//"log"
	"os"
)

var (
	dictPath = "public/dict/"
	userPassPath = dictPath+"userpass.txt"
	userPath = dictPath+"user.txt"
	passPath = dictPath+"pass.txt"
)
func UserPassIsExist() bool{
if IsExist(userPassPath) {
	return true
}
return false
}

func PwdIsExist() bool{
if IsExist(userPassPath) {
	return true
}
if IsExist(userPath) {
	return true
}
if IsExist(passPath) {
	return true
}
return false
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func TxtRead(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open ",filename,"error, %v", err)
	}
	fi,_:=os.Stat(filename)
	if fi.Size() ==0 {
	fmt.Println("Error: "+filename+" is null!")
	os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip != "" {
			lines = append(lines, ip)
		}
	}
	return lines
}
func UserDic() (users []string) {
	dicname:=userPath
	file, err := os.Open(dicname)
	if err != nil {
		fmt.Println("Open "+dicname+" error, %v", err)
	}
	fi,_:=os.Stat(dicname)
	if fi.Size() ==0 {
	fmt.Println("Error: "+dicname+" is null!")
	os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	return users
}

func PassDic() (password []string) {
	dicname:=passPath
	file, err := os.Open(dicname)
	if err != nil {
		fmt.Println("Open "+dicname+" error, %v", err)
	}
	fi,_:=os.Stat(dicname)
	if fi.Size() ==0 {
	fmt.Println("Error: "+dicname+" is null!")
	os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			password = append(password, passwd)
		}
	}
	return password
}

func UserPassDic() (userpass []string) {
	dicname:=userPassPath
	file, err := os.Open(dicname)
	if err != nil {
		fmt.Println("Open "+dicname+" error, %v", err)
	}
	fi,_:=os.Stat(dicname)
	if fi.Size() ==0 {
	fmt.Println("Error: "+dicname+" is null!")
	os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			userpass = append(userpass, passwd)
		}
	}
	return userpass
}