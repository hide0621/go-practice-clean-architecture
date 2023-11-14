package ch6

import "fmt"

func Sub1() {
	sqliteManager := NewSQLiteManager()
	// roomManager := NewRoomManager()

	// どの具体的なデータベースマネージャを使用するかを選択
	mailHandler := MailHandler{dbi: sqliteManager}

	mail := mailHandler.FetchMailById(123)

	fmt.Println(mail)

}

func Sub2() {

	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}

	circle := Circle{
		Radius: 3,
	}

	calculateAndPrintArea(rectangle)
	calculateAndPrintArea(circle)

}
