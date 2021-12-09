package service

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) Login(loginDTO model.LoginDTO) bool {
	var login model.User
	utility.Convert(&loginDTO, &login)
	var user model.User
	found := obj.db.Read(&model.User{}, &login, &user)
	if found {
		return utility.CompareHash(loginDTO.Password, login.Password)
	}
	return false
}

//this version doesnt have interests (will be updated)
func (obj *Service) SignUp(userDTO model.UserDTO) {
	var user model.User
	utility.Convert(&userDTO, &user)
	obj.db.Create(&model.User{}, &user)
}

func (obj *Service) GetGroupUsers(id string) []model.UserDTO {
	uid := utility.ConvertID(id)
	var usersDTO []model.UserDTO
	obj.db.Query(&model.User{}, &model.User{GroupID: uid}, &usersDTO)
	return usersDTO
}

func (obj *Service) UpdateUser(id string, userDTO model.UserDTO) {
	var user model.User
	utility.Convert(&userDTO, &user)
	obj.db.Update(&model.User{}, id, &user)
}

func (obj *Service) DeleteUser(id string) {
	obj.db.Delete(&model.User{}, id)
}

/* These are unnecessary as we will only be taking DTO struct from the request and can do them all once
func (obj *Service) UpdateUserName(userDTO model.UserDTO, name string) {
	//obj.db.Update(&model.UserDTO{}, &model.UserDTO{ID: userDTO.ID}, &model.UserDTO{UserName: name})
	//obj.db.Update(&model.User{}, &model.User{Model: gorm.Model{ID: userDTO.ID}}, &model.User{UserName: name})

}

func (obj *Service) UpdateProfilePicture(userDTO model.UserDTO, picture string) {
	//obj.db.Update(&model.UserDTO{}, &model.UserDTO{ID: userDTO.ID}, &model.UserDTO{ProfilePicture: picture})
	//obj.db.Update(&model.User{}, &model.User{Model: gorm.Model{ID: userDTO.ID}}, &model.User{ProfilePicture: picture})
}
*/
