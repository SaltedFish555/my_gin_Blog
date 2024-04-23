package validator

import (
	"fmt"
	"my_gin_blog/utils/errmsg"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{})(string,int){
	// 传interface而不是结构体是因为struct是值传递
	validate:=validator.New()
	uni:=unTrans.New(zh_Hans_CN.New())
	trans,_:=uni.GetTranslator("zh_Hans_CN")

	err:=zhTrans.RegisterDefaultTranslations(validate,trans)
	if err!=nil{
		fmt.Println("注册翻译方法错误",err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField)string{
		label:=field.Tag.Get("label")
		return label
	})
	// 由于一定是struct，所以可以不断言
	err=validate.Struct(data)
	if err!=nil{
		// .() 断言
		for _,v:=range err.(validator.ValidationErrors){
			return v.Translate(trans),errmsg.ERROR
		}
	}
	return "",errmsg.SUCCESS

}




















