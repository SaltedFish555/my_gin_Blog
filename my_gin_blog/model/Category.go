package model

import (
	"fmt"
	"my_gin_blog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct{
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Articles []Article `gorm:"foreignKey:Cid"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}



// 查询分类名是否已被占用
func CheckCategory(username string) (code int) {
	var category Category
	// 查询满足条件的第一条记录，并存入user
	db.Select("id").Where("name=?", username).First(&category)
	if category.ID > 0 { // 说明用户名已被占用
		return errmsg.ERROR_CATENAME_USED //3001
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) int {
	// data.Password=ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategories(pageSize int, pageNum int) ([]Category ,int64){
	var categories []Category
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == 0 && pageSize == 0 {
		offset = -1
		pageSize = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&categories).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("分页错误")
		return nil,total
	}
	return categories,total
}

// 编辑分类信息
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"]=data.Name
	err:=db.Model(&category).Where("id=?",id).Updates(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// 删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id=?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}


// 查询分类下的所有文章
