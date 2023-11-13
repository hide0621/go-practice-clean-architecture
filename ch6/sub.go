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
	rectangle := NewShape("rectangle")
	rectangle.Width = 10
	rectangle.Height = 5
	fmt.Println(rectangle.Area())

	circle := NewShape("circle")
	circle.Radius = 3
	fmt.Println(circle.Area())
}
