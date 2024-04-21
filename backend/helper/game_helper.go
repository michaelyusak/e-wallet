package helper

import (
	"strconv"

	"e-wallet/apperror"
	"e-wallet/constants"
	"e-wallet/entity"
)

func CheckForGacha(selectionStr string) (int, error) {
	selectionInt, err := strconv.Atoi(selectionStr)
	if err != nil {
		return 0, apperror.BadRequest("selection not number")
	}

	if selectionInt < constants.MinBoxNumber || selectionInt > constants.MaxBoxNumber {
		return 0, apperror.BadRequest("user must select one of 9 boxes")
	}

	return selectionInt, nil
}

func SelectBox(gachaBoxes []entity.GachaBox, selection int) entity.GachaBox {
	selectedBox := gachaBoxes[selection-1]
	selectedBox.BoxNumber = selection

	return selectedBox
}
