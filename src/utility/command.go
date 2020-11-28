package utility

import (
	"strings"
)

type Command struct {
	Command   string
	Arguments []string
}

func ParseMessage(message string) (*Command, error) {
	if strings.HasPrefix(message, "!lba") {
		operations := strings.Split(message, " ")
		command := Command{Command: strings.ToLower(operations[0]), Arguments: operations[1:]}
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
	case "help":
		if len(command.Arguments) > 0 {
			return InvalidCommandError{Command: command.Command, Message: "Does not take any argument"}
		}
		return nil
	case "collections":
		if len(command.Arguments) > 0 {
			return InvalidCommandError{Command: command.Command, Message: "Does not take any argument"}
		}
		return nil
	case "summon":
		if len(command.Arguments) != 1 {
			return InvalidCommandError{Command: command.Command, Message: "Only take the name of the card as an argument"}
		}
		return nil
	case "trade":
		if len(command.Arguments) != 2 {
			return InvalidCommandError{Command: command.Command, Message: "Take username of the person you wish to trade with and the card name as argument"}
		}
		return nil
	case "open":
		if len(command.Arguments) != 1 {
			return InvalidCommandError{Command: command.Command, Message: "Take the booster name as an argument"}
		}
		return nil
	case "booster":
		if len(command.Arguments) != 0 {
			return InvalidCommandError{Command: command.Command, Message: "Does not take any argument"}
		}
		return nil
	default:
		return CommandDoesNotExist{Command: command.Command}
	}
}
