package util

import "golang.org/x/crypto/bcrypt"

func GenerateUserPwdHash(pwd string) (*string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hash := string(hashedPwd)
	return &hash, nil
}

func CompareUserPassword(hash, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)) == nil
}
