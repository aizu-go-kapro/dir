//scripts/loacl.sh, remote.shそれぞれを実行するためのパッケージ
//スクリプトの内容は対象リソースを指定した場所に日時を決めて保存
//学内用とリモート用がある
package dir

import (
	"os/exec"
)

//scriptファイルのパスを設定
type Executor struct {
	Path string
}

func NewExecutor(path string) *Executor {
	executor := &Executor{Path: path}
	return executor
}

//Pathに設定されたスクリプトを実行
func (e *Executor) Do() error {
	cmd := exec.Command(e.Path)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
