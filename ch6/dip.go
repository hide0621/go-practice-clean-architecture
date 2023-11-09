package ch6

type SQLiteManager struct {
	// SQLiteManagerの具体的な実装の詳細がここに含まれる
}

func NewSQLiteManager() *SQLiteManager {
	// SQLiteManagerの初期化ロジック
	return &SQLiteManager{}
}

func (s *SQLiteManager) LoadMailById(id int) Mail {
	// SQLiteデータベースからデータを読み込む具体的な実装
	return Mail{}
}

func (s *SQLiteManager) SaveMail(mail Mail) bool {
	// SQLiteデータベースにデータを保存する具体的な実装
	return true
}

type RoomManager struct {
	// RoomManagerの具体的な実装の詳細がここに含まれる
}

func NewRoomManager() *RoomManager {
	// RoomManagerの初期化ロジック
	return &RoomManager{}
}

func (r *RoomManager) LoadMailById(id int) Mail {
	// Roomデータベースからデータを読み込む具体的な実装
	return Mail{}
}

func (r *RoomManager) SaveMail(mail Mail) bool {
	// Roomデータベースにデータを保存する具体的な実装
	return true
}

type Mail struct {
	// Mailの構造体の詳細がここに含まれる
}
