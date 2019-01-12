---
layout: post
title: golang进阶:使用go生成私有证书实现gRPC+TLS服务
category: golang
tags: golang golang进阶
description: 
keywords: golang,gRPC+TLS,qcmd-slim
date: 2019-01-09T13:19:54+08:00
score: 5.0
coverage: qcmd_grpc.jpg
published: false
---

## 1.简洁

## 2.什么是SSL/TLS通信
不使用SSL/TLS的HTTP通信，就是不加密的通信。所有信息明文传播，带来了三大风险。

1. 窃听风险（eavesdropping）：第三方可以获知通信内容。
2. 篡改风险（tampering）：第三方可以修改通信内容。
3. 冒充风险（pretending）：第三方可以冒充他人身份参与通信。

SSL/TLS协议是为了解决这三大风险而设计的，希望达到：

1. 所有信息都是加密传播，第三方无法窃听。
2. 具有校验机制，一旦被篡改，通信双方会立刻发现。
3. 配备身份证书，防止身份被冒充。

互联网是开放环境，通信双方都是未知身份，这为协议的设计带来了很大的难度。而且，协议还必须能够经受所有匪夷所思的攻击，这使得SSL/TLS协议变得异常复杂。
## 3.SSL/TLS运行过程
**SSL/TLS协议的基本思路是采用公钥加密法，也就是说，客户端先向服务器端索要公钥，然后用公钥加密信息，服务器收到密文后，用自己的私钥解密。**

SSL/TLS协议的基本过程(握手阶段handshake)是这样的：

1. 客户端向服务器端索要并验证公钥。
2. 双方协商生成"对话密钥"。
3. 双方采用"对话密钥"进行加密通信。

![tsl_handshake](/assets/image/tls_handshake.gif)

## 3.为什么要要使用TLS
## 4.go签发私钥
## 5.最佳实践
## 6.总结

