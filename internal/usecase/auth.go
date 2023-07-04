package usecase

import (
	"context"
	"fmt"
	"recipes/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *usecase) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *usecase) SignIn(ctx context.Context, req domain.User) (domain.User, bool, error) {
	user, err := u.stor.ReadUser(ctx, req.Login)
	if err != nil {
		return domain.User{}, false, fmt.Errorf("u.stor.ReadUser: %w", err)
	}
	if !u.checkPasswordHash(req.Password, user.Password) {
		return domain.User{}, false, nil
	}
	return user, true, nil
}
