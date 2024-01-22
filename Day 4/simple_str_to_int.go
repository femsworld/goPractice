package sprint

func SimpleStrToInt(s string) int {
	var num int
	for _, digit := range s {
		digitValue := int(digit - '0')
		if digitValue >= 0 && digitValue <= 9 {
			num = num * 10 + digitValue
		} else {
			return 0
		}
	}
	return num

}