package dir

const (
	SUNDAY = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SUTURDAY
)

type Cron struct{}

//option for functional options pattern
//type option func(*Cron)

func NewCronJob(options ...func(*Cron)) (*Cron, error) {
	return nil, nil
}

func SetDay(day uint) func(*Cron) {
	return func(c *Cron) {}
}

func (c *Cron) Do() error {
	return nil
}
