package mixcoin

import (
	"github.com/conformal/btcjson"
	"github.com/conformal/btcrpcclient"
	"github.com/conformal/btcutil"
	"github.com/conformal/btcwire"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

type RpcClient interface {
	NotifyBlocks() error
	WalletPassphrase(string, int64) error
	CreateEncryptedWallet(string) error
	GetNewAddress() (btcutil.Address, error)
	GetBestBlock() (*btcwire.ShaHash, int32, error)
	CreateRawTransaction([]btcjson.TransactionInput, map[btcutil.Address]btcutil.Amount) (*btcwire.MsgTx, error)
	SignRawTransaction(*btcwire.MsgTx) (*btcwire.MsgTx, bool, error)
	SendRawTransaction(*btcwire.MsgTx, bool) (*btcwire.ShaHash, error)
	NotifyReceivedAsync([]btcutil.Address) btcrpcclient.FutureNotifyReceivedResult
	ListUnspentMinMaxAddresses(int, int, []btcutil.Address) ([]btcjson.ListUnspentResult, error)
	ImportPrivKey(*btcutil.WIF) error
	SendToAddress(btcutil.Address, btcutil.Amount) (*btcwire.ShaHash, error)
	//getblockhash() (string,error)
}

func NewRpcClient() RpcClient {
	log.Println("starting rpc client")

	log.Printf("Reading btcd cert file %v", cfg.CertFile)
	//certs, err := ioutil.ReadFile(cfg.CertFile)
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		log.Panicf("couldn't read btcd certs")
	}

	//connCfg := &btcrpcclient.ConnConfig{
	//	Host:         cfg.RpcAddress,
	//	Endpoint:     "ws",
	//	User:         cfg.RpcUser,
	//	Pass:         cfg.RpcPass,
	//	Certificates: certs,
	//	DisableTLS:   false,
	//	HttpPostMode: false,
	//	DisableConnectOnNew: false,
	//	DisableAutoReconnect: false,
	//}

	connCfg := &btcrpcclient.ConnConfig{
		Host:         "localhost:18334",
		Endpoint:     "ws",
		User:         "mixcoin",
		Pass:         "Mixcoin1",
		Certificates: certs,
	}
	println("證書：",string(certs))
	client, err := btcrpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal("lineno57:",err)
	}
	//defer client.Shutdown()
	//ntfnHandlers := &btcrpcclient.NotificationHandlers{
	//	OnBlockConnected: onBlockConnected,
	//}

	//client, err := btcrpcclient.New(connCfg, nil)
	//if err != nil {
	//	log.Printf("error creating rpc client: %v", err)
	//}
	log.Printf("*************已经生成Client*************")
	if client==nil {
		log.Printf("生成的Client是TM空的！！！！")
	}
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("**********************Block count: %d*********************", blockCount)
	// Register for block connect and disconnect notifications.
	if err = client.NotifyBlocks(); err != nil {
		log.Printf("error setting notifyblock")
	}

	log.Printf("unlocking wallet")

	err = client.WalletPassphrase(cfg.WalletPass, 7200)
	if err != nil {
		log.Printf("error unlocking wallet: %v", err)
		log.Printf("trying to create a new wallet:")
		err = client.CreateEncryptedWallet(cfg.WalletPass)
		if err != nil {
			log.Panicf("error creating new wallet: %v", err)
		}
		log.Printf("successfully created a new wallet")
	}
	log.Printf("successfully unlocked wallet")
	return client
}

func getNewAddress() (btcutil.Address, error) {
	addr, err := rpc.GetNewAddress()
	if err != nil {
		log.Panicf("error getting new address: %v", err)
	}
	/**
	err = rpcClient.SetAccount(addr, cfg.MixAccount)
	if err != nil {
		log.Panicf("error setting account: %v", err)
		return nil, err
	}
	*/
	return addr, nil
}

// TODO only update occasionally; no need to check every time
func getBlockchainHeight() int {
	log.Printf("getting blockchain height")
	_, height32, err := rpc.GetBestBlock()
	if err != nil {
		log.Panicf("error getting blockchain height: %v", err)
	}
	log.Printf("got blockchain height: %v", height32)
	return int(height32)
}
func AddMultiSigAddress()string{
	time.Sleep(1 * time.Second)
	log.Printf("getting MultiSigAddress")
	//multiSign,err := rpc.addMultiSigAddress()
	multiSign := "3GJm9skU7JEw32zxAsFxDXNGbEGrFdSGsD"
	//if err != nil{
	//	log.Panicf("error getting MultiSigAddress: %v", err)
	//}
	log.Printf("got MultiSigAddress: %v", multiSign)
	return multiSign
}