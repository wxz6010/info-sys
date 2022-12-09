package utils

import (
	"com.alex.infosys/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

func CryptoPass(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash)
}

func ISEqScr(pass ...string) error {
	if len(pass) != 2 {
		return errs.New("新旧密码")
	}
	err := bcrypt.CompareHashAndPassword([]byte(pass[0]), []byte(pass[1]))
	return err
}

func CheckPass(pass string) error {
	if len(pass) < 6 {
		return errs.New("密码必须大于六位")
	}
	return nil
}
