package utils

import (
	"regexp"
	"strings"
)

/**
 * 正则验证：ip
 * 1.0.0.0~255.255.255.255
 */
func CheckIp(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

/**
 * 验证pn 的url 地址是否正确
 * http://192.168.1.129:5145
 */
func CheckPNUrl(url string ) bool {
	addr :=strings.Trim(url," ")
	addr=strings.ToLower(addr)
	regStr := `^(https?://)(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(:5145/)$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}