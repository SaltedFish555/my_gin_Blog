package model

import (
	"encoding/base64"
	"fmt"
	"log"
	"my_gin_blog/utils/errmsg"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model //内嵌结构体，其中包含主键及创建、更新、删除的时间四个字段
	Articles []Article `gorm:"foreignKey:Uid"`
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int; DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
}

// 查询用户名是否已被占用
func CheckUser(username string) (code int) {
	var user User
	// 查询满足条件的第一条记录，并存入user
	db.Select("id").Where("username=?", username).First(&user)
	if user.ID > 0 { // 说明用户名已被占用
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	// data.Password=ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User,int64) {
	var users []User
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == 0 && pageSize == 0 {
		offset = -1
		pageSize = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("分页错误")

		return nil,0
	}

	return users,total
}

// 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"]=data.Username
	maps["role"]=data.Role
	err:=db.Model(&user).Where("id=?",id).Updates(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 钩子函数，程序自动调用
func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	return nil
}

// 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	// salt:=make([]byte,8)
	salt := []byte{12, 32, 4, 6, 66, 22, 22, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		fmt.Println("密码加密错误")
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登录验证
func CheckLogin(username string,password string)int{
	var user User
	db.Where("username=?",username).First(&user)
	if user.ID==0{
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password)!=user.Password{
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role <1{ // 权限不正常（管理权限为1，其他权限应该>=2）
		return errmsg.ERROR_USER_NO_RIGHT
	} 
	return errmsg.SUCCESS
}

// 通过username获取uid
func GetUID(username string) (uid uint,code int){
	var user User
	err = db.Where("Username=?",username).First(&user).Error
	if err != nil {
		fmt.Println("查询用户ID错误")
		return 0,errmsg.ERROR_USER_NOT_EXIST
	}
	return user.ID,errmsg.SUCCESS

}