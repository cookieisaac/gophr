package main

import "errors"
type ValidationError error

var (
	errNoUsername = ValidationError(errors.New("You must supply a username"))
	errNoEmail = ValidationError(errors.New("You must supply an email"))
	errNoPassword = ValidationError(errors.New("You must supply a password"))
	errPasswordTooShort = ValidationError(errors.New("Your password is too short"))
	errUsernameExists = ValidationError(errors.New("That username is taken"))
	errEmailExists = ValidationError(errors.New("That email address has an account"))
	errCredentialsIncorrect = ValidationError(errors.New("We couldn't find a user with supplied username and password combination"))
	errOldPasswordIncorrect = ValidationError(errors.New("Old password is incorrect"))
	errNewPasswordNotMatch = ValidationError(errors.New("New password does not match. Enter new password and confirm again."))
	errInvalidImageType = ValidationError(errors.New("Please upload only jpeg, gif or png images"))
	errNoImage = ValidationError(errors.New("Please select an image to upload"))
	errImageURLInvalid = ValidationError(errors.New("Couldn't download image from the URL you provided"))
)

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}