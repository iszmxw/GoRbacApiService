package helpers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"gorbac/pkg/config"
	"io"
)

const (
	base64Table = "bbUGUNiuy54851asaGYUGJBiugubnjkhu78584215jhbgbkjbuy+/"
)

func Md5(password string) string {
	year := "2021"
	w := md5.New()
	// 拼接密码
	str := password + config.Env("PW_SALT", "iszmxw").(string)
	_, _ = io.WriteString(w, str) // 将str写入到w中
	str1 := fmt.Sprintf("%x", w.Sum(nil)) + year
	w1 := md5.New()
	_, _ = io.WriteString(w1, str1)       // 将str写入到w中
	return fmt.Sprintf("%x", w1.Sum(nil)) // w.Sum(nil)将w的hash转成[]byte格式
}

func Base64Encode(src []byte) []byte { //编码
	return []byte(base64.NewEncoding(base64Table).EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) { //解码
	return base64.NewEncoding(base64Table).DecodeString(string(src))
}
