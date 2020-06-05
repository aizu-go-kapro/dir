//授業のある曜日にプログラム実行するためのcron用のパッケージ
package dir

//Schedulerは実行するjobと実行日を持たせる
type Scheduler struct {
	*Executor
	Date []string
}

func (j *Scheduler) Do() error {
	//TODO: cron job
	return nil
}
