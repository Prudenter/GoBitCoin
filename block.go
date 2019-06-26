/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午4:34
 */

package main

import (
	"bytes"
	"crypto/sha256"
)

/*
	定义区块结构
 	第一阶段: 先实现基础字段：前区块哈希，哈希，数据
 	第二阶段: 补充字段：Version，时间戳，难度值等
*/
type Block struct {
	// 前区块哈希
	PrevHash []byte
	// 当前区块的哈希,为了方便,我们将当前区块的哈希也放入Block中
	Hash []byte
	// 数据
	Data []byte
}

/*
	定义创建一个区块的函数
	输入:数据,前区块的哈希值
	输出:区块
*/
func NewBlock(data string, prevHash []byte) *Block {
	b := Block{
		PrevHash: prevHash,
		Hash:     nil,
		Data:     []byte(data),
	}

	// 计算哈希值
	b.setHash()
	return &b
}

/*
	定义计算区块哈希值的方法
*/
func (b *Block) setHash() {
	// 比特币哈希算法:func Sum256(data []byte) [Size]byte
	// 这里data是block中各个字段拼成的字节流
	// 我们使用Join(s [][]byte, sep []byte) []byte拼接各个切片字段,接收一个二维的切片,使用一个一维切片拼接,返回拼接后的一维切片.

	// 构造一个二维字节切片
	temp := [][]byte{
		b.PrevHash,
		b.Hash,
		b.Data,
	}
	// 使用Join方法进行拼接
	data := bytes.Join(temp, []byte{})

	// 计算当前区块的哈希
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}
