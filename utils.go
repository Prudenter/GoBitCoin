/**
* @Author: ASlowPerson
* @Date: 19-6-26 下午11:22
 */

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
	定义将uint64转换为[]byte的函数
*/
func uintToByte(num uint64) []byte {
	var buffer bytes.Buffer
	// 使用二进制编码
	// Write(w io.Writer, order ByteOrder, data interface{}) error
	err := binary.Write(&buffer, binary.LittleEndian, &num)
	if err != nil {
		fmt.Println("binary.Write err: ", err)
		return nil
	}
	return buffer.Bytes()
}
