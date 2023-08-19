package dto

import "Appitems/models"

// 给前端返回封装好的用户信息
func ToUserDto(user models.UserDatabase) models.UserDto {
	return models.UserDto{
		Name:           user.Username,
		Email:          user.Email,
		VerificationID: user.VerificationID,
	}
}
