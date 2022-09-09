package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

//统计题目数量
var (
	count    float64
	errCount float64
)

//定义答案字典
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

func ans(num, line string) {
	var choose string
	fmt.Print("请选择: ")
	fmt.Scanln(&choose)
	//fmt.Println("你的选择是", choose)
	//打错了保存到错题文件
	if choose != res[num] {
		//fmt.Println("打错了, 正确答案为: ", res[num])
		errCount++
		file, err := os.OpenFile("ansErrof.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			exit("创建错题文件失败")
		}
		info := line + "你的选的的是: " + choose + " 正确答案为: " + res[num] + "\n\n"
		file.WriteString(info)
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
		index := strings.Index(line, ".")
		num := line[:index]
		if errR != nil {
			if b != nil {
				fmt.Println(line)
				count++
				line = line + "\n#"
				ans(num, line)
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
		count++
		ans(num, line)
	}
}

func startAnswer() {
	res = result{}
	//获取答案
	getRes(res, "res.txt")
	//获取题目，答题
	getProblem("problem")
}

func printInfo(time float64, count, errcount float64) {
	correct := count - errcount

	fmt.Printf("答题结束,题目总数为: %.1f, 错误题数: %.1f, 正确题数量为: %.1f, 正确率为: %.2f%%, 耗时: %.2f分钟\n", count, errcount, correct, correct/count*100, time)
}

func main() {
	//删除之前的错题
	os.Remove("ansErrof.txt")
	startTime := time.Now()
	//开始答题
	startAnswer()

	//fmt.Println(time.Since(startTime))
	t := time.Since(startTime)
	time := t.Minutes()
	//fmt.Printf("答题结束,耗时: %.2f分钟\n", t.Minutes())
	printInfo(time, count, errCount)
}
