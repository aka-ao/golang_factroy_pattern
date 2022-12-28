package repo

import (
	"di_sample/domain"
	"di_sample/dto"
	"di_sample/modules"
	"log"
)

type UserRepositoryImpl struct{}

func NewUserRepo() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) All() (users []domain.User, err error) {

	cmd := `select user_id, user_name, status from user`
	rows, err := modules.Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for rows.Next() {
		var userDto dto.UserDTO
		err = rows.Scan(
			&userDto.UserID,
			&userDto.UserName,
			&userDto.Status,
		)
		if err != nil {
			log.Fatalln(err)
		}

		user, err := domain.GetUser(userDto.UserID, userDto.UserName, userDto.Status)

		if err != nil {
			log.Println(err.Error())
			continue
		}

		users = append(users, user)
	}

	return users, err
}
