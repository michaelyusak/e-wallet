package entity

type User struct {
	Id                 int
	Email              string
	Name               string
	Password           string
	HashPassword       []byte
	ProfilePictureName string
	Wallet             Wallet
}
