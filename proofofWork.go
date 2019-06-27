/**
* @Author: ASlowPerson
* @Date: 19-6-27 上午12:12
 */

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

/*
	工作量证明,pow
	实现挖矿功能
*/
type ProofofWork struct {
	// 区块:block
	block *Block
	// 目标值:target,这个目标值要与挖矿生成的哈希值比较
	target *big.Int // 结构体,提供了比较大小的方法
}

/*
	定义创建ProofofWork的函数
	block由用户提供
	target目标由系统提供
*/
func NewProofofWork(block *Block) *ProofofWork {
	pow := ProofofWork{
		block: block,
	}
	// 难度值先写死,不去推导,后续再补充推导方式
	// 16进制的哈希值形式的字符串,共64位
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	tempBigInt := new(big.Int)

	// 将自定义的难度值赋值给tempBigInt
	tempBigInt.SetString(targetStr, 16)
	pow.target = tempBigInt
	return &pow
}

/*
	定义函数,拼接nonce和block数据
*/
func (pow *ProofofWork) PrepareData(nonce uint64) []byte {
	b := pow.block

	// 构造一个二维字节切片
	temp := [][]byte{
		uintToByte(b.Version), //将uint64转换为[]byte
		b.PrevHash,
		b.MerkleRoot,
		uintToByte(b.TimeStamp),
		uintToByte(b.Bits),
		uintToByte(nonce),
		b.Data,
	}
	// 使用Join方法进行拼接,将二维切片转为一维切片
	data := bytes.Join(temp, []byte{})
	return data
}

/*
	定义挖矿函数,不断变化nonce,使得sha256(数据+noce) < 难度值
	返回值: 当前区块的哈希,nonce
*/
func (pow *ProofofWork) Run() ([]byte, uint64) {
	// 定义随机数
	var nonce uint64
	var hash [32]byte
	fmt.Println("开始挖矿..")
	for {
		fmt.Printf("%x\r", hash[:])
		// 1.拼接字符串+nonce
		data := pow.PrepareData(nonce)

		// 2.计算哈希值
		hash = sha256.Sum256(data)

		// 3.将hash转换为bigInt类型
		tmpInt := new(big.Int)
		tmpInt.SetBytes(hash[:])

		// 4.比较当前哈希值与目标难度值的大小
		// func (x *Int) Cmp(y *Int) (r int)
		// -1 if x <  y
		// 0 if x == y
		// +1 if x >  y
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功!计算出的hash为:%x,随机数nonce为:%d\n", hash[:], nonce)
			break
		}
		// 如果计算出的哈希值不小于目标难度值
		nonce++
	}
	return hash[:], nonce
}

/*
	定义其他旷工验证区块的函数
*/
func (pow *ProofofWork) IsValid() bool {
	fmt.Println("开始验证..")
	// 1.拼接字符串+nonce
	data := pow.PrepareData(pow.block.Nonce)

	// 2.计算哈希值
	hash := sha256.Sum256(data)

	// 3.将hash转换为bigInt类型
	tmpInt := new(big.Int)
	tmpInt.SetBytes(hash[:])

	// 4.比较当前哈希值与目标难度值的大小,满足条件,返回true
	return tmpInt.Cmp(pow.target) == -1
}
