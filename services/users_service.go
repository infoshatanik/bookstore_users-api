package services

import(
	"github.com/infoshatanik/bookstore_users-api/domain/users"
	"github.com/infoshatanik/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User,*errors.RestErr){
	// user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	// if user.Email == "" {
	// 	return nil,errors.NewBadRequstError("invalid email add")
	// }

	// if err := users.Validate(&user); err != nil {
	// 	return nil, err
	// }
    if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err !=nil {
		return nil,err
	}
	return &user ,nil
	
}

func GetUser(userId int64) (*users.User,*errors.RestErr){
	result := &users.User{Id:userId}
	if err := result.Get(); err != nil {
		return nil,err
	}
	return result ,nil
}