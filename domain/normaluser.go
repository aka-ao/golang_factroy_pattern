package domain

import "fmt"

type NormalUser struct {
	Id   int
	Name string
}

const Postage = 500

func (u *NormalUser) Pay(amount int) string {
	return fmt.Sprintf("一般ユーザ %vさんの支払金額は送料合わせて%v円です", u.Name, amount+500)
}
