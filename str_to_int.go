package sprint

func StrToInt(s string) int {
	var num int
	var sign int = 1
	for i, char := range s {
		switch i {
		case 0:
			if char == '-' {
				sign = -1
			} else if char == '+' {
				sign = 1
			} else {
				num = num*10 + int(char-'0')
			}
		default:
			if char < '0' || char > '9' {
				return 0
			}
			num = num*10 + int(char-'0')
		}
	}
	return num * sign
}
