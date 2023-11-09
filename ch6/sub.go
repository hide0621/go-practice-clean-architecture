package ch6

import "fmt"

func Sub1() {
	sqliteManager := NewSQLiteManager()
	// roomManager := NewRoomManager()

	// どの具体的なデータベースマネージャを使用するかを選択
	sqliteMail := sqliteManager.LoadMailById(123)
	// または、以下のように別のデータベースマネージャを使用できる
	// roomMail := roomManager.LoadMailById(123)

	// ロードされたメールを使用または処理する

	fmt.Println(sqliteMail)
}
