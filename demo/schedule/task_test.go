package schedule

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

// Redis configuration
var (
	ctx                   = context.Background()
	lockTTL               = 5 * time.Second       // Lock expiration time
	taskExecutionInterval = 10 * time.Second      // Task check interval
	redisAddr             = "192.168.3.112:16378" // Redis address
	password              = "beta-redis-0"
	db                    = 85
)

// TaskManager struct
type TaskManager struct {
	cron        *cron.Cron
	tasks       map[cron.EntryID]*Task
	taskMutex   sync.Mutex
	redisClient *redis.Client
}

// Task struct
type Task struct {
	ID       cron.EntryID
	Spec     string
	Function func()
	LockKey  string // Redis lock key to uniquely identify the task
}

// Initialize TaskManager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		cron:  cron.New(cron.WithSeconds()), // Use cron expressions with seconds
		tasks: make(map[cron.EntryID]*Task),
		redisClient: redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: password,
			DB:       db,
		}),
	}
}

// Start the task scheduler
func (tm *TaskManager) Start() {
	tm.cron.Start()
}

// Stop the task scheduler
func (tm *TaskManager) Stop() {
	tm.cron.Stop()
	tm.redisClient.Close()
}

// Try to acquire a distributed lock
func (tm *TaskManager) acquireLock(lockKey string) bool {
	success, err := tm.redisClient.SetNX(ctx, lockKey, "1", lockTTL).Result()
	if err != nil {
		log.Println("Error acquiring lock:", err)
		return false
	}
	return success
}

// Release the distributed lock
func (tm *TaskManager) releaseLock(lockKey string) {
	_, err := tm.redisClient.Del(ctx, lockKey).Result()
	if err != nil {
		log.Println("Error releasing lock:", err)
	}
}

// Create a new task
func (tm *TaskManager) CreateTask(spec string, function func(), lockKey string) (cron.EntryID, error) {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	id, err := tm.cron.AddFunc(spec, func() {
		// Use distributed lock to ensure only one instance executes the task
		if tm.acquireLock(lockKey) {
			defer tm.releaseLock(lockKey)
			function()
		}
	})
	if err != nil {
		return 0, err
	}

	tm.tasks[id] = &Task{
		ID:       id,
		Spec:     spec,
		Function: function,
		LockKey:  lockKey,
	}
	fmt.Println("Task created with ID:", id)
	return id, nil
}

// Update an existing task
// Update an existing task
func (tm *TaskManager) UpdateTask(id cron.EntryID, newSpec string, newFunction func(), newLockKey string) error {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	// Check if the task exists
	oldTask, exists := tm.tasks[id]
	if !exists {
		return fmt.Errorf("task with ID %d does not exist", id)
	}

	// Release old lock and remove old task
	tm.releaseLock(oldTask.LockKey)
	tm.cron.Remove(id)
	delete(tm.tasks, id)

	// Create new task
	newID, err := tm.cron.AddFunc(newSpec, func() {
		// Use distributed lock to ensure only one instance executes the task
		if tm.acquireLock(newLockKey) {
			defer tm.releaseLock(newLockKey)
			newFunction()
		}
	})
	if err != nil {
		return err
	}

	// Update the task map with the new task ID
	tm.tasks[newID] = &Task{
		ID:       newID,
		Spec:     newSpec,
		Function: newFunction,
		LockKey:  newLockKey,
	}
	fmt.Println("Task updated with new ID:", newID)
	return nil
}

// Delete a task
func (tm *TaskManager) DeleteTask(id cron.EntryID) error {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	task, exists := tm.tasks[id]
	if exists {
		tm.releaseLock(task.LockKey) // Release lock
	}
	tm.cron.Remove(id)
	delete(tm.tasks, id)
	fmt.Println("Task deleted with ID:", id)
	return nil
}

// View all tasks
func (tm *TaskManager) ViewTasks() {
	tm.taskMutex.Lock()
	defer tm.taskMutex.Unlock()

	fmt.Println("Current scheduled tasks:")
	for id, task := range tm.tasks {
		fmt.Printf("ID: %d, Spec: %s, LockKey: %s\n", id, task.Spec, task.LockKey)
	}
}

func Test_task(t *testing.T) {
	// Create TaskManager and start it
	tm := NewTaskManager()
	tm.Start()
	defer tm.Stop()

	// Create Task 1
	taskID, _ := tm.CreateTask("*/5 * * * * *", func() {
		fmt.Println("Executing Task 1 - every 5 seconds:", time.Now())
	}, "task1-lock")

	// Create Task 2
	tm.CreateTask("*/10 * * * * *", func() {
		fmt.Println("Executing Task 2 - every 10 seconds:", time.Now())
	}, "task2-lock")

	// View tasks
	tm.ViewTasks()

	// Update Task 1 with a new interval
	time.Sleep(15 * time.Second)
	tm.UpdateTask(taskID, "*/3 * * * * *", func() {
		fmt.Println("Executing Updated Task 1 - every 3 seconds:", time.Now())
	}, "task1-lock")

	// View updated tasks
	tm.ViewTasks()

	// Wait for a while and then delete Task 1
	time.Sleep(10 * time.Second)
	tm.DeleteTask(taskID)

	// View remaining tasks
	tm.ViewTasks()
	select {}
}
