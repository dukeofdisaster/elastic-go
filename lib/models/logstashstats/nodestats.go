package logstashstats

// GetNodeStatsResponse is the response from GET /_node/stats
type GetNodeStatsResponse struct {
	EphemeralID string         `json:"ephemeral_id,omitempty"`
	Events      EventsObject   `json:"events,omitempty"`
	Host        string         `json:"host,omitempty"`
	HTTPAddress string         `json:"http_address,omitempty"`
	ID          string         `json:"id,omitempty"`
	Jvm         JvmObject      `json:"jvm,omitempty"`
	Name        string         `json:"name,omitempty"`
	OS          OsObject       `json:"os,omitempty"`
	Pipeline    PipelineObject `json:"pipeline,omitempty"`
	//Pipelines   []PipelinesItemObject `json:"pipelines,omitempty"`
	Pipelines []map[string]interface{} `json:"pipelines,omitempty"`
	Process   ProcessObject            `json:"process,omitempty"`
	Queue     QueueObject              `json:"queue,omitempty"`
	Reloads   ReloadsObject            `json:"reloads,omitempty"`
	Snapshot  bool                     `json:"snapshot,omitempty"`
	Status    string                   `json:"status,omitempty"`
	Version   string                   `json:"version,omitempty"`
}

// ReloadsObject contains count of reload success/fail
type ReloadsObject struct {
	Successes int64 `json:"successes,omitempty"`
	Failures  int64 `json:"failures,omitempty"`
}

// QueueObject contains events_count
type QueueObject struct {
	EventsCount int64 `json:"events_count,omitempty"`
}

// ProcessObject contains metadata about file desriptors, mem, and cpu
type ProcessObject struct {
	CPU                 ProcCPUObject `json:"cpu,omitempty"`
	MaxFileDescriptors  int64         `json:"max_file_descriptors,omitempty"`
	Mem                 ProcMemObject `json:"mem,omitempty"`
	OpenFileDescriptors int64         `json:"open_file_descriptors,omitempty"`
}

// ProcCPUObject contains info about total cpu and load average up to 15m
type ProcCPUObject struct {
	LoadAverage   LoadAverageObject `json:"load_average,omitempty"`
	Percent       int64             `json:"percent,omitempty"`
	TotalInMillis int64             `json:"total_in_millis,omitempty"`
}

// LoadAverageObject contains metrics for cpu load over 1m,5m,15m
type LoadAverageObject struct {
	OneMin     float64 `json:"1m,omitempty"`
	FiveMin    float64 `json:"5m,omitempty"`
	FifteenMin float64 `json:15m,omitempty"`
}

// ProcMemObject is part of a process object that contains info about virt mem
type ProcMemObject struct {
	TotalVirtualInBytes int64 `json:"total_virtual_in_bytes,omitempty"`
}

// PipelinesItemObject are part of an array that can have arbitrary names
// so probably best to hadle unmarshalling in progras. PipelinesItemObject
// array is just an array of map[string]interface{}
//type PipelinesItemObject struct {
//}

// OsObject contains CGroupObject which has cpu stats
type OsObject struct {
	CGroup CGroupObject `json:"cgroup,omitempty"`
}

// CGroupObject contains meta about cpu stats
type CGroupObject struct {
	CPUAcct CPUAcctObject `json:"cpuacct,omitempty"`
	CPU     CPUObject     `json:"cpu,omitempty"`
}

// CPUAcctObject contains control group and usage metrics
type CPUAcctObject struct {
	ControlGroup string `json:"control_group,omitempty"`
	UsageNanos   int64  `json:"usage_nanos,omitempty"`
}

// CPUObject contains similar info to CPUAcctObject
type CPUObject struct {
	ControlGroup    string        `json:"control_group,omitempty"`
	CfsQuotaMicros  int64         `json:"cfs_quota_micros,omitempty"`
	Stat            CPUStatObject `json:"stat,omitempty"`
	CfsPeriodMicros int64         `json:"cfs_period_micros,omitempty"`
}

// CPUStatObject contains 3 metrics about stat
type CPUStatObject struct {
	NumberOfElapsedPeriods int64 `json:"number_of_elapsed_periods,omitempty"`
	NumberOfTimesThrottled int64 `json:"number_of_times_throttled,omitempty"`
	TimeThrottledNanos     int64 `json:"time_throttled_nanos,omitempty"`
}

