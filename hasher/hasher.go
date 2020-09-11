package hasher

import "golang.org/x/crypto/bcrypt"

//GenerateHashedPassword creates the hashed password for security
func GenerateHashedPassword(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

//ComparePasswords handles the password and hash comparison
func ComparePasswords(password []byte, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err
}