package utils

import (
	"regexp"
)

/**
*** 进行手机号校验
 */
func VerifyMobilePhone(phone string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	run := regexp.MustCompile(reg)
	return run.MatchString(phone)
}
