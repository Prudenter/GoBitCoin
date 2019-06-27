/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午4:34
 */

package main

import (
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
