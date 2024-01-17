package sprint

import "strings"

func Season(month string) string {
	lowerMonth := strings.ToLower(month)

	switch lowerMonth {
	case "jan", "feb", "dec":
		return "winter\ninvalid input"
	case "mar", "apr", "may":
		return "spring\ninvalid input"
	case "jun", "jul", "aug":
		return "summer\ninvalid input"
	case "sep", "oct", "nov":
		return "autumn\ninvalid input"
	default:
		return "invalid input: " + month
	}
}
