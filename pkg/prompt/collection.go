package prompt

import (
	"github.com/manifoldco/promptui"
)

// Confirmation prompt for deletion of a collection.
func ConfirmCollectionDelete() (string, error) {
	prompt := promptui.Prompt{
		Label:     "WARNING: DELETES ALL LOGINS IN COLLECTION. Confirm deletion",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
