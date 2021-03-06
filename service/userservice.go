package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) Login(loginDTO model.LoginDTO) (uint, uint, bool) {
	var user model.User
	found := obj.db.Read(&model.User{}, &model.User{UserName: loginDTO.UserName}, &user)
	if found && utility.CompareHash(loginDTO.Password, user.Password) {
		return user.ID, *user.GroupID, true
	}
	return 0, 0, false
}

//this version doesnt have interests (will be updated)
func (obj *Service) SignUp(signupDTO model.SignUpDTO) (uint, bool) {
	var user model.User
	foundName := obj.db.Read(&model.User{}, &model.User{UserName: signupDTO.Username}, &user)
	foundMail := obj.db.Read(&model.User{}, &model.User{UserMail: signupDTO.Email}, &user)
	if !foundName && !foundMail {
		newUser := model.User{
			UserName: signupDTO.Username,
			UserMail: signupDTO.Email,
			Password: utility.GetHash(signupDTO.Password),
		}
		obj.db.Create(&model.User{}, &newUser)
		return newUser.ID, true
	}
	return 0, false
}

func (obj *Service) GetGroupUsers(groupid float64) []model.UserDTO {
	var userDTOs []model.UserDTO
	gid := uint(groupid)
	obj.db.Query(&model.User{}, &model.User{GroupID: &gid}, &userDTOs)
	return userDTOs
}

func (obj *Service) GetUser(userid float64) model.UserDTO {
	var userDTO model.UserDTO
	obj.db.Query(&model.User{}, uint(userid), &userDTO)
	return userDTO
}

func (obj *Service) UpdateUser(id float64, userDTO model.UserDTO) {
	user := model.User{
		UserName:       userDTO.UserName,
		UserMail:       userDTO.UserMail,
		ProfilePicture: userDTO.ProfilePicture,
	}
	obj.db.Update(&model.User{}, uint(id), &user)
}

func (obj *Service) DeleteUser(id float64) {
	obj.db.Delete(&model.User{}, uint(id))
}
