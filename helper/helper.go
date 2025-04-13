package helper

import "strings"

func UserInputValidation(firstname string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) > 2 && len(lastname) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}

