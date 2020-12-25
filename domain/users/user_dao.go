package users

import(
	"github.com/infoshatanik/bookstore_users-api/utils/errors"
	"fmt"
	"time"
)

var (
	usersDB = make(map[int64]*User) 
)

func (user *User) Get() *errors.RestErr{
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found",user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email{
			return errors.NewBadRequstError(fmt.Sprintf("email %d already exists",user.Email))
		}
		return errors.NewBadRequstError(fmt.Sprintf("user %d already exists",user.Id))
	} 
	now := time.Now()
	user.DateCreated = now.Format("2006-01-02 15:04:05")
	usersDB[user.Id] = user
	return nil
}