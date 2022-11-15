package caputils

type CommandResult struct {
	Command    string
	ReturnCode int
	Result     interface{}
}
