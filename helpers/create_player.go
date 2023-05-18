package helpers

type CreatePlayerRequestBody struct {
	Name     string `json:"name"`
	Grade    string `json:"grade"`
	Position string `json:"position"`
}

func (player *CreatePlayerRequestBody) IsCreatePlayerRequestBodyValid() bool {
	if player.Name != "" && isGradeValid(player.Grade) && isPositionValid(player.Position) {
		return true
	}
	return false
}

func isGradeValid(grade string) bool {
	if grade == "" {
		return false
	}
	asciiValueOfGrade := []rune(grade)
	if len(asciiValueOfGrade) > 1 {
		return false
	}
	if asciiValueOfGrade[0] < 65 || asciiValueOfGrade[0] > 67 {
		return false
	}
	return true
}

func isPositionValid(position string) bool {
	if position == "" {
		return false
	}
	validPositions := []string{
		"ST", "CF", "LW", "RW", "RS", "LS",
		"CM", "RCM", "LCM", "CAM", "LM", "RM",
		"CDM", "CB", "RCB", "LCB", "RB", "LB", "GK",
	}
	isValidFlag := false
	for _, validPosition := range validPositions {
		if position == validPosition {
			isValidFlag = true
			break
		}
	}
	return isValidFlag
}
