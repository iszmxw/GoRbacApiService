package service

import (
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var Validate *validator.Validate
var trans ut.Translator

func init() {
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	err := zh_translations.RegisterDefaultTranslations(Validate, trans)
	if err != nil {
		return
	}
}
//Translate 翻译工具
func Translate(err error, s interface{}) map[string]string {
	r := make(map[string]string)
	t := reflect.TypeOf(s)
	for _, err := range err.(validator.ValidationErrors) {
		//使用反射方法获取struct种的json标签作为key --重点2
		var k string
		if field, ok := t.FieldByName(err.StructField()); ok {
			k = field.Tag.Get("json")
		}
		if k == "" {
			k = err.StructField()
		}
		r[k] = err.Translate(trans)
	}
	return r
}
