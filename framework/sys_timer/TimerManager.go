package sys_timer

/*
		定时器
	  - author yh @ 2023-11-1 17:58
	    *注意:
*/
/*
type TimerTask struct {
	taskID    int
	interval  time.Duration
	taskFunc  func(taskId int)
	stopChan  chan bool
	isRunning bool
}

func (t *TimerTask) Start() {
	if t.interval == 0 {
		t.doTaskFunc()
		return
	}
	GetTimerManagerInstance().AddTask(t)
}
func (t *TimerTask) Stop() {
	GetTimerManagerInstance().DeleteTask(t.taskID)
}
func (t *TimerTask) doTaskFunc() {
	defer func() {
		if r := recover(); r != nil {
			err := debug.Stack()
			log4.Error("TimerManager->taskFunc:%s  %s", r, err)
			t.Stop()
		}
	}()
	t.taskFunc(t.taskID)
}
*/
/*
		全局定时管理器
	  - author yh @ 2023-11-1 17:55
	    *注意:
*/
/*
type TimerManager struct {
	tasks       map[int]*TimerTask
	addTaskChan chan *TimerTask
	delTaskChan chan int
	index       int
}

func NewTimerTask(interval time.Duration, taskFunc func(taskId int)) *TimerTask {
	return &TimerTask{
		taskID:    -1,
		interval:  interval,
		taskFunc:  taskFunc,
		stopChan:  make(chan bool),
		isRunning: false,
	}
}

var instance *TimerManager
var once sync.Once

func GetTimerManagerInstance() *TimerManager {
	once.Do(func() {
		instance = _NewTimerManager() //&TimerManager{}
	})
	return instance
}

func _NewTimerManager() *TimerManager {
	manager := &TimerManager{
		tasks:       make(map[int]*TimerTask),
		addTaskChan: make(chan *TimerTask),
		delTaskChan: make(chan int),
		index:       0,
	}
	go manager.start()
	return manager
}

func (t *TimerManager) AddTask(task *TimerTask) {
	task.taskID = GetTimerManagerInstance().GetIndexId()
	t.addTaskChan <- task
}

func (t *TimerManager) DeleteTask(taskID int) {
	t.delTaskChan <- taskID
}

func (t *TimerManager) start() {
	for {
		select {
		case task := <-t.addTaskChan:
			if !task.isRunning {
				task.isRunning = true
				go t.runTask(task)
				t.tasks[task.taskID] = task
			}
		case taskID := <-t.delTaskChan:
			if task, ok := t.tasks[taskID]; ok {
				task.stopChan <- true
				delete(t.tasks, taskID)
			}
		}
	}
}

func (t *TimerManager) runTask(task *TimerTask) {
	ticker := time.NewTicker(task.interval)
	for {
		select {
		case <-ticker.C:
			go task.doTaskFunc()
		case <-task.stopChan:
			ticker.Stop()
			return
		}
	}
}

func (t *TimerManager) GetIndexId() int {
	t.index++
	return t.index
}
*/
