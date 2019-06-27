/**
* @Author: ASlowPerson
* @Date: 19-6-27 上午12:12
 */

package main

import "math/big"

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
