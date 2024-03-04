package validation

var (
	emailRegexp = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func ValidatePassword(password string) bool {
	return len(password) > 8

}

func ValidateEmail(email string) bool {
	return len(email) > 8
}
