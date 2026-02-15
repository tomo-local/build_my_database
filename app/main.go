package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	path := "./data.txt"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("書き込む内容を入力してください\n>")

	if scanner.Scan() {
		line := scanner.Bytes()
		fmt.Printf("入力された文字: %s\n", string(line))
		SaveData1(path, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}

}


func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Sync()
}