// EventsObject is the top level events struct in the response
type EventsObject struct {
	DurationInMillis          int64 `json:"duration_in_millis,omitempty"`
	Filtered                  int64 `json:"filtered,omitempty"`
	In                        int64 `json:"in,omitempty"`
	Out                       int64 `json:"out,omitempty"`
	QueuePushDurationInMillis int64 `json:"queue_push_duration_in_millis,omitempty"`
}

// PipelineObject contains metadata about workers and batches
type PipelineObject struct {
	Workers    int64 `json:"workers,omitempty"`
	BatchSize  int64 `json:"batch_size,omitempty"`
	BatchDelay int64 `json:"batch_delay,omitempty"`
}

// JvmObject contains meta about the underlying java jvm running the logstash node
type JvmObject struct {
	Threads ThreadsObject `json:"threads,omitempty"`
	Mem     MemObject     `json:"mem,omitempty"`
	GC      GCObject
}

// ThreadsObject contains meta about the number of threads in the JVM
type ThreadsObject struct {
	Count     int64 `json:"count,omitempty"`
	PeakCount int64 `json:peak_count,omitempty"`
}

// MemObject conains meta about the JVM heap
type MemObject struct {
	HeapUsedPercent         int64       `json:"heap_used_percent,omitempty"`
	HeapCommittedInBytes    int64       `json:"heap_committed_in_bytes,omitempty"`
	HeapMaxInBytes          int64       `json:"heap_max_in_bytes,omitempty"`
	HeapUsedInBytes         int64       `json:"heap_used_in_bytes,omitempty"`
	NonHeapUsedInBytes      int64       `json:"non_heap_used_in_bytes,omitempty"`
	NonHeapCommittedInBytes int64       `json:"non_heap_committed_in_bytes,omitempty"`
	Pools                   PoolsObject `json:"pools,omitempty"`
}

// PoolsObject contains meta about jvm pools by type survivor,young,old; measured in bytes
type PoolsObject struct {
	Survivor       SurvivorObject `json:"survivor_object,omitempty"`
	Young          YoungObject
	Old            OldObject
	UptimeInMillis int64 `json:"uptime_in_millis,omitempty"`
}

// SurvivorObject contains bytes info relevant to survivors
type SurvivorObject struct {
	CommittedInBytes int64 `json:"comitted_in_bytes,omitempty"`
	MaxInBytes       int64 `json:"max_in_bytes,omitempty"`
	PeakMaxInBytes   int64 `json:"peak_max_in_bytes,omitempty"`
	PeakUsedInBytes  int64 `json:"peak_used_in_bytes,omitempty"`
	UsedInBytes      int64 `json:"used_in_bytes,omitempty"`
}

// YoungObject contains bytes info relevent to young gc
type YoungObject struct {
	CommittedInBytes int64 `json:"comitted_in_bytes,omitempty"`
	MaxInBytes       int64 `json:"max_in_bytes,omitempty"`
	PeakMaxInBytes   int64 `json:"peak_max_in_bytes,omitempty"`
	PeakUsedInBytes  int64 `json:"peak_used_in_bytes,omitempty"`
	UsedInBytes      int64 `json:"used_in_bytes,omitempty"`
}

// OldObject contains bytes info relevant to old gc
type OldObject struct {
	CommittedInBytes int64 `json:"comitted_in_bytes,omitempty"`
	MaxInBytes       int64 `json:"max_in_bytes,omitempty"`
	PeakMaxInBytes   int64 `json:"peak_max_in_bytes,omitempty"`
	PeakUsedInBytes  int64 `json:"peak_used_in_bytes,omitempty"`
	UsedInBytes      int64 `json:"used_in_bytes,omitempty"`
}

// GCObject is part of JvmObject and contains children of type CollectorsObject
type GCObject struct {
	Collectors CollectorsObject `json:"collectors,omitempty"`
}

// CollectorsObject contains meta about old vs young GC CollectionStats
type CollectorsObject struct {
	Young CollectionStats `json:"young,omitempty"`
	Old   CollectionStats `json:"old,omitempty"`
}

// CollectionStats is the atomic data field of GCObject containing collection time and count
type CollectionStats struct {
	CollectionTimeInMillis int64 `json:"collection_time_in_millis,omitempty"`
	CollectionCount        int64 `json:"collection_count,omitempty"`
}
