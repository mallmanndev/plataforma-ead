package services

import "github.com/google/uuid"

type UUIDService struct {
}

func NewUUIDService() *UUIDService {
	return &UUIDService{}
}

func (u *UUIDService) Generate() string {
	return uuid.NewString()
}
