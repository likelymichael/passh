package prompt

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

var promptsWizard = []func() error{
	getPass,
}

var validate = func(input string) error {
	if len(input) < 12 {
		return errors.New("Password must have 12 or more characters.")
	}
	return nil
}

func GetMasterPassword() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Enter your Passh master password:",
		IsConfirm: false,
		Mask:      '*',
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func WelcomeWizard() error {
	for _, prompt := range promptsWizard {
		if err := prompt(); err != nil {
			return err
		}
	}

	return nil
}

func getPass() error {
	prompt := promptui.Prompt{
		Label:    "Create a Passh Password",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	persist, persistErr := persistPass()
	if persistErr != nil {
		return persistErr
	}

	if persist == "y" || persist == "Y" {
		fmt.Print(result)
	} else {
		// for some reason, this isn't setting even though it'll set above
		// will come back to this, but for now i can just document setting it after wizard manually
		// config.SaveConfigValue("auth", "persist_pass", "")
	}

	return nil
}

func persistPass() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Enable persistent password? If enabled, Passh will not timeout and prompt to re-enter your master password. This can be changed later in your config.ini",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
