/**
* @Author: ASlowPerson
* @Date: 19-6-29 上午1:39
 */

package main

import "fmt"

/*
	定义创建区块链的方法
*/
func (cli *CLI) createBlockChain() {
	// 创建区块链
	err := CreateBlockChain()
	if err != nil {
		fmt.Println("CreateBlockChain err:", err)
		return
	}
	fmt.Println("创建区块链成功!")
}

/*
	定义添加区块的方法
*/
func (cli *CLI) addBlock(data string) {
	// 获取区块链实例
	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("GetBlockChainInstance err:", err)
		return
	}

	// 使用完关闭数据库
	defer bc.db.Close()

	// 添加一个区块数据
	err = bc.AddBlock(data)
	if err != nil {
		fmt.Println("AddBlock err:", err)
		return
	}
	fmt.Println("添加区块成功!")
}

/*
	定义打印整个区块链的方法
*/
func (cli *CLI) printBlock() {

	// 获取区块链实例
	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("GetBlockChainInstance err:", err)
		return
	}

	// 调用迭代器,遍历输出blockChains
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
