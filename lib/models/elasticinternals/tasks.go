package elasticinternals

// CatTaskItem; []CatTaskItem is returned by GET /_cat/tasks
// - StartTime is milli epoch
// - RunningTime can be: 3.2micros, 4.6d, 6.4h; i.e. <float><unit>
type CatTaskItem struct {
	Action       string `json:"action"`
	Ip           string `json:"ip"`
	Node         string `json:"node"`
	ParentTaskID string `json:"parent_task_id"`
	RunningTime  string `json:"running_time"`
	StartTime    int64  `json:"start_time"`
	TaskID       string `json:"task_id"`
	Timestamp    string `json:"timestamp"`
	Type         string `json:"type"`
}
