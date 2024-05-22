package authentication

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Creata a slice of bytes from the passsword param and hash it with the cost of 16, then save it into the bytes var
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err
}
func ComparePassword(password, hash string) bool {
	// Compare the slice of bytes from the password and the slice of bytes from the unhashed password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
