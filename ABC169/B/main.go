package main

import (
	"bufio"
	"fmt"
	// "math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// 1行読み込み
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 読み込みをint型へキャスト
func nextInt() int64 {
	sc.Scan()
	num, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func nextFloat() float64 {
	sc.Scan()
	num, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return num
}

const max = 1000000000000000000
func main() {
	// scannerの挙動を改行区切り → 空白区切りに変更
	sc.Split(bufio.ScanWords)
	n := nextInt()
	num := nextInt()
	var i int64 = 1
	for ;i < n; {
		num *= nextInt()
		i++
	}
	if num <= max && num >= 0 {
		fmt.Println(num)
	} else {
		fmt.Println("-1")
	}
}