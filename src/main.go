package main

import (
	//"crypto"
	//"crypto/rand"
	//"crypto/rsa"
	//_ "crypto/sha256"
	////"fmt"
	//"github.com/cryptoballot/fdh"
	//"github.com/cryptoballot/rsablind"
	//"io/ioutil"
	"runtime"
	//"os"
	"fmt"
	"time"
)
//var quit chan int = make(chan int)

func loop(id int) { // id: 该goroutine的标号
	for i := 0; i < 10; i++ { //打印10次该goroutine的标号
		fmt.Printf("%d ", id)
	}
	quit <- 0
}
func main() {
	time.Sleep(100 * time.Millisecond)
	runtime.GOMAXPROCS(1) // 最多同时使用2个核

	for i := 0; i < 300; i++ { //开三个goroutine
		go loop(i)
	}

	for i := 0; i < 300; i++ {
		<- quit
	}
	//timeUnix_end:=time.Now()
	//shutdownTime := timeUnix_end.Format("2006-01-02 03:04:05 PM")
	//content := "program shutdown time:"+shutdownTime+"\r\n"
	//
	//// 以只写的模式，打开文件
	//f, err := os.OpenFile("F:\\clienttimelog.txt",os.O_WRONLY,0644)
	//if err != nil {
	//	fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	//} else {
	//	// 查找文件末尾的偏移量
	//	n, _ := f.Seek(0, os.SEEK_END)
	//	// 从末尾的偏移量开始写入内容
	//	_, err = f.WriteAt([]byte(content), n)
	//}
	//defer f.Close()
	/*message := []byte("1SDf32kj3asdf")

	keysize := 2048
	hashize := 1536
    println("消息明文：",string(message))
	// We do a SHA256 full-domain-hash expanded to 1536 bits (3/4 the key size)
	hashed := fdh.Sum(crypto.SHA256, hashize, message)
	println("加密hash值：",hashed)
	// Generate a key
	key, _ := rsa.GenerateKey(rand.Reader, keysize)
	println("生成加密密钥key：",key.PublicKey.N)
	// Blind the hashed message
	blinded, unblinder, err := rsablind.Blind(&key.PublicKey, hashed)
	if err != nil {
		panic(err)
	}
	println("生成盲化因子：",blinded)
	println("生成去盲因子：",unblinder)

	// Blind sign the blinded message
	sig, err := rsablind.BlindSign(key, blinded)
	if err != nil {
		panic(err)
	}
	println("生成盲签名：",sig)
	// Unblind the signature
	unblindedSig := rsablind.Unblind(&key.PublicKey, sig, unblinder)
	println("将盲签名去盲：",unblindedSig)
	// Verify the original hashed message against the unblinded signature
	println("用去盲信息校验原始加密hash值是否正确：")
	if err := rsablind.VerifyBlindSignature(&key.PublicKey, hashed, unblindedSig); err != nil {
		panic("failed to verify signature")
	} else {
		println("校验成功：")
		println("ALL IS WELL")
	}*/
}
