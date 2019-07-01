/**
* @Author: ASlowPerson
* @Date: 19-6-30 下午9:38
 */

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

/*
	定义交易结构体
*/
type Transaction struct {
	// 交易id
	TXId []byte
	// 当前交易的输入集合
	TXInputs []TXInput
	// 当前交易的输出集合
	TXOutputs []TXOutput
	// 创建交易的时间戳
	TimeStamp uint64
}

/*
	定义交易输入
*/
type TXInput struct {
	// 这个input所引用的output所在的交易id
	TXId []byte
	// 这个input所引用的output在交易中的索引
	Index int64
	// 付款人对当前交易的签名,(新交易，而不是引用的交易)
	ScriptSig string
}

/*
	定义交易输出
*/
type TXOutput struct {
	// 收款人的公钥哈希,可以理解为地址
	ScriptPubKey string
	// 转账金额
	Value float64
}

/*
	定义获取交易ID的方法
	对交易做哈希处理
*/
func (tx *Transaction) setHash() error {
	// 对tx做gob编码得到字节流,做sha256,赋值给TXId
	var buffer bytes.Buffer
	// 定义编码器
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		fmt.Println("Encode err: ", err)
		return err
	}
	// 计算tx字节流的哈希值
	hash := sha256.Sum256(buffer.Bytes())
	// 使用tx字节流的哈希值作为交易id
	tx.TXId = hash[:]
	return nil
}

// 自定义挖矿奖励
var reward = 12.5

/*
	定义创建挖矿交易的函数
	参数1:挖矿人,  参数2:付款人对当前交易的签名,挖矿交易不需要签名，所以这个签名字段可以书写任意值，只有矿工有权利写
*/
func NewCoinbaseTx(miner string, data string) *Transaction {
	// 特点：没有输入，只有一个输出，得到挖矿奖励
	// 挖矿交易不需要签名，所以这个签名字段可以书写任意值，只有矿工有权利写
	// 中本聪：写的创世语,现在都是由矿池来写，写自己矿池的名字
	inPut := TXInput{TXId: nil, Index: -1, ScriptSig: data}
	outPut := TXOutput{Value: reward, ScriptPubKey: miner}
	timeStamp := time.Now().Unix()
	// 定义交易结构体并赋值
	tx := Transaction{
		TXId:      nil,
		TXInputs:  []TXInput{inPut},
		TXOutputs: []TXOutput{outPut},
		TimeStamp: uint64(timeStamp),
	}

	// 获取交易ID
	tx.setHash()
	return &tx
}
