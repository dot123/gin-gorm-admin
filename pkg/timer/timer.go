package timer

import (
	"github.com/robfig/cron/v3"
	"sync"
)

func NewTimer() *Timer {
	return &Timer{taskList: make(map[string]*cron.Cron)}
}

// Timer 定时任务管理
type Timer struct {
	taskList map[string]*cron.Cron
	sync.Mutex
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *Timer) AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New(cron.WithSeconds())
	}
	id, err := t.taskList[taskName].AddFunc(spec, task)
	t.taskList[taskName].Start()
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *Timer) AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New(cron.WithSeconds())
	}
	id, err := t.taskList[taskName].AddJob(spec, job)
	t.taskList[taskName].Start()
	return id, err
}

// FindCron 获取对应taskName的cron 可能会为空
func (t *Timer) FindCron(taskName string) (*cron.Cron, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.taskList[taskName]
	return v, ok
}

// StartTask 开始任务
func (t *Timer) StartTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Start()
	}
	return
}

// StopTask 停止任务
func (t *Timer) StopTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Stop()
	}
	return
}

// Remove 从taskName 删除指定任务
func (t *Timer) Remove(taskName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Remove(cron.EntryID(id))
	}
	return
}

// Clear 清除任务
func (t *Timer) Clear(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Stop()
		delete(t.taskList, taskName)
	}
}

// Close 释放资源
func (t *Timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.taskList {
		v.Stop()
	}
}
