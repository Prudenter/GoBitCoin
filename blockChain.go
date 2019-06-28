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

// 定义最后一个区块哈希值的key,用于访问bolt数据库，得到最后一个区块的哈希值
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

			// 5.将最后一个区块的哈希值写入到数据库
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
	db, err := bolt.Open(blockChainDB, 0400, nil)
	if err != nil {
		fmt.Println("bolt.Open err: ", err)
		return nil, err
	}
	// 在这里不关闭数据库,因为后续添加区块时还要使用这个句柄

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
	参数:当前区块的数据,不需要提供前一个区块的哈希值,因为bc中有最后一个区块的哈希值
*/
func (bc *BlockChain) AddBlock(data string) error {
	// 获取区块链中最后一个区块的哈希值
	lastBlockHash := bc.lastHash

	// 1.创建区块
	block := NewBlock(data, lastBlockHash)

	// 2.写入bolt数据库
	err := bc.db.Update(func(tx *bolt.Tx) error {
		// 创建bucket
		bucket := tx.Bucket([]byte(blockBucket))
		// 判断bucket是否存在
		if bucket == nil {
			return errors.New("AddBlock时Bucket不应为空!")
		}

		// 3.写入区块数据,key是区块的哈希值，value是block的字节流
		bucket.Put(block.Hash, block.Serialize())

		// 4.更新最后一个区块的哈希值
		bucket.Put([]byte(lastBlockHashKey), block.Hash)

		// 5.更新bc的lastHash,保证后续AddBlock时可以基于block追加
		bc.lastHash = block.Hash
		return nil
	})
	if err != nil {
		fmt.Printf("bc.db.Update err:", err)
		return err
	}
	return nil
}

/*
	定义迭代器,遍历整个区块链
*/
type Iterator struct {
	// bolt数据库,用于存储数据
	db *bolt.DB
	// 不断变化的哈希值,用于访问链上所有区块
	currentHash []byte
}

/*
	定义方法,将Iterator绑定到BlockChain
*/
func (bc *BlockChain) NewIterator() *Iterator {
	it := Iterator{
		db:          bc.db,
		currentHash: bc.lastHash,
	}
	return &it
}

/*
	给Iterator绑定一个方法:Next
	两个功能: 1.返回当前currentHash所指向的区块
             2.更新currentHash,指向前一个区块
*/
func (it *Iterator) Next() (block *Block, err error) {
	// 读取Bucket中最后一个区块的数据
	// 1.创建bucket
	err = it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		// 判断bucket是否存在
		if bucket == nil {
			return errors.New("Iterator Next时bucket不应为nil")
		}

		// 2.从数据库中读取最后一个区块的数据
		// 得到的数据是block的字节流,需要做解码处理
		// 注意,这里是在做迭代操作,所以一定是currentHash,而不是lastHash
		blockTempInfo := bucket.Get([]byte(it.currentHash))
		// 3.进行解码操作
		block = Deserialize(blockTempInfo)

		// 4.更新currentHash,指向前一个区块
		it.currentHash = block.PrevHash
		return nil
	})
	if err != nil {
		fmt.Println("it.db.View err: ", err)
		return nil, err
	}
	return
}
