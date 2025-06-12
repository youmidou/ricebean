package sys_timer

import (
	"fmt"
	"time"
)

type TimerTest struct{}

func (t *TimerTest) Do() {

	time.Sleep(5 * time.Second) // Let the tasks stop

	fmt.Println("Tasks finished")

}
func (t *TimerTest) DoTest() {

}
