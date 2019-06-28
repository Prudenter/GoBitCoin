/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午3:52
 */
package main

import (
	"fmt"
)

/*
	打印区块链
*/
func main() {

	// 创建区块链
	err := CreateBlockChain()
	if err != nil {
		fmt.Println("CreateBlockChain err:", err)
		return
	}

	// 获取区块链实例
	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("GetBlockChainInstance err:", err)
		return
	}

	// 使用完关闭数据库
	defer bc.db.Close()

	// 添加一个区块数据
	err = bc.AddBlock("hello world!")
	if err != nil {
		fmt.Println("AddBlock err:", err)
		return
	}

	err = bc.AddBlock("hello BTC!")
	if err != nil {
		fmt.Println("AddBlock err:", err)
		return
	}

	/*// 添加时间缓冲
	time.Sleep(1 * time.Second)
	// 添加区块数据到区块链中
	bc.AddBlock("go语言是世界上最好的语言!")
	time.Sleep(1 * time.Second)
	bc.AddBlock("2019年6月26号btc暴涨20%,突破9万元/个")

	// 打印区块链
	for i, block := range bc.Blocks {
		fmt.Printf("\n---------当前区块高度:%d---------\n", i)
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PrevHash : %x\n", block.PrevHash)
		fmt.Printf("MerkleRoot : %x\n", block.MerkleRoot)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Bits : %d\n", block.Bits)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", block.Data)

		// 其他旷工验证区块
		pow := NewProofofWork(block)
		fmt.Printf("验证结果为:%v\n", pow.IsValid())
	}*/
}
