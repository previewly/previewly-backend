package token

import (
	"context"

	"wsw/backend/lib/utils"

	tokenModel "wsw/backend/model/token"

	"github.com/golobby/container/v3"
)

func ResolveCreateToken(ctx context.Context) (string, error) {
	var model tokenModel.Token
	err := container.Resolve(&model)
	if err != nil {
		utils.F("Couldnt resolve model Token: %v", err)
		return "", err
	}
	token, err := model.CreateToken()
	if err != nil {
		return "", err
	}
	return *token, nil
}
