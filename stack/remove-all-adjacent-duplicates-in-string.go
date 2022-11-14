package main

func removeDuplicates(s string) string {
	if len(s) == 0 {
		return ""
	}
	var res []byte
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if len(res) == 0 || res[len(res)-1] != str[i] {
			res = append(res, str[i])
		} else {
			res = res[:len(res)-1]
		}
	}
	return string(res)
}
