package domain

import "fmt"

type PrimeUser struct {
	Id   int
	Name string
}

func (u *PrimeUser) Pay(amount int) string {
	return fmt.Sprintf("プライムユーザ %vさんの支払金額は送料無料で%v円です", u.Name, amount)
}
