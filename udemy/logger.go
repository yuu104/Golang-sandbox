package main

import (
	"fmt"
	"log"
	"os"
)

func logger() {
	log.Println("情報：アプリケーションが開始されました")

	// log.Fatalln("致命的エラー：設定ファイルが見つかりません")
	log.Panicln("パニックエラー")

	defer fmt.Println("defer")
}

func customLogger() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("カスタムロガー")
}
