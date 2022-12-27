package domain

import (
	"errors"
	"fmt"
)

const (
	Normal = 1
	Prime  = 2
)

func GetUser(userId int, userName string, status int) (User, error) {
	switch status {
	case Normal:
		return &NormalUser{Name: userName, Id: userId}, nil
	case Prime:
		return &PrimeUser{Name: userName, Id: userId}, nil
	default:
		errMsg := fmt.Sprintf("user factory error: status %v", status)
		return nil, errors.New(errMsg)
	}
}
