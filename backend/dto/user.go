package dto

import "e-wallet/entity"

type UserDTO struct {
	Id                 int       `json:"id"`
	Email              string    `json:"email" binding:"required"`
	Name               string    `json:"name"`
	ProfilePictureName string    `json:"profile_picture"`
	Password           string    `json:"password,omitempty" binding:"required"`
	Wallet             WalletDTO `json:"wallet,omitempty"`
}

type UserDataDTO struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

func ToUserData(userData UserDataDTO) entity.User {
	return entity.User{
		Email: userData.Email,
		Name:  userData.Name,
	}
}

func ToUser(userDTO UserDTO) entity.User {
	return entity.User{
		Email:    userDTO.Email,
		Name:     userDTO.Name,
		Password: userDTO.Password,
	}
}

func ToUserDTO(user entity.User) UserDTO {
	return UserDTO{
		Id:                 user.Id,
		Email:              user.Email,
		Name:               user.Name,
		ProfilePictureName: user.ProfilePictureName,
		Wallet: WalletDTO{
			Id:         user.Wallet.Id,
			Number:     user.Wallet.Number,
			Balance:    user.Wallet.Balance,
			Income:     user.Wallet.Income,
			Expense:    user.Wallet.Expense,
			GachaTrial: user.Wallet.GachaTrial,
		},
	}
}
