package helper

import (
	"encoding/json"
	"os"
	"sync"
)

// id: A unique identifier for the task
// description: A short description of the task
// status: The status of the task (todo, in-progress, done)
// createdAt: The date and time when the task was created
// updatedAt: The date and time when the task was last updated

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// 定义数据结构
type Config struct {
	Data []Task `json:"data"`
}

var (
	configMutex sync.Mutex // 如果有并发读写，必须加锁
)

var path = "data.json"

func LoadFromFile() (*Config, error) {
	// 1. 读取文件字节流
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 2. 解析 JSON 到结构体
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func SyncToFile(data interface{}) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	// 1. 将数据转为带缩进的 JSON，方便人类阅读
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// 2. 打开文件：不存在则创建，存在则清空（O_TRUNC）
	// 0644 代表文件所有者可读写，其他人只读
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
