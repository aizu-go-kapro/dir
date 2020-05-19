package dir

import (
	"testing"
)

const (
	TestCommand = "pwd"
)

var (
	TestArgs = []string{"-L", "-P"}
)

func TestSetCommand(t *testing.T) {
	testExecutor := new(Executor)
	optionFunc := SetCommand(TestCommand)
	optionFunc(testExecutor)
	if testExecutor.ScriptPath != TestCommand {
		t.Errorf("[ERROR]: testExecutor.ScriptPath must be %s. But it's %s\n", TestCommand, testExecutor.ScriptPath)
	}
}

func TestSetArgs(t *testing.T) {
	testExecutor := new(Executor)
	optionFunc := SetArgs(TestArgs...)
	optionFunc(testExecutor)
	for i := 0; i < len(TestArgs); i++ {
		if testExecutor.Args[i] != TestArgs[i] {
			t.Errorf("[ERROR]: testExecutor.Args must be %s. But it's %s\n", TestArgs, testExecutor.ScriptPath)
		}
	}
}

func TestNewExecutor(t *testing.T) {
	testExecutor := NewExecutor(
		SetCommand(TestCommand),
		SetArgs(TestArgs...),
	)
	if testExecutor.ScriptPath != TestCommand {
		t.Errorf("[ERROR]: testExecutor.ScriptPath must be %s. But it's %s\n", TestCommand, testExecutor.ScriptPath)
	}
	for i := 0; i < len(TestArgs); i++ {
		if testExecutor.Args[i] != TestArgs[i] {
			t.Errorf("[ERROR]: testExecutor.Args must be %s. But it's %s\n", TestArgs, testExecutor.ScriptPath)
		}
	}
}

func TestDo(t *testing.T) {
	executor := NewExecutor(
		SetCommand(TestCommand),
		SetArgs(TestArgs...),
	)

	if err := executor.Do(); err != nil {
		t.Error(err)
	}
}
