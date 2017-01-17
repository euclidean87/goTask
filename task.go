package goTask

import (
	"log"
	"reflect"
	"time"
)

type Task struct {
	StartTime time.Time
	Duration  time.Duration
	Func      interface{}
	Params    []reflect.Value
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) SetTaskTime(taskTime time.Time) *Task {
	t.StartTime = taskTime
	return t
}

func (t *Task) SetDuration(d time.Duration) *Task {
	t.Duration = d
	return t
}

func (t *Task) SetFunc(f interface{}) *Task {
	t.Func = f
	return t
}

func (t *Task) SetParams(params ...interface{}) *Task {
	t.Params = getFuncParams(params)
	return t
}

func getFuncParams(params ...interface{}) []reflect.Value {
	v := make([]reflect.Value, len(params))
	for index, p := range params {
		v[index] = reflect.ValueOf(p)
	}
	return v
}

func (t *Task) Run() {
	f := reflect.ValueOf(t.Func)
	if f.Kind() != reflect.Func {
		log.Fatal("Task.Func type is invalid")
	}
	go func() {
		nextTime, tic := start(t.StartTime, t.Duration)
		for {
			<-tic.C
			f.Call(t.Params)
			nextTime, tic = start(nextTime, t.Duration)
		}
	}()
}

func start(t time.Time, d time.Duration) (time.Time, *time.Ticker) {
	// 第一次执行时间
	if !taskTime.After(time.Now()) {
		if !taskTime.Add(d).After(time.Now()) {
			taskTime = time.Now().Add(d)
		} else {
			taskTime = taskTime.Add(d)
		}
	}
	fmt.Println(taskTime)
	// 当前时间和下一次执行时间差
	diff := taskTime.Sub(time.Now())
	return taskTime, time.NewTicker(diff)
}
