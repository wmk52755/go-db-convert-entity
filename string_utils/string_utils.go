package string_utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ExplodeStr(rawStr string, seq string) []string {
	res := make([]string, 0)
	strArr := strings.Split(rawStr, seq)
	for _, str := range strArr {
		if str != "" {
			res = append(res, str)
		}
	}
	return res
}

func Implode(list interface{}, seq string) string {
	listValue := reflect.Indirect(reflect.ValueOf(list))
	if listValue.Kind() != reflect.Slice {
		return ""
	}
	count := listValue.Len()
	listStr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr = append(listStr, str)
		}
	}
	return strings.Join(listStr, seq)
}

func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}

func ReplaceAllQuotationMarks(str string) (string, map[string]string) {
	qmIsclosure := true
	lastIndex := 0
	s := str
	sQmIndex := 0
	subStrMap := make(map[string]string)
	collectNum := 0
	for {
		index := strings.Index(s, "'")
		if index < 0 {
			break
		}
		rep, err := regexp.Compile("[\\\\]*'")
		if err != nil {
			log.Println(err)
			return str, subStrMap
		}
		res := rep.Find([]byte(s))
		escapeStringNum := len(string(res)) - 1
		if escapeStringNum%2 == 0 {
			if qmIsclosure { // 引号之前是闭合的
				lastIndex += index
				qmIsclosure = false
			} else { // 引号之前是未闭合的
				if len(s) > index+2 && string(s[index+1]) == "'" {
					index++
					sQmIndex += index + 1
				} else {
					qmIsclosure = true
					variable := "$QM" + strconv.Itoa(collectNum)
					subStrMap[variable] = str[lastIndex : lastIndex+sQmIndex+index+2]
					str = str[:lastIndex] + variable + str[lastIndex+sQmIndex+index+2:]
					lastIndex = lastIndex + len(variable)
					sQmIndex = 0
					collectNum++
				}
			}
		} else {
			sQmIndex += index + 1
		}
		s = s[index+1:]
	}
	return str, subStrMap
}

func GetContentInFirstBrackets(str string) (string, error) {
	startIndex := -1
	endIndex := -1
	recordNum := 0
	for i, s := range str {
		if string(s) == "(" {
			if startIndex == -1 {
				startIndex = i
			}
			recordNum++
		} else if string(s) == ")" {
			recordNum--
			if startIndex != -1 && recordNum == 0 {
				endIndex = i
				break
			}
		}
	}
	if startIndex == -1 || endIndex == -1 {
		return str, errors.New("匹配失败，请检查字符串是否正确")
	}
	return str[startIndex+1 : endIndex], nil
}

func WriteToFile(filePath, resText string) (int, error) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return f.WriteString(resText)
}

func IsNeedClear(r rune) bool {
	if string(r) == " " {
		return true
	}
	if string(r) == "	" {
		return true
	}
	if string(r) == "\n" {
		return true
	}
	return false
}