package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"crypto/md5"
)

type User struct {
	Username 		string
	Email 			string
	HashedPassword	string
	ID				string
}

const (
	hashCost = 10
	passwordLength = 6
	userIDLength = 16
)

func NewUser(username, email, password string) (User, error) {
	user := User{
		Email: email,
		Username: username,
	}
	
	if username == "" {
		return user, errNoUsername
	}
	
	if email == "" {
		return user, errNoEmail
	}
	
	if password == "" {
		return user, errNoPassword
	}
	
	if len(password) < passwordLength {
		return user, errPasswordTooShort
	}
	
	existingUser, err := globalUserStore.FindByUsername(username)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errUsernameExists
	}
	
	existingUser, err = globalUserStore.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errEmailExists
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	
	user.HashedPassword = string(hashedPassword)
	user.ID = GenerateID("usr", userIDLength)
	
	return user, err
	
}

func FindUser(username, password string) (*User, error) {
	out := &User{
		Username: username,
	}
	
	existingUser, err := globalUserStore.FindByUsername(username)
	if err != nil {
		return out, err
	}
	if existingUser == nil {
		return out, errCredentialsIncorrect
	}
	
	if bcrypt.CompareHashAndPassword(
		[]byte(existingUser.HashedPassword),
		[]byte(password),
	) != nil {
		return out, errCredentialsIncorrect
	}
	
	return existingUser, nil
}

func UpdateUser(user *User, email, currentPassword, newPassword, confirmNewPassword string) (User, error) {
	out := *user
	out.Email = email
	
	existingUser, err := globalUserStore.FindByEmail(email)
	if err != nil {
		return out, err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return out, errEmailExists
	}
	
	user.Email = email
	
	if currentPassword == "" {
		return out, nil
	}
	
	if bcrypt.CompareHashAndPassword(
		[]byte(user.HashedPassword),
		[]byte(currentPassword),
	) != nil {
		return out, errOldPasswordIncorrect
	}
	
	if newPassword == "" || confirmNewPassword == "" {
		return out, errNoPassword
	}
	
	if len(newPassword) < passwordLength {
		return out, errPasswordTooShort
	}
	
	if newPassword != confirmNewPassword {
		return out, errNewPasswordNotMatch
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), hashCost)
	user.HashedPassword = string(hashedPassword)
	return out, err
	
}

func (user *User) AvatarURL() string {
	return fmt.Sprintf(
		"//www.gravatar.com/avatar/%x",
		md5.Sum([]byte(user.Email)),
	)
}

func (user *User) ImagesRoute() string {
	return "/user/" + user.ID
}