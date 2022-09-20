package module

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Ans struct {
	ID       string
	Question string
	Answer   string
	A        string
	B        string
	C        string
	D        string
}

var (
	Qs []*Ans = []*Ans{}
	//Err []*string = []*string{}
)

func List() []*Ans {
	return Qs
}

func ADD(id, q, ans, a, b, c, d string) {
	Qs = append(Qs, &Ans{id, q, ans, a, b, c, d})
}

func Check(w http.ResponseWriter, r *http.Request) []*string {
	Err := []*string{}
	for _, q := range Qs {
		if q.Answer == r.PostFormValue(q.ID) {
			continue
		}
		info := fmt.Sprintf("第 %s 题错误, %s\n,你选择的是: %s, 正确答案是: %s\n", q.ID, q.Question, r.PostFormValue(q.ID), q.Answer)
		//w.Write([]byte(info))
		Err = append(Err, &info)
	}
	return Err
}

/生成count个[start, end)结束的不重复的随机数
func randNums(start, end, count int) []int {
	//检查范围
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器,加入时间戳保证每次生成的随机数白羊
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(nums) < count {

		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}

		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}