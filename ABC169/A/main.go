package main

import (
	"bufio"
	"fmt"
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
func nextInt() int {
	sc.Scan()
	num, err := strconv.Atoi(sc.Text())
	if(err != nil) {
		panic(err)
	}
	return num
}

func main() {
	// scannerの挙動を改行区切り → 空白区切りに変更
	sc.Split(bufio.ScanWords)
	a, b := nextInt(), nextInt()

	fmt.Println(a * b)
}