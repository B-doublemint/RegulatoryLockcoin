package main

import (
	"./mixcoin"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/cryptoballot/fdh"
	"crypto"
	"github.com/cryptoballot/rsablind"
	"runtime"
	"time"
	//"sync"
	"math/big"
	"crypto/rsa"
	"crypto/rand"
	//"os"
	//"fmt"
	//"strconv"
	"crypto/x509"
	"os"
	"encoding/pem"
	"crypto/sha256"
)

type MixcoinServer struct {
	Address string
	Name string
	PubKey string
}

func chooseServers(servers []*MixcoinServer, numMixes int) []*MixcoinServer {
	//n := len(servers)
	var chosen []*MixcoinServer
	for i := 0; i < numMixes; i++ {
		//chosen = append(chosen, servers[rand.Intn(n)])
	}

	return chosen
}

type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
}

var (
	chunk = &mixcoin.ChunkMessage{
		Val:      4000000,
		SendBy:   601075,
		ReturnBy: 601080,
		OutAddr:  "1SDf32kj3asdf",
		Fee:      2,
		Nonce:    123,
		Confirm:  1,
	}
)
var count int =0
//var key *rsa.PrivateKey
//var blinded []byte
//var unblinder []byte
//var err error
func missionMixCoin(count int) {
	//lock.Lock()
	keysize := 2048
	hashize := 1536
	//log.Printf("OutAddr: %s",chunk.OutAddr)

	hashed := fdh.Sum(crypto.SHA256, hashize, []byte(chunk.OutAddr))
	//send message and id to the TTP
	//生成签名
	sign:=GetSign(hashed,"private.pem")
	if(storedUserInfo(count,hashed,sign)){
		println("已在TTP成功注册")
	}
	println("客户端将转账地址加密：",hashed)

	/*//generate multisingature KA
	multisingatureKA := "03d2cde63d0dca2d3d31c667372b91a33787d6c230700501bc216c0d0229026aeb"
	println("客户端账户publicKey-KA：",multisingatureKA)*/

	// receive a key from server
	key, _ := rsa.GenerateKey(rand.Reader, keysize)
	//key := mixcoin.SendPublicKey()

	// Blind the hashed message
	blinded, unblinder, err := rsablind.Blind(&key.PublicKey, hashed)
	if err != nil {
		panic(err)
	}
	/*println("客户端生成盲化地址：",blinded)
	println("客户端生成去盲因子：",unblinder)*/

	chunk.OutAddr = string(blinded)
	marshaled, err := json.Marshal(chunk)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(marshaled)

	res, err := http.Post("http://localhost:8082/chunk", "application/json", reader)

	if err != nil {
		log.Printf("err")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(bytes.NewBuffer(body))
	responseChunk := mixcoin.ChunkMessage{}
	decoder.Decode(&responseChunk)

	log.Printf("response: %v", responseChunk)
	log.Printf("mixaddr: %s", responseChunk.MixAddr)
	//get blind singature from server
	//sig := mixcoin.Sendsig()
	// Blind sign the blinded message
	sig, err := rsablind.BlindSign(key, []byte(blinded))
	if err != nil {
		panic(err)
	}
	// Unblind the signature
	//outAddrbytes := []byte(responseChunk.OutAddr)
	unblindedSig := rsablind.Unblind(&key.PublicKey, sig, unblinder)
	println("将盲签名去盲：",unblindedSig)
	// Verify the original hashed message against the unblinded signature
	println("用去盲信息校验原始加密hash值是否正确：")
	if err := rsablind.VerifyBlindSignature(&key.PublicKey, hashed, unblindedSig); err != nil {
		panic("failed to verify signature")
	} else {
		println("校验成功！")
	}
	//d1 := []byte("hello\ngo\n")
	//err = ioutil.WriteFile("test.txt", d1, 0644)
	//if err != nil {
	//	panic(err)
	//}

	/*timeUnix_end:=time.Now()
	shutdownTime := timeUnix_end.Format("2006-01-02 03:04:05 PM")
	content := "client NO "+strconv.Itoa(count)+": shutdown time:"+shutdownTime+"\r\n"
	// 以只写的模式，打开文件
	f, err := os.OpenFile("F:\\clienttimelog.txt",os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()*/
	quit <- 0
	//lock.Unlock()
}
//func SendBlindedAddrToServer()([]byte){
//	return blinded
//}
var quit chan int = make(chan int)
func main(){
	GenerateRSAKey(2048)
	timeUnix_begin:=time.Now()
	log.Printf("任务开始，时间: %s", timeUnix_begin.Format("2006-01-02 03:04:05 PM"))
	//beginTime:=timeUnix_begin.Format("2006-01-02 03:04:05 PM")
	//content := "任务开始，时间:"+beginTime+"\r\n"

	/*// 以只写的模式，打开文件
	f, err := os.OpenFile("F:\\clienttimelog.txt",os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()*/

	runtime.GOMAXPROCS(2) // 最多同时使用2个核

	for i := 0; i < 10000; i++ { //开三个goroutine
		go missionMixCoin(i)
	}

	for i := 0; i < 10000; i++ {
		<- quit
	}

	/*//var c int = 0
	lock:=&sync.Mutex{}
	for i:=0;i<1000;i++{
		//传递指针是为了防止 函数内的锁和 调用锁不一致
		go missionMixCoin(lock)
	}
	for{
		lock.Lock()
		c:=count
		lock.Unlock()
		///把时间片给别的goroutine  未来某个时刻运行该routine
		runtime.Gosched()
		if c>=1000{
			break
		}
	}*/
	timeUnix_end:=time.Now()
	log.Printf("任务结束，时间: %s", timeUnix_end.Format("2006-01-02 03:04:05 PM"))
	log.Printf("任务耗时: %d 秒", timeUnix_end.Unix()-timeUnix_begin.Unix())
	log.Printf("任务耗时: %d 纳秒", timeUnix_end.UnixNano()-timeUnix_begin.UnixNano())
}

func verifyWarrant(msg *mixcoin.ChunkMessage, mixPubKey string) bool {
	return true
}
func GenerateRSAKey(bits int){
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err!=nil{
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err!=nil{
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock:= pem.Block{Type: "RSA Private Key",Bytes:X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile,&privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey:=privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey,err:=x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err!=nil{
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock:= pem.Block{Type: "RSA Public Key",Bytes:X509PublicKey}
	//保存到文件
	pem.Encode(publicFile,&publicBlock)
}
//读取RSA私钥
func GetRSAPrivateKey(path string)*rsa.PrivateKey{
	//读取文件内容
	file, err := os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privateKey
}
//读取RSA公钥
func GetRSAPublicKey(path string)*rsa.PublicKey{
	//读取公钥内容
	file, err := os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey
}
//对消息的散列值进行数字签名
func GetSign(msg []byte,path string)[]byte{
	//取得私钥
	privateKey:=GetRSAPrivateKey(path)
	//计算散列值
	hash := sha256.New()
	hash.Write(msg)
	bytes := hash.Sum(nil)
	//SignPKCS1v15使用RSA PKCS#1 v1.5规定的RSASSA-PKCS1-V1_5-SIGN签名方案计算签名
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, bytes)
	if err!=nil{
		panic(sign)
	}
	return sign
}
//验证数字签名 (pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte)
func VerifySign(msg []byte,sign []byte,path string) bool{
	//取得私钥
	publicKey:=GetRSAPublicKey(path)
	//计算散列值
	hash := sha256.New()
	hash.Write(msg)
	bytes := hash.Sum(nil)
	//SignPKCS1v15使用RSA PKCS#1 v1.5规定的RSASSA-PKCS1-V1_5-SIGN签名方案计算签名
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, bytes, sign)
	if err!=nil{
		panic(err)
	}
	return true
}
type userInfoStruct struct{
	userId int
	msg []byte
}
var userInfoStrList [10000]userInfoStruct
var i=0
//regist the user's information
func storedUserInfo(userId int,msg []byte,sign []byte) bool{
	if verifyTheInfo(msg,sign) {
	/*	i := len(userInfoStrList)
		println("length:",i)*/
		userInfoStrList[i].userId = userId
		userInfoStrList[i].msg = msg
		i=i+1
		return true
	}
	return false
}
//varify the message
func verifyTheInfo(msg []byte,sign []byte)bool{

	resule := VerifySign(msg,sign,"public.pem")
	/*	hashize := 1536

		hashed := fdh.Sum(crypto.SHA256, hashize, []byte("1SDf32kj3asdf"))
		if encryptedInfo == string(hashed){
			return true
		}*/
	return resule
}