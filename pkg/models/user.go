package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `gorm:"name" json:"name"`
	LastName string `gorm:"last_name" json:"lastName"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"password" json:"password"`
}

func hash(user *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func isTheSame(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

func (u *User) Create(db *gorm.DB) error {
	if err := hash(u); err != nil {
		return err
	}
	return db.Save(&u).Error
}

func (u *User) Delete(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Delete(User{}).Error
}

func (u *User) Update(db *gorm.DB, id uint) error {
	if err := hash(u); err != nil {
		return err
	}
	return db.Where("id = ?", id).Save(&u).Error
}

func (u *User) FindByPk(db *gorm.DB, id uint) (*User, error) {
	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (u *User) FindByName(db *gorm.DB, name string) (*User, error) {
	var user User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	} else {
		return &user, err
	}
}

func (u *User) FindByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	} else {
		return &user, err
	}
}

func (u *User) Login(db *gorm.DB) bool {
	user, err := u.FindByEmail(db, u.Email)
	if err != nil {
		return false
	}
	u.ID = user.ID
	u.Email = user.Email
	u.Name = user.Name
	return isTheSame(u.Password, user.Password)
}
