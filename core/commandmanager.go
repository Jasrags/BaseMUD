package core

type CommandManager struct {
}

func NewCommandManager() *CommandManager {
	return &CommandManager{}
}

func (cm *CommandManager) Add(cmd Command) {
}

func (cm *CommandManager) Find(search string, returnAlias bool) Command {
	return Command{}
}

func (cm *CommandManager) Get(cmd string) Command {
	return Command{}
}

func (cm *CommandManager) Remove(cmd Command) {
}
