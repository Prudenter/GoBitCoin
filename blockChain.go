/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午4:34
 */

package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

/*
	定义区块链结构
	结合bolt设计,进行持久化操作
*/
type BlockChain struct {
	// bolt数据库,用于存储数据
	db *bolt.DB
	// 最后一个区块的哈希值
	lastHash []byte
}

// 定义创世语
const genesisInfo = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// 定义数据库名
const blockChainDB = "blockChain.db"

// 定义装所有block的桶
const blockBucket = "blockBucket"

// 定义最后区块哈希值的key,用于访问bolt数据库，得到最后一个区块的哈希值
const lastBlockHashKey = "lastBlockHashKey"

/*
	定义创建区块链的函数,从无到有,这个函数仅执行一次
*/
func CreateBlockChain() error {
	// 1.打开数据库
	// func Open(path string, mode os.FileMode, options *Options) (*DB, error)
	db, err := bolt.Open(blockChainDB, 0600, nil)
	if err != nil {
		fmt.Println("bolt.Open err: ", err)
		return err
	}
	// 2.使用完关闭数据库
	defer db.Close()

	// 3.创建bucket
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		// 判断bucket是否存在
		if bucket == nil {
			// 如果不存在,就创建一个
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				fmt.Println("CreateBucket err :", err)
				return err
			}
			// 4.创建bucket后,写入创世块
			// 创建创世块
			genesisBlock := NewBlock(genesisInfo, nil)
			// 将创世块写入,key是区块的哈希值，value是block的字节流
			// Serialize()是将block序列化的方法
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())

			// 5.将最后区块哈希值写入到数据库
			bucket.Put([]byte(lastBlockHashKey), genesisBlock.Hash)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("db.Update err:", err)
		return err
	}
	return nil
}

/*
	定义获取区块链实例的函数,用于后续操作,每一次有业务时都会调用
	两个功能: 1.如果区块链不存在,则创建,同时返回blockChain的实例
	         2.如果区块链存在,则直接返回blockChain实例
*/
func GetBlockChainInstance() (*BlockChain, error) {
	// 定义变量,接收数据库中最后一个区块的哈希值
	var lastHash []byte

	// 1.打开数据库
	// func Open(path string, mode os.FileMode, options *Options) (*DB, error)
	db, err := bolt.Open(blockChainDB, 0400, nil)
	if err != nil {
		fmt.Println("bolt.Open err: ", err)
		return nil, err
	}
	// 在这里不关闭数据库,因为后续还要重复使用这个句柄

	// 2.创建bucket
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		// 判断bucket是否存在
		if bucket == nil {
			return errors.New("bucket不应为nil!")
		}

		// 3.从数据库中读取最后一个区块的哈希值
		lastHash = bucket.Get([]byte(lastBlockHashKey))
		return nil
	})
	if err != nil {
		fmt.Println("db.View err: ", err)
		return nil, err
	}

	// 4.构造blockChain实例并返回
	bc := BlockChain{db, lastHash}
	return &bc, nil
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
