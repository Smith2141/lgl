package robot

import "strings"

import "fmt"

// Robot — тип робота
type Robot struct {
	Model       string
	SerialId    int
	WorkCounter int
}

func (r Robot) String() string {
	return fmt.Sprintf("Robot %s serialID %d", r.Model, r.SerialId)
}

// Work — робот выполняет работы и запоминает количество выполненных задач. Поэтому получатель метода — по указателю
func (r *Robot) Work(tasks []string) string {
	var res strings.Builder
	fmt.Fprintf(&res, "%s work:", r)
	for _, task := range tasks {
		res.WriteString("\n I do ")
		res.WriteString(task)
	}
	r.WorkCounter += len(tasks)
	return res.String()
}
