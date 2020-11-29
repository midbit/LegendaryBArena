package utility

import "fmt"

type InvalidCommandError struct {
	Command string
	Message string
}

func (e InvalidCommandError) Error() string {
	return fmt.Sprintf("Invalid %s Error: %s", e.Command, e.Message)
}

type CommandDoesNotExist struct {
	Command string
}

func (e CommandDoesNotExist) Error() string {
	return fmt.Sprintf("Command %s does not exist", e.Command)
}

type CardDoesNotExistError struct {
	Card string
}

func (e CardDoesNotExistError) Error() string {
	return fmt.Sprintf("Card %s does not exist in your collections", e.Card)
}
