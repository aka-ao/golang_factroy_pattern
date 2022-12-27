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
		return &NormalUser{UserName: userName, UserId: userId}, nil
	case Prime:
		return &PrimeUser{UserName: userName, UserId: userId}, nil
	default:
		errMsg := fmt.Sprintf("user factory error: status %v", status)
		return nil, errors.New(errMsg)
	}
}
