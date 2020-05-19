//scripts/loacl.sh, remote.shそれぞれを実行するためのパッケージ
//スクリプトの内容は対象リソースを指定した場所に日時を決めて保存
//学内用とリモート用がある
package dir

import (
	"log"
	"os/exec"
)

type Executor struct {
	ScriptPath string
	Args       []string
}

func NewExecutor(options ...func(*Executor)) *Executor {
	executor := new(Executor)
	for _, option := range options {
		option(executor)
	}
	return executor
}

func SetCommand(src string) func(*Executor) {
	return func(e *Executor) {
		e.ScriptPath = src
	}
}

func SetArgs(args ...string) func(*Executor) {
	return func(e *Executor) {
		for _, arg := range args {
			e.Args = append(e.Args, arg)
		}
	}
}

func (e *Executor) Do() error {
	cmd := exec.Command(e.ScriptPath, e.Args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	log.Printf("[STDOUT]: %s", string(out))
	return nil
}
