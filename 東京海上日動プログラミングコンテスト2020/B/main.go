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
	a, v := nextInt(), nextInt()
	b, w := nextInt(), nextInt()
	t := nextInt()

	dis := a - b
	spd := w - v

	if dis < 0 {
		dis = dis * -1
	} else if dis == 0 {
		fmt.Println("YES")
		return
	}
	
	if spd >= 0 {
		fmt.Println("NO")
	} else {
		spd = spd * -1
		if dis <= spd * t {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")		
		}
	}
}