/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午3:52
 */
package main

import "fmt"

/*
	打印区块链
*/
func main() {

	// 创建区块链
	bc := NewBlockChain()

	// 添加区块数据到区块链中
	bc.AddBlock("go语言是世界上最好的语言!")
	bc.AddBlock("2019年6月26号btc暴涨20%,突破9万元/个")

	// 打印区块链
	for i, block := range bc.Blocks {
		fmt.Printf("当前区块高度:%d\n", i)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", block.Data)
	}
}
