package main

import (
	"bufio"
	"os"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	path := "./data.txt"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("書き込む内容を入力してください\n>")

	if scanner.Scan() {
		line := scanner.Bytes()
		fmt.Printf("入力された文字: %s\n", string(line))
		SaveData2(path, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}

}


func SaveData2(path string, data []byte) error {
	rand.Seed(time.Now().UnixNano())
	tmp := fmt.Sprintf("%s.%d.tmp", path, rand.Int())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}

	defer func(){
		fp.Close()
		// 正常に一時的なファイルが作成されている場合は削除する
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err = fp.Write(data); err != nil {
		return err
	}

	if err = fp.Sync(); err != nil {
		return err
	}

	err = os.Rename(tmp, path)

	return err
}
