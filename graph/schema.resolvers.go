package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"errors"

	"wsw/backend/domain/image/process"
	"wsw/backend/domain/image/upload"
	"wsw/backend/graph/convertor"
	"wsw/backend/graph/model"
	"wsw/backend/lib/utils"
	tokenModel "wsw/backend/model/token"
	urlModel "wsw/backend/model/url"
	"wsw/backend/resolvers/token"

	container "github.com/golobby/container/v3"
)

// CreateToken is the resolver for the createToken field.
func (r *mutationResolver) CreateToken(ctx context.Context) (string, error) {
	return token.ResolveCreateToken(ctx)
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

// Upload is the resolver for the upload field.
func (r *mutationResolver) Upload(ctx context.Context, token string, images []*model.UploadInput) ([]*model.UploadImageStatus, error) {
	var resolver upload.Resolver
	err := container.Resolve(&resolver)
	if err != nil {
		utils.F("Couldnt resolve UploadResolver: %v", err)
		return nil, err
	}
	var tokenModelImpl tokenModel.Token
	errResolve := container.Resolve(&tokenModelImpl)
	if errResolve != nil {
		return nil, errResolve
	}

	if !tokenModelImpl.IsTokenExist(token) {
		return nil, errors.New("invalid token")
	}

	return resolver.Resolve(ctx, images)
}

// ProcessImage is the resolver for the processImage field.
func (r *mutationResolver) ProcessImage(ctx context.Context, token string, imageID int, processes []*model.ImageProcessesInput) (*model.ImageProcess, error) {
	var resolver process.Resolver
	err := container.Resolve(&resolver)
	if err != nil {
		utils.F("Couldnt resolve Image Process Resolver: %v", err)
		return nil, err
	}
	var tokenModelImpl tokenModel.Token
	var urlModelImpl urlModel.Url

	errTokenModel := container.Resolve(&tokenModelImpl)
	errURLModel := container.Resolve(&urlModelImpl)

	if errURLModel != nil {
		return nil, errTokenModel
	}
	if errTokenModel != nil {
		return nil, errTokenModel
	}

	if !tokenModelImpl.IsTokenExist(token) {
		return nil, errors.New("invalid token")
	}
	return resolver.Resolve(ctx, imageID, processes)
}

// GetPreviewData is the resolver for the getPreviewData field.
func (r *queryResolver) GetPreviewData(ctx context.Context, token string, url string) (*model.PreviewData, error) {
	var tokenModelImpl tokenModel.Token
	var urlModelImpl urlModel.Url

	errTokenModel := container.Resolve(&tokenModelImpl)
	errURLModel := container.Resolve(&urlModelImpl)

	if errURLModel != nil {
		return nil, errTokenModel
	}
	if errTokenModel != nil {
		return nil, errTokenModel
	}

	if !tokenModelImpl.IsTokenExist(token) {
		return nil, errors.New("invalid token")
	}

	previewData, err := urlModelImpl.GetPreviewData(url)
	if err != nil {
		return nil, err
	}
	return convertor.ConvertPreviewData(previewData), nil
}

// VerifyToken is the resolver for the verifyToken field.
func (r *queryResolver) VerifyToken(ctx context.Context, token string) (*bool, error) {
	var tokenModelImpl tokenModel.Token
	err := container.Resolve(&tokenModelImpl)
	if err != nil {
		return nil, err
	}
	isExist := tokenModelImpl.IsTokenExist(token)
	return &isExist, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
