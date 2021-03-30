package database

type User struct {
	ID    uint   `gorm:"primarykey"`
	Login string `gorm:"unique"`
	Hash  string
}
type SessionTable struct {
	ID      uint `gorm:"primarykey"`
	Login   string
	Session string
	Expiry  int
}
