package adapters

type UserService interface {
	RegisterUser(username, email, password string) error
	AuthenticateUser(username, password string) (bool, error)
}
