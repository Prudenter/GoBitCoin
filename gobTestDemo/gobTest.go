/**
* @Author: ASlowPerson
* @Date: 19-6-28 下午5:43
 */
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
	gob编解码测试Demo
*/

type Person struct {
	Name string
	Age  uint64
}

func main() {
	// 编解码Person结构
	p1 := Person{
		"张三",
		18,
	}

	// 编码
	var buffer bytes.Buffer
	// 1.创建编码器
	encoder := gob.NewEncoder(&buffer)
	// 2.编码
	err := encoder.Encode(&p1)
	if err != nil {
		fmt.Printf("encode err:", err)
		return
	}
	data := buffer.Bytes()
	fmt.Printf("编码后的数据:%x\n", data)

	// 解码
	var p2 Person
	// 1.创建解码器
	decoder := gob.NewDecoder(bytes.NewReader(data))
	// 2.解码
	err = decoder.Decode(&p2)
	if err != nil {
		fmt.Printf("Decode err:", err)
		return
	}
	fmt.Printf("解码后的数据:%v\n", p2)
}
