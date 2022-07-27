package bcrypt

import "golang.org/x/crypto/bcrypt"

func Encrypt(pwd string) (string, error) {
	rawHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(rawHash), nil
}

func Verify(hash string, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)) == nil
}
