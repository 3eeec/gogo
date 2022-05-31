package models

import (
	"strconv"
)

type User_role struct{
	UserId string `json:"user_id"`
	RoleId uint `json:"role_id"`
}
func FindRoleByUserId(userId interface{})User_role{
	var uRole User_role
	DB.Where("user_id = ?", userId).First(&uRole)
	return uRole
}
func CreateUserRole(userId uint,roleId uint){
	uRole:=&User_role{
		UserId:strconv.Itoa(int(userId)),
		RoleId: roleId,
	}
	DB.Create(uRole)
}