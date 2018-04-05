/**
 * 通用工具类
 */
package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"math/rand"
	"encoding/base64"
	"encoding/json"
	"strconv"
)

/**
 * @brief 生成32位MD5
 * @param string data 需要加密的字符
 * @return mixed
 */
func Md5(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

/**
 * @brief 生成随机字符串-数字加字母
 * @param length int 固定长度
 * @return mixed
 */
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * @brief 生成随机字符串-纯数字
 * @param length int 固定长度
 * @return mixed
 */
func GetRandomIntString(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * @brief base64编码
 * @param string name 待编码的字符
 * @return mixed
 */
func Base64Encode(data string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(data))

	return encodeString
}

/**
 * @brief base64解码
 * @param string encodeString base64编码的字符
 * @return mixed
 */
func Base64Decode(encodeString string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)

	if err == nil {
		return string(decodeBytes)
	} else {
		return ""
	}
}

/**
 * @brief json编码
 * @param mixed data 需要json编码的数据
 * @return mixed
 */
func JsonEncode(data interface{}) string {
	jsonEncodeRes,err := json.Marshal(data)

	if err == nil {
		return string(jsonEncodeRes)
	} else {
		return ""
	}
}

/**
 * @brief json解码
 * @param string data 待解码的字符串
 * @return mixed
 */
func JsonDecode(encodeString string) map[string]interface{} {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(encodeString), &data)
	if err == nil {
		return data
	} else {
		return nil
	}
}

/**
 * @brief int转string
 * @param int nu 需要转为string的整型
 * @return mixed
 */
func IntToString(nu int) string {
	res := strconv.Itoa(nu)
	return res
}

/**
 * @brief string转int
 * @param string str 需要转换为int的字符
 * @return mixed
 */
func StringToInt(str string) int {
	i,err := strconv.Atoi(str)

	if err == nil {
		return i
	} else {
		return 0
	}
}