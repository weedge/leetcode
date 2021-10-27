package main


// 定义匹配关系
// 不匹配入栈
// 匹配出栈
// 栈为空为有效
func isValid(s string) bool {
	if len(s)==0 || len(s)%2 == 1{
		return false
	}
	pairs := map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}
	a := []rune{}
	for _, i:=range s{
		if pairs[i]>0{
			if len(a)==0 || a[len(a)-1] != pairs[i]{
				return false
			}
			a = a[:len(a)-1]
		}else{
			a = append(a, i)
		}
	}
	return len(a)==0
}
