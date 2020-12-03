package elasticinternals

// The Json object returned is a slice of indices objects
type CatIndicesJsonResponse struct {
	Indices []CatIndex
}

type CatIndex struct {
	// From the api these are all string; can't unmarshal string into int,
	// so have to try parse int
	//DocsCount   int64  `json:"docs.count"`
	DocsCount string `json:"docs.count"`
	//DocsDeleted int64  `json:"docs.deleted"`
	DocsDeleted string `json:"docs.deleted"`
	Health      string `json:"health"`
	Index       string `json:"index"`
	//Pri          int64  `json:"pri"`
	Pri          string `json:"pri"`
	PriStoreSize string `json:"pri.store.size"`
	Rep          string `json:"rep"`
	Status       string `json:"status"`
	StoreSize    string `json:"store.size"`
	Uuid         string `json:"uuid"`
}

// response from /_tasks/<TASKID>
type GetTaskIdResponse struct {
	Completed bool       `json:"completed"`
	Task      TaskObject `json:"task"`
	// when a task is completed, it will have response objet with same/smilar meta as task
	Response ResponseObject `json:"response"`
}
type ResponseObject struct {
	Took                 int64         `json:"took"`
	TimedOut             int64         `json:"timed_out"`
	Total                int64         `json:"total"`
	Updated              int64         `json:"updated"`
	Created              int64         `json:"created"`
	Deleted              int64         `json:"deleted"`
	Batches              int64         `json:"batches"`
	VersionConflicts     int64         `json:"version_conflicts"`
	Noops                int64         `json:"nops"`
	Retries              RetriesObject `json:"retries"`
	Throttled            string        `json:"throttled"`
	ThrottledMillis      int64         `json:"throttled_millis"`
	RequestsPerSecond    float64       `json:"requests_per_second"`
	ThrottledUntil       string        `json:"thorttled_until"`
	ThrottledUntilMillis int64         `json:"throttled_until_milis"`
	Failures             []string      `json:"failures"`
}

type TaskObject struct {
	Node             string `json:"node"`
	Id               int64  `json:"id"`
	Type             string `json:"type"`
	Action           string `json:"action"`
	Status           StatusObject
	Description      string        `json:"description"`
	StartTimeMillis  int64         `json:"start_time_in_millis"`
	RunningTimeNanos int64         `json:"running_time_in_nanos"`
	Cancellable      bool          `json:"cancellable"`
	Headers          HeadersObject `json:"headers"`
}
type StatusObject struct {
	Total                int64 `json:"total"`
	Updated              int64 `json:"updated"`
	Created              int64 `json:"created"`
	Deleted              int64 `json:"deleted"`
	Batches              int64 `json:"batches"`
	VersionConflicts     int64 `json:"version_conflicts"`
	Noops                int64 `json:"noops"`
	Retries              RetriesObject
	ThrottledMillis      int64   `json:"throttled_millis"`
	RequestsPerSecond    float64 `json:"requests_per_second"`
	ThrottledUntilMillis int64   `json:"throttled_until_milis"`
}
type RetriesObject struct {
	Bulk   int64 `json:"bulk"`
	Search int64 `json:"search"`
}
type HeadersObject struct {
}
