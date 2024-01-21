package util

import (
	"strings"
)

func ToPkgName(s string) string {
	var ret string
	ret = strings.ToLower(s)
	ret = strings.ReplaceAll(ret, "_", "")
	ret = strings.ReplaceAll(ret, ".", "")
	return ret
}

func ToUpperCamelCase(s string) string {
	var ret string
	sliced := strings.Split(s, "_")
	for i := 0; i < len(sliced); i++ {
		substr := sliced[i]
		ret += strings.ToUpper(substr[0:1]) + substr[1:]
	}

	return ret
}

func ToLowerCamelCase(s string) string {
	var ret string
	sliced := strings.Split(s, "_")
	for i := 0; i < len(sliced); i++ {
		substr := sliced[i]
		if i == 0 {
			ret += strings.ToLower(substr)
		} else {
			ret += strings.ToUpper(substr[0:1]) + substr[1:]
		}
	}
	return ret
}

func SnakeCaseToKebabCase(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

// func ToUpperSnakeCase(s string) string {

// 	if slices.Contains([]byte(s), '_') {
// 		return strings.ToUpper(s)
// 	}

// 	var ret string
// 	start := 0
// 	for i := 0; i < len(s); i++ {
// 		c := s[i]
// 		if c >= 'A' && c <= 'Z' {
// 			if start == -1 {
// 				start = i
// 			} else {
// 				ret += strings.ToUpper(s[start:i+1]) + "_"
// 				start = -1
// 			}
// 		}
// 	}

// 	if len(ret) == 0 {
// 		return s
// 	}

// 	ret = ret[:len(ret)-1]

// 	return ret
// }

// func ToLowerSnakeCase(s string) string {

// 	if slices.Contains([]byte(s), '_') {
// 		return strings.ToLower(s)
// 	}

// 	var tmp []string
// 	var prevTmp string
// 	start := 0
// 	for i := 0; i < len(s); i++ {
// 		c := s[i]
// 		if i != 0 && c >= 'A' && c <= 'Z' {
// 			tmplen := len(tmp)
// 			if tmplen != 0 {
// 				prevTmp = tmp[tmplen-1]
// 			}
// 			if len(prevTmp) == 1 {
// 				tmp[tmplen-1] += s[start:i]
// 			} else {
// 				tmp = append(tmp, s[start:i])
// 			}
// 			// ret += s[start:i] + "_"
// 			start = i
// 		}
// 	}

// 	if len(tmp) == 0 {
// 		return strings.ToLower(s)
// 	}

// 	ret := strings.Join(tmp, "_")
// 	return strings.ToLower(ret)
// }
