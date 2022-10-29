package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var task = sync.WaitGroup{}

func timer(duration time.Duration, content string) {
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			fmt.Println(content)
			return
		default:
			print("")
		}
	}
}
func ticker2(duration time.Duration, content string) {
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			fmt.Println(content)
			return
		default:
			print("")
		}
	}
	task.Done()
}

func knock() {
	fmt.Println("请设置提醒的时间,格式为:2006-01-02 15:04:05")
	var str string
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	time1, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		fmt.Println(err)
		return
	}

	//时间差值
	timenow1 := time.Now()
	timenow := timenow1.Add(8 * time.Hour)
	subnum := time1.Sub(timenow)
	//打印内容
	var content string
	fmt.Printf("请设置提醒的内容:\n")
	fmt.Scanf("%s", &content)
	fmt.Printf("你预定的\"%v\"将于%v分钟后开始\n\n\n", content, subnum.Minutes())
	timer(subnum, content)
}
func knock3(k int, content string) {
	a := time.Now()
	p := a.Weekday()
	wow := k - int(p)
	if wow < 0 {
		wow = 7 + wow
	}
	time.Sleep(time.Duration(wow) * 24 * time.Hour)
	for {
		time.Sleep(time.Hour * 24 * 7)
		println(content)
	}
}
func knock2() {
	fmt.Println("请设置提醒的时间,格式为:2006-01-02 15:04:05")
	var str string
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	time1, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		fmt.Println(err)
		return
	}

	//时间差值
	timenow1 := time.Now()
	timenow := timenow1.Add(8 * time.Hour)
	subnum := time1.Sub(timenow)
	//打印内容
	var content string
	fmt.Printf("请设置提醒的内容:\n")
	fmt.Scanf("%s", &content)
	fmt.Printf("你预定的\"%v\"将于%v分钟后开始,并之后每天提醒你一次\n\n\n", content, subnum.Minutes())
	task.Add(1)
	ticker2(subnum, content)
	task.Wait()

	for {
		time.Sleep(time.Hour * 24)
		fmt.Println(content)
	}

}

func main() {
	fmt.Println("提醒功能如下\n1.单次日程提醒功能\n2.每天多久提醒一次\n3.每周多久提醒一次\n请输入序号:")
	var a int
	fmt.Scanf("%d\n", &a)
	if a == 1 {
		knock()
	}
	if a == 2 {
		knock2()
	}
	if a == 3 {
		fmt.Println("请输入你想在7个星期中设置几个星期进行此项目?")
		var n int
		k := make([]int, 7)
		fmt.Scanf("%d\n", &n)
		fmt.Println("输入星期,输入一个你就回车一个,懂吧,你懂的,如果不懂可以看我的代码,嘻嘻(以1234567代表星期一...)")
		for i := 0; i < n; i++ {
			fmt.Scanf("%d\n", &k[i])
		}
		fmt.Println("请设置提醒的时间,格式为:2006-01-02 15:04:05")
		var str string
		reader := bufio.NewReader(os.Stdin)
		str, _ = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		time1, err := time.Parse("2006-01-02 15:04:05", str)
		if err != nil {
			fmt.Println(err)
			return
		}
		//时间差值
		timenow1 := time.Now()
		timenow := timenow1.Add(8 * time.Hour)
		subnum := time1.Sub(timenow)
		//打印内容
		var content string
		fmt.Printf("请设置提醒的内容:\n")
		fmt.Scanf("%s\n", &content)
		fmt.Printf("你预定的\"%v\"将于于你所设定的那时候提醒你,并且在你设定的周也会在那个时间提醒你!\n\n\n", content)
		timer(subnum, content)
		for j := 0; j < n; j++ {
			go knock3(k[j], content)
		}
	}
}
