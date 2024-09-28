package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"
	"wsw/backend/graph/convertor"
	"wsw/backend/graph/model"
	"wsw/backend/lib/utils"
	tokenModel "wsw/backend/model/token"
	urlModel "wsw/backend/model/url"

	container "github.com/golobby/container/v3"
)

// CreateToken is the resolver for the createToken field.
func (r *mutationResolver) CreateToken(ctx context.Context) (string, error) {
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

// AddURL is the resolver for the addUrl field.
func (r *mutationResolver) AddURL(ctx context.Context, token string, url string) (*model.PreviewData, error) {
	var tokenModelImpl tokenModel.Token
	err := container.Resolve(&tokenModelImpl)
	if err != nil {
		return nil, err
	}

	if !tokenModelImpl.IsTokenExist(token) {
		return nil, errors.New("invalid token")
	}
	var urlModelImpl urlModel.Url
	errURLModel := container.Resolve(&urlModelImpl)
	if errURLModel != nil {
		return nil, errURLModel
	}

	previewData, errData := urlModelImpl.AddURL(url)
	if errData != nil {
		return nil, errData
	}
	return convertor.ConvertPreviewData(previewData), nil
}

// GetPreviewData is the resolver for the getPreviewData field.
func (r *queryResolver) GetPreviewData(ctx context.Context, token string, url string) (*model.PreviewData, error) {
	var model tokenModel.Token
	err := container.Resolve(&model)
	if err != nil {
		return nil, err
	}

	previewData, err := model.GetPreviewData(token)
	if err != nil {
		return nil, err
	}
	return convertor.ConvertPreviewData(previewData), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
