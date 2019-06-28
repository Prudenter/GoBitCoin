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

	// 调用迭代器,遍历输出blockChain
	it := bc.NewIterator()
	for {
		// 调用Next方法,从后向前逐一获取区块
		block, err := it.Next()
		if err != nil {
			fmt.Println("Next err: ", err)
			return
		}
		// 打印每个区块的详细信息
		fmt.Printf("\n--------------------\n")
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

		// 退出条件
		if block.PrevHash == nil {
			fmt.Println("区块链遍历结束!")
			break
		}
	}
}
