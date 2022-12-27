package domain

type User interface {
	Pay(amount int) string
}
