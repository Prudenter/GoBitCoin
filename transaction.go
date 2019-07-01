/**
* @Author: ASlowPerson
* @Date: 19-6-30 下午9:38
 */

package main

/*
	定义交易结构体
*/
type Transaction struct {
	// 交易id
	TXId []byte
	// 当前交易的输入集合
	TXInputs []TXInput
	// 当前交易的输出集合
	TXOutputs []TXOutput
	// 创建交易的时间戳
	TimeStamp uint64
}

/*
	定义交易输入
*/
type TXInput struct {
	// 这个input所引用的output所在的交易id
	TXId []byte
	// 这个input所引用的output在交易中的索引
	Index int64
	// 付款人对当前交易的签名,(新交易，而不是引用的交易)
	ScriptSig string
}

/*
	定义交易输出
*/
type TXOutput struct {
	// 收款人的公钥哈希,可以理解为地址
	ScriptPubKey string
	// 转账金额
	Value float64
}
