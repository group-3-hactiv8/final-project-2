package models

import (
	"final-project-2/helpers"
	"final-project-2/pkg/errs"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" valid:"required~Your age is required"`
}

// stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func (user *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	isAbove8 := govalidator.InRangeInt(user.Age, 9, MaxInt) // kalau 9 tetep true (lower bound)

	if !isAbove8 {
		return errs.NewUnprocessableEntity("Age must have value above 8")
	}

	user.Password = helpers.HashPass(user.Password)

	return nil
}

// func (user *User) BeforeUpdate(tx *gorm.DB) error {
// 	_, err := govalidator.ValidateStruct(user)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
