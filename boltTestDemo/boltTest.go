/**
* @Author: ASlowPerson
* @Date: 19-6-27 上午11:27
 */
package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

// 定义数据库名
const testDb = "boltDataBase.db"

func main() {
	// 打开数据库
	// func Open(path string, mode os.FileMode, options *Options) (*DB, error)
	db, err := bolt.Open(testDb, 0600, nil)
	if err != nil {
		fmt.Println("bolt open err :", err)
		return
	}
	defer db.Close()

	// 创建bucket
	//func (db *DB) Update(fn func(*Tx) error) error
	err = db.Update(func(tx *bolt.Tx) error {
		// 打开一个bucket
		b1 := tx.Bucket([]byte("testBucket"))
		// 判断b1是否存在
		if b1 == nil {
			// 如果不存在,就创建一个bucket
			b1, err = tx.CreateBucket([]byte("testBucket"))
			if err != nil {
				fmt.Println("CreateBucket err: ", err)
				return err
			}
		}

		// 写入数据
		b1.Put([]byte("key1"), []byte("hello2"))
		b1.Put([]byte("key2"), []byte("world2"))

		// 读取数据
		v1 := b1.Get([]byte("key1"))
		v2 := b1.Get([]byte("key2"))
		v3 := b1.Get([]byte("key3"))

		fmt.Printf("v1:%s\n", string(v1))
		fmt.Printf("v2:%s\n", string(v2))
		fmt.Printf("v3:%s\n", string(v3))

		return nil
	})
	if err != nil {
		fmt.Printf("db.Update err:", err)
		return
	}
	return
}
