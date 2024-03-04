package adapters

type AuthAdapter interface {
	Register(username string, email string, password string) error
	Login(username string, password string) (string, error)
}
