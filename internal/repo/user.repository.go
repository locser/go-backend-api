package repo

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetDetailUser() string {

	return "Locser"
}
