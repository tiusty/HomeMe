package models

import (
	"fmt"
	"github.com/revel/revel"
)

type User struct {
	Id                 int
	Created            int64
	Name               string
	Username, Password string
	Email		   string
	HashedPassword     []byte
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

func (user *User) Validate(v *revel.Validation) {

	//Verify Components of the Name field
	ValidateName(v, user.Name).
		Key("user.Name")

	//Verify components of the Username field
	ValidateUsername(v, user.Username).
		Key("user.Username")

	//Verify Components of the password field
	ValidatePassword(v, user.Password).
		Key("user.Password")

	//Verify components of the email field
	v.Email(user.Email)
	ValidateEmail(v, user.Email).
		Key("user.Email")
}
func ValidateName(v *revel.Validation, name string) *revel.ValidationResult {
	return v.Check(name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidateUsername(v *revel.Validation, username string) *revel.ValidationResult {
	return v.Check(username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}


func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}

func ValidateEmail(v *revel.Validation, email string) *revel.ValidationResult {
	return v.Check(email,
		revel.Required{},
		revel.MaxSize{40},
		revel.MinSize{5},
		)
}
