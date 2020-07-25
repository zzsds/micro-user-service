package service

import (
	"fmt"
	"log"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	// Validate ...
	Validate *Valid
)

// Valid ...
type Valid struct {
	ut.Translator
	*validator.Validate
}

// NewValid ...
func NewValid(translator locales.Translator) *Valid {
	var (
		validate = validator.New()
		trans    ut.Translator
		err      error
	)
	uni := ut.New(translator)
	trans, _ = uni.GetTranslator(translator.Locale())
	//验证器注册翻译器
	switch trans.Locale() {
	case "zh":
		err = zh_translations.RegisterDefaultTranslations(validate, trans)
	}
	if err != nil {
		log.Printf("Translator registration failed: %v", err)
	}
	return &Valid{
		trans,
		validate,
	}
}

// FirstError 第一个错误
func (v *Valid) FirstError(err error) error {
	if err == nil {
		return nil
	}
	for _, err := range err.(validator.ValidationErrors) {
		return fmt.Errorf(err.Translate(v.Translator))
	}
	return nil
}

// Errors 翻译后错误列表
func (v *Valid) Errors(err error) map[string]string {
	if err == nil {
		return nil
	}
	list := make(map[string]string, 0)
	for _, err := range err.(validator.ValidationErrors) {
		list[err.Field()] = err.Translate(v.Translator)
		continue
	}
	return list
}
