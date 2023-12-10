package helper

import "errors"

func ScoringMatrix(probability string, score int) (int, string, error) {
	var categoryNum int
	var categoryStr string

	switch probability {
	case "High":
		switch {
		case score >= 0 && score <= 15:
			categoryNum, categoryStr = 2, "Mitigate/Defer"
		case score >= 16 && score <= 29:
			categoryNum, categoryStr = 2, "Mitigate/Defer"
		case score >= 30 && score <= 45:
			categoryNum, categoryStr = 1, "Mitigate"
		default:
			return 0, "", errors.New("invalid score for High probability")
		}
	case "Medium":
		switch {
		case score >= 0 && score <= 15:
			categoryNum, categoryStr = 3, "Defer/Accept"
		case score >= 16 && score <= 29:
			categoryNum, categoryStr = 2, "Mitigate/Defer"
		case score >= 30 && score <= 45:
			categoryNum, categoryStr = 2, "Mitigate/Defer"
		default:
			return 0, "", errors.New("invalid score for Medium probability")
		}
	case "Low":
		switch {
		case score >= 0 && score <= 15:
			categoryNum, categoryStr = 4, "Accept"
		case score >= 16 && score <= 29:
			categoryNum, categoryStr = 3, "Defer/Accept"
		case score >= 30 && score <= 45:
			categoryNum, categoryStr = 3, "Defer/Accept"
		default:
			return 0, "", errors.New("invalid score for Low probability")
		}
	default:
		return 0, "", errors.New("invalid probability")
	}

	return categoryNum, categoryStr, nil
}
