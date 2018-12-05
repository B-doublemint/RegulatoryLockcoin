# 可监管的比特币隐私保护混淆服务
本文提出了一种可监管的比特币隐私保护方案。引入可信第三方，使用RSA公平盲签名算法增加了协议的可监管性。同时使用Tor匿名网络和OP_RETURN交易保证用户隐私。

## Environment:

The experiment was performed atop a Dell desktop machine having an Intel Core i5-6500 CPU at 3.20GHz and 4.00G of RAM, running 64-bit windows 10.


## Version of the installation software:

go:  1.8.4

btcd:  0.12.0-beta

btcwallet:  0.7.0-alpha


## Installation:

go sdk https://golang.org/dl/

btcd https://github.com/btcsuite/btcd/releases

btcwallet https://github.com/btcsuite/btcwallet/releases

## Notice:

If you have already completed the above preparations, you can import the project with you IDE(my IDE is [IntelliJ IDEA](https://www.jetbrains.com/idea/)). The IDE will help you install some dependency packages that are required for the program to run. Please note the version number of the dependent package which is very important. 

## Getting Started:

run btcd:

    btcd -u rpcuser -P rpcpass (--testnet --mainnet)

create btcwallet:

    btcwallet -u rpcuser -P rpcpass --create

run bctwallet:

    btcwallet -u rpcuser -P rpcpass -d trace

run server.go and client.go

    You can run the files in the IDE or in the command line.

If everything's working correctly, it'll say "listening on port 8082" or something like that and start downloading a lot of blocks. 
Just like this:

![image](https://github.com/B-doublemint/LockCoin/blob/master/startup.PNG)

Afterwards, it'll just run. And then you can run the client.go, it will send a quest to the server and receive a response with a long signature at the end. As shown below:

![image](https://github.com/B-doublemint/LockCoin/blob/master/responsefromserver.PNG)

PS: We use a parallel strategy to simulate multiple users and test the time required to mix coins in different numbers of users. Of course if you don't want use the Multi-threaded mode, you can comment out the code in the client.go.

If you have any question, please send message to neublockchain@163.com. Good luck.
