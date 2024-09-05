package user

import (
	"product-engine/config"
	"product-engine/domain/user"
)

type UserUsecase struct {
	cfg        *config.Config
	userDomain user.UserDomainItf
}

func InitUserUsecase(
	cfg *config.Config,
	userDomain user.UserDomain,
) *UserUsecase {
	return &UserUsecase{
		cfg:        cfg,
		userDomain: userDomain,
	}
}
