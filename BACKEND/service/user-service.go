package service

import (
	"learning-go/dto"
	"learning-go/entity"
	"learning-go/repository"
	"log"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userId string) entity.User
}

type userSerivce struct {
	userRepository repository.UserRepository
}

func NewUserService(u repository.UserRepository) UserService {
	return &userSerivce{
		userRepository: u,
	}
}

func (service *userSerivce) Update(user dto.UserUpdateDTO) entity.User {
	userUpdate := entity.User{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map %v:", err)
	}
	updateUser := service.userRepository.UpdateUser(userUpdate)
	return updateUser
}
func (service *userSerivce) Profile(userId string) entity.User {
	return service.userRepository.ProfileUser(userId)
}
