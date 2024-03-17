package helper

import (
	// "strings"
	"regexp"
	
)

func isValidEmail(email string) bool {
	// Regular expression for basic email validation
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func ValidateUserInput(userName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	isValiedName := len(userName) >= 2
	// isValiedEmail := strings.Contains(email, "@")
	isValidEmail := isValidEmail(email)
	isValiedUserTicket := userTicket > 0 && userTicket <= remainingTickets

	return isValiedName, isValidEmail, isValiedUserTicket
}
