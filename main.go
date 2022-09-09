package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type result map[string]string

var res result

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func getRes(res result, filePath string) result {
	data, err := os.ReadFile(filePath)
	if err != nil {
		exit("打开答案文件失败")
	}

	s := string(data)
	line := regexp.MustCompile(`[\s]+`).ReplaceAllString(strings.TrimSpace(s), "")
	line = line[1 : len(line)-1]
	resList := strings.Split(line, ",")
	for _, r := range resList {
		rlist := strings.Split(r, ":")
		//fmt.Println(rlist)
		res[rlist[0]] = rlist[1]
	}
	return res
}

func ans(num string) {
	var choose string
	fmt.Print("请选择: ")
	fmt.Scanln(&choose)
	fmt.Println("你的选择是", choose)
	if choose != res[num] {
		fmt.Println("打错了, 正确答案为: ", res[num])
	}
}

func getProblem(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		exit("打开文件失败")
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('#')
		line := regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(string(b)), "\n")
		num := line[:2]
		if errR != nil {
			if b != nil {
				fmt.Println(string(b))
				ans(num)
			}
			if errR == io.EOF {
				//fmt.Println(string(b))
				break
			}

			fmt.Println(errR.Error())
		}
		//line := regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(string(b)), "\n")
		//num := line[:2]
		//fmt.Println("num: ", num)
		fmt.Println(line[:len(line)-1])
		ans(num)
	}
}

func startAnswer() {
	res = result{}
	//获取答案
	getRes(res, "res.txt")
	//获取题目，答题
	getProblem("problem")
}

func main() {
	//开始答题
	startAnswer()
}