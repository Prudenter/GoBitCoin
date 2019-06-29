/**
* @Author: ASlowPerson
* @Date: 19-6-29 上午1:38
 */

package main

import (
	"fmt"
	"os"
)

/*
	定义命令行,处理用户输入命令，完成具体函数的调用
*/
type CLI struct {
}

/*
	定义使用说明,帮助用户正确使用
*/
const Usage = `
	Usage:
	./blockchain create "创建区块链"
	./blockchain addBlock <需要写入的的数据> "添加区块"
	./blockchain print "打印区块链"
`

/*
	定义方法,负责解析输入命令
*/
func (cli *CLI) Run() {
	// 获取命令行参数
	cmds := os.Args
	// 判断参数个数,用户至少要输入两个参数
	if len(cmds) < 2 {
		fmt.Println("输入参数无效，请检查!")
		fmt.Println(Usage)
		return
	}

	switch cmds[1] {
	case "create":
		fmt.Println("创建区块方法被调用!")
		cli.createBlockChain()
	case "addBlock":
		// 判断参数个数是否正确
		if len(cmds) != 3 {
			fmt.Println("输入参数无效，请检查!")
			return
		}
		fmt.Println("添加区块方法被调用!")
		data := cmds[2]
		cli.addBlock(data)
	case "print":
		fmt.Println("打印区块链方法被调用!")
		cli.printBlock()
	default:
		fmt.Println("输入参数无效，请检查!")
		fmt.Println(Usage)
	}
}
