package model

import (
	"fmt"
	"my_gin_blog/utils/errmsg"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
	Cid      uint     `gorm:"type:int;not null" json:"cid"`
	Uid      uint     `gorm:"type:int;not null" json:"uid"`
}

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// 查询分类下的所有文章
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var cate Category
	var total int64
	// offset := (pageNum - 1) * pageSize
	// if pageNum == 0 && pageSize == 0 {
	// 	offset = -1
	// 	pageSize = -1
	// }
	err := db.Preload("Articles").First(&cate,cid).Error //查找id=cid的分类，并预加载它的Articles
	if err != nil {
		fmt.Printf("在查找id=%d的分类时错误",cid)
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	db.Model(&cate).Association("Articles").Find(&cateArtList) // 把这个分类里的Articles传给userArtList

	return cateArtList, errmsg.SUCCESS, total

}

// 查询某用户下的所有文章
func GetUserArt(uid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var userArtList []Article
	var user User
	var total int64=5
	// offset := (pageNum - 1) * pageSize
	// if pageNum == 0 && pageSize == 0 {
	// 	offset = -1
	// 	pageSize = -1
	// }
	// err := db.Preload("Articles").Limit(pageSize).Offset(offset).Where("uid=?", uid).Find(&userArtList).Count(&total).Error
	err := db.Preload("Articles").First(&user,uid).Error //查找id=uid的用户，并预加载它的Articles
	if err!=nil{
		fmt.Printf("查找id=%d的用户时错误\n",uid)
		return nil, errmsg.ERROR_USER_NOT_EXIST, 0
	}
	db.Model(&user).Association("Articles").Find(&userArtList) // 把这个用户里的Articles传给userArtList
	return userArtList, errmsg.SUCCESS, total

}


// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var article Article
	err := db.First(&article,id).Error
	if err != nil {
		fmt.Println("查询单个文章错误")
		return article, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCESS

}

// 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int, int64) {
	var articles []Article
	var total int64

	offset := (pageNum - 1) * pageSize
	if pageNum == 0 && pageSize == 0 {
		offset = -1
		pageSize = -1
	}
	err := db.Limit(pageSize).Offset(offset).Find(&articles).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("分页错误")
		return nil, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, total
}


// 查询推荐文章(最近三天内发布的文章)
func GetRecomArt() ([]Article, int, int64) {
	var articles []Article
	var total int64


	// 计算三天前的时间
	threeDaysAgo := time.Now().AddDate(0, 0, -3)
	err := db.Where("created_at >= ?", threeDaysAgo).Find(&articles).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("查找推荐文章错误")
		return nil, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, total


}

// 编辑文章信息（bug:没检查文章id是否存在）
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&article).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id=?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
