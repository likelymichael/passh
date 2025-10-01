package prompt

import (
	"github.com/manifoldco/promptui"
)

// Confirmation prompt to delete a login item.
func ConfirmLoginItemDelete() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Confirm deletion of the login item",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
