package utility

import (
	"strings"
)

type Command struct {
	Command   string
	Arguments []string
}

func ParseMessage(message string) (*Command, error) {
	message = strings.TrimSpace(message)
	if strings.HasPrefix(message, "!lba") {
		operations := strings.Split(message, " ")
		command := Command{Command: strings.ToLower(operations[1]), Arguments: operations[2:]}
		err := validateCommand(&command)
		if err != nil {
			return &command, err
		}
		return &command, nil
	} else {
		return &Command{Command: "none"}, nil
	}
}

func validateCommand(command *Command) error {
	switch command.Command {
	case "boosters":
		if len(command.Arguments) != 0 {
			return InvalidCommandError{Command: command.Command, Message: "Does not take any argument"}
		}
		return nil
	case "collections":
		if len(command.Arguments) > 0 {
			return InvalidCommandError{Command: command.Command, Message: "Does not take any argument"}
		}
		return nil
	case "help":
		if len(command.Arguments) > 0 {
			return InvalidCommandError{Command: command.Command, Message: "Does not take any argument"}
		}
		return nil

	case "open":
		if len(command.Arguments) == 0 {
			return InvalidCommandError{Command: command.Command, Message: "Need to input the name of the booster pack"}
		}
		return nil
	case "summon":
		if len(command.Arguments) == 0 {
			return InvalidCommandError{Command: command.Command, Message: "Need to input the name of unit"}
		}
		return nil

	case "trade":
		if len(command.Arguments) != 2 {
			return InvalidCommandError{Command: command.Command, Message: "Take username of the person you wish to trade with and the card name as argument"}
		}
		return nil

	default:
		return CommandDoesNotExist{Command: command.Command}
	}
}
