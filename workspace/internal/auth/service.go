package auth

type ServiceImpl struct {
}

var _ Service = (*ServiceImpl)(nil)

func (s ServiceImpl) Create(email, password string) (string, error) {
	// check if the user already exists
	// if not, create a new user with the given email and password
	// hash the password before storing it
	// generate a userId
	// save the user to the database
	// return the userId
	return "generatedUserToken", nil
}

func (s ServiceImpl) Login(email, password string) (string, error) {
	// check if the user already exists
	// if not, create a new user with the given email and password
	// hash the password before storing it
	// generate a userId
	// save the user to the database
	// return the userId
	return "generatedUserToken", nil
}
