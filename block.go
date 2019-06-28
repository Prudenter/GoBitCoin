/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午4:34
 */

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

/*
	定义区块结构
 	第一阶段: 先实现基础字段：前区块哈希，哈希，数据
 	第二阶段: 补充字段：Version，时间戳，难度值等
*/
type Block struct {
	//版本号
	Version uint64

	// 前区块哈希
	PrevHash []byte

	//交易的根哈希值
	MerkleRoot []byte

	//时间戳
	TimeStamp uint64

	//难度值, 系统提供一个数据，用于计算出一个哈希值
	Bits uint64

	//随机数，挖矿要求的数值
	Nonce uint64

	// 当前区块的哈希,为了方便,将当前区块的哈希也放入Block中
	Hash []byte

	//数据
	Data []byte
}

/*
	定义创建一个区块的函数
	输入:数据,前区块的哈希值
	输出:区块
*/
func NewBlock(data string, prevHash []byte) *Block {
	b := Block{
		Version:    0,
		PrevHash:   prevHash,
		MerkleRoot: nil,
		TimeStamp:  uint64(time.Now().Unix()),
		Bits:       0,
		Nonce:      0,
		Hash:       nil,
		Data:       []byte(data),
	}

	// 将pow集成到Block中
	pow := NewProofofWork(&b)
	hash, nonce := pow.Run()
	b.Hash = hash
	b.Nonce = nonce

	return &b
}

/*
	定义序列化方法,使用gob编码,将block序列化
	gob编码,go语言内置编解码包,可以支持变长类型的编解码（通用）
*/
func (b *Block) Serialize() []byte {

	// 编码
	var buffer bytes.Buffer
	// 1.创建编码器
	encoder := gob.NewEncoder(&buffer)
	// 2.编码
	err := encoder.Encode(&b)
	if err != nil {
		fmt.Printf("encode err:", err)
		return nil
	}

	return buffer.Bytes()
}

/*
	定义反序列化方法
	输入[]byte,返回block
*/
func Deserialize(src []byte) *Block {
	// 解码
	var block Block
	// 1.创建解码器
	decoder := gob.NewDecoder(bytes.NewReader(src))
	// 2.解码
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Printf("Decode err:", err)
		return nil
	}

	return &block
}
