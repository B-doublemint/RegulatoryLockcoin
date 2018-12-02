package main
/*import (
	"os"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/rand"
	"crypto"
)

//生成RSA私钥和公钥，保存到文件中
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
}*/
func main(){
	//生成密钥文件
	GenerateRSAKey(2048)
/*	//模拟发送方
	//要发送的消息
	msg:=[]byte("hello world")
	//生成签名
	sign:=GetSign(msg,"private.pem")

	//模拟接收方
	//接受到的消息
	acceptmsg:=[]byte("hello world")
	//接受到的签名
	acceptsign:=sign
	//验证签名
	result:=VerifySign(acceptmsg,acceptsign,"public.pem")
	fmt.Println("验证结果：",result)*/
}