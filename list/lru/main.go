// 用单链表实现LRU算法
// LRU 最近最少使用策略
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/daymenu/gods/list"
)

// Meal  饭
type Meal string

const (
	//MaxLength  链表的最大长度
	MaxLength = 10
)

//Compare 比较链表元素是否相等
//实现Comparer 接口
func (m *Meal) Compare(data interface{}) bool {
	meal := data.(*Meal)
	if *meal == *m {
		return true
	}
	return false
}

func (m *Meal) String() string {
	return string(*m)
}
func main() {
	lru := list.Single()
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("你吃了什么饭？")
	for input.Scan() {
		meal := Meal(input.Text())
		if meal == "exit" {
			break
		} else {
			fmt.Println("你吃了什么饭？输入exit退出")
		}
		i, err := lru.LocateElem(&meal)

		if err != list.ErrNotFound {
			_, err := lru.Delete(i)
			if err != nil {
				log.Print(err)
			}
		}

		if err := lru.Insert(1, &meal); err != nil {
			log.Print(err)
		}

		if lru.Length() > MaxLength {
			lru.Delete(MaxLength + 1)
		}

		fmt.Println("最近最常吃的饭是：", lru)
	}
}
