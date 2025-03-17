package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// CustomValidator เป็น struct ที่ครอบ validator instance และ translator
type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

// ValidationError เป็น struct ที่เก็บข้อมูลข้อผิดพลาดจากการตรวจสอบ
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// ValidationErrors เป็น slice ของ ValidationError
type ValidationErrors []ValidationError

// Error implements the error interface
func (ve ValidationErrors) Error() string {
	var errMsgs []string
	for _, err := range ve {
		errMsgs = append(errMsgs, fmt.Sprintf("Field '%s': %s", err.Field, err.Message))
	}
	return strings.Join(errMsgs, ", ")
}

// New สร้าง CustomValidator instance ใหม่พร้อมกับการตั้งค่า
func New() *CustomValidator {
	// สร้าง validator instance
	validate := validator.New()

	// ตั้งค่าให้ใช้ชื่อฟิลด์จาก json tag
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// ตั้งค่า translator
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	// ลงทะเบียน custom validation
	registerCustomValidations(validate)

	return &CustomValidator{
		validator:  validate,
		translator: trans,
	}
}

// registerCustomValidations ลงทะเบียน custom validation functions
func registerCustomValidations(validate *validator.Validate) {
	// ตัวอย่าง custom validation
	_ = validate.RegisterValidation("nowhitespace", validateNoWhitespace)
	
	// สามารถเพิ่ม custom validation อื่นๆ ได้ตามต้องการ
}

// validateNoWhitespace ตรวจสอบว่าไม่มีช่องว่างในข้อความ
func validateNoWhitespace(fl validator.FieldLevel) bool {
	return !strings.ContainsAny(fl.Field().String(), " \t\n")
}

// Validate ตรวจสอบข้อมูลตาม struct tag
func (cv *CustomValidator) Validate(i interface{}) ValidationErrors {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	// แปลง validator.ValidationErrors เป็น ValidationErrors
	validationErrors := ValidationErrors{}
	for _, err := range err.(validator.ValidationErrors) {
		var element ValidationError
		element.Field = err.Field()
		element.Tag = err.Tag()
		element.Value = err.Param()
		element.Message = err.Translate(cv.translator)
		validationErrors = append(validationErrors, element)
	}

	return validationErrors
}

// ValidateVar ตรวจสอบตัวแปรเดี่ยวๆ
func (cv *CustomValidator) ValidateVar(field interface{}, tag string) error {
	return cv.validator.Var(field, tag)
}

// GetValidator คืนค่า validator instance
func (cv *CustomValidator) GetValidator() *validator.Validate {
	return cv.validator
}

// RegisterCustomValidation ลงทะเบียน custom validation functions
func (cv *CustomValidator) RegisterCustomValidation(tag string, fn validator.Func) error {
	return cv.validator.RegisterValidation(tag, fn)
}