package cmd

type CommandResult struct {
	Command    string
	ReturnCode int
	Result     interface{}
}
