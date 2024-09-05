package user

import (
	"context"
	"errors"
	"product-engine/common/auth"
)

func (userUC *UserUsecase) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {

	userDetail, err := userUC.userDomain.Get(ctx, req.Username)
	if err != nil {
		return LoginResponse{}, err
	}

	if req.Password != userDetail.Password {
		return LoginResponse{}, errors.New("wrong password")
	}

	token, err := auth.GenerateJWT(userDetail.ID, userUC.cfg.JWT.Secret)
	if err != nil {
		return LoginResponse{}, err
	}

	res := LoginResponse{
		Token: token,
	}

	return res, nil
}
