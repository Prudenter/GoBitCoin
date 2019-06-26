/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午4:34
 */

package main

/*
	定义区块链结构
	这里使用数组模拟区块链
*/
type BlockChain struct {
	// 区块链
	Blocks []*Block
}

// 自定义创世语
const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

/*
	定义一个创建区块链的方法
*/
func NewBlockChain() *BlockChain {
	// 创建BlockChain,同时添加一个创世快
	genesisBlock := NewBlock(genesisInfo, nil)
	bc := BlockChain{
		Blocks: []*Block{genesisBlock},
	}
	return &bc
}

/*
	定义一个向区块链中添加区块的方法
	参数:当前区块的数据,不需要提供前一个区块的哈希值,因为bc可以通过自己的下标拿到前一个区块的哈希值
*/
func (bc *BlockChain) AddBlock(data string) {
	// 通过下标,得到最后一个区块
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	// 最后一个区块的哈希值就是是新区块的的前哈希值
	prevHash := lastBlock.Hash

	// 创建block
	newBlock := NewBlock(data, prevHash)

	// 添加区块到区块链中
	bc.Blocks = append(bc.Blocks, newBlock)
}
