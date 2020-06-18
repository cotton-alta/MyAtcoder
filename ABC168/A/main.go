package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

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
	num := strings.Split(nextLine(), "")

	last_str, _ := strconv.Atoi(num[len(num) - 1])

	if last_str == 3 {
		fmt.Println("bon")
	} else if last_str == 0 || last_str == 1 || last_str == 6 || last_str == 8 {
		fmt.Println("pon")
	} else {
		fmt.Println("hon")
	}
}