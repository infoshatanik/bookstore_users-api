package users

import(
	"github.com/infoshatanik/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id           int64    `json:"id,string,omitempty"`
	FirstName    string   `json:"FirstName"`
	LastName     string   `json:"LastName"` 
	Email        string   `json:"Email"`
	DateCreated  string   `json:"DateCreated"`
}

// func Validate (user *User) *errors.RestErr{
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return nil,errors.NewBadRequstError("invalid email add")
// 	}
// 	return nil
// }
func (user *User) Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequstError("invalid email add")
	}
	return nil
}