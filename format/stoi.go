package format

import (
	"strconv"
	"strings"
)

//StrToIntArr return 将逗号分割的字符串 转换为 int型数组
func StrToIntArr(str string) ([]int, error) {
	strs := strings.Split(str, ",")

	tois := make([]int, 0, len(strs))

	for _, val := range strs {
		toi, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		tois = append(tois, toi)
	}

	return tois, nil
}

//StrArrToIntArr return 将字符串数组 转换为 int型数组
func StrArrToIntArr(strs []string) ([]int, error) {
	tois := make([]int, 0, len(strs))

	for _, val := range strs {
		toi, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		tois = append(tois, toi)
	}

	return tois, nil
}

//IntArrToStrArr return 将int型数组  转换为 字符串数组
func IntArrToStrArr(params []int) []string {
	res := make([]string, 0, len(params))

	for _, val := range params {
		res = append(res, strconv.Itoa(val))
	}

	return res
}
