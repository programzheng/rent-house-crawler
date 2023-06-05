package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func GetJSON(value interface{}) string {
	result, err := json.Marshal(value)
	if err != nil {
		log.Fatal("helper GetJSON func:", err)
	}
	return string(result)
}

func ConvertUrlToMd5(url string) string {
	md5Hash := md5.Sum([]byte(url))
	// fmt.Printf("ConvertUrlToMd5:%v", md5Hash[:])
	md5String := hex.EncodeToString(md5Hash[:])

	return md5String
}

func ConvertCommaSeparatedStringToInt(s string) (int, error) {
	// 移除字串中的逗號
	str := strings.ReplaceAll(s, ",", "")

	// 將字串轉換為整數
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func ConvertStringToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0.0, err
	}
	return f, nil
}
