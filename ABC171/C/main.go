package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// スライス内に指定した要素があるか判定
// 第1引数 : 要素, 第2要素 : スライス
// return : result, err
func contains(target interface{}, list interface{}) (bool, error) {
	switch list.(type) {
		case []int:
			revert := list.([]int)
			for _, r := range revert {
				if target == r {
					return true, nil
				}
			}
			return false, nil
			
		case []uint64:
			revert := list.([]uint64)
			for _, r := range revert {
				if target == r {
					return true, nil
				}
			}
			return false, nil
			
		case []string:
			revert := list.([]string)
			for _, r := range revert {
				if target == r {
					return true, nil
				}
			}
			return false, nil
		default:
			return false, fmt.Errorf("%v is an unsupported type", reflect.TypeOf(list))
	}

	return false, fmt.Errorf("processing failed")
}

// 1行読み込み
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 読み込みをint型へキャスト
func nextInt() int {
	sc.Scan()
	num, err := strconv.Atoi(sc.Text())
	if(err != nil) {
		panic(err)
	}
	return num
}

// 読み込みをfloat型へキャスト
func nextFloat() float64 {
	sc.Scan()
	num, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	// scannerの挙動を改行区切り → 空白区切りに変更
	sc.Split(bufio.ScanWords)
	n := nextInt()
	sum := 26
	// befo_sum := 0
	truss := 1
	for ; ; {
		if n <= sum {
			break
		}
		// befo_sum = sum
		sum = sum + sum * 26
		truss += 1
	}
	
	lists := make([]int, 0, truss)
	// truss_loop := n / 26 / 26
	cut := n
	rem := 0
	for i := 0; i < truss; i++ {
		for m := 0; m < truss - 5; m++ {
			rem = cut % 26
			cut = cut / 26
		}
		lists = append(lists, cut)
	}
	num_list := make([]int, 0, truss)
	for m := 0; m < truss - 2; m++ {
		num_list = append(num_list, lists[m] / 26)
	}
	fmt.Println(rem, cut, truss, lists, num_list)
}