package access_token

import (
	"oauthService/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}
type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(token AccessToken) *errors.RestErr {
	err := token.Validate()
	if err != nil {
		return err
	}
	return s.repository.Create(token)
}

func (s *service) UpdateExpirationTime(token AccessToken) *errors.RestErr {
	err := token.Validate()
	if err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(token)
}
