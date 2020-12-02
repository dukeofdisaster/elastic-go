package elastic

// start at the top and decompose
type SuricataFlowSuccess struct {
	Took     int64      `json:"took"`
	TimedOut bool       `json:"timed_out"`
	Shards   ShardsMeta `json:"_shards"`
	Hits     HitObject  `json:"hits"`
}
type ShardsMeta struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}
type HitObject struct {
	Total    TotalObject         `json:total"`
	MaxScore float64             `json:"max_score"`
	Hits     []SuricataHitObject `json:"hits"`
}
type TotalObject struct {
	Value    int64  `json:"value"`
	Relation string `json:relation"`
}
type SuricataHitObject struct {
	Index          string               `json:"_index"`
	Type           string               `json:"_type"`
	Id             string               `json:"_id"`
	Score          float64              `json:"_score"`
	SourceInternal SuricataSourceObject `json:"_source"`
}
type SuricataSourceObject struct {
	Input       InputObject       `json:"input"`
	Timestamp   string            `json:"@timestamp"`
	Agent       AgentObject       `json:"agent"`
	Tags        []string          `json:"tags"`
	Ecs         EcsObject         `json:"ecs"`
	Version     string            `json:"@version"`
	Log         LogObject         `json:"log"`
	Source      SourceObject      `json:"source"`
	Destination DestinationObject `json:"destination"`
	Suricata    SuricataObject    `json:"suricata"`
}
type InputObject struct {
	Type string `json:"type"`
}
type SuricataObject struct {
	SrcPort   int64       `json:"src_port"`
	InIface   string      `json:"in_iface"`
	Proto     string      `json:"proto"`
	EventType string      `json:"event_type"`
	Timestamp string      `json:"timestamp"`
	DestIp    string      `json:"dest_ip"`
	FlowId    int64       `json:"flow_id"`
	SrcIp     string      `json:"src_ip"`
	DestPort  int64       `json:"dest_port"`
	AppProto  string      `json:"app_proto"`
	Agent     AgentObject `json:"agent"`
	Flow      FlowObject  `json:"flow"`
}
type DestinationObject struct {
	// By default if type isn't a pointer, json package
	// will unmarshal missing value to 0 for numbers and "" for strings
	Port int64  `json:"port"`
	Ip   string `json:"ip"`
}
type SourceObject struct {
	Port int64  `json:"port"`
	Ip   string `json:"ip"`
}
type LogObject struct {
	Offset int64      `json:"offset"`
	File   FileObject `json:"file"`
}
type FileObject struct {
	Path string `json:"path"`
}
type EcsObject struct {
	Version string `json:"version"`
}
type AgentObject struct {
	Version     string `json:"version"`
	Hostname    string `json:"hostname"`
	Id          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	EphemeralId string `json:"ephemeral_id"`
}
type FlowObject struct {
	Reason        string `json:"reason"`
	Start         string `json:"start"`
	BytesToServer int64  `json:"bytes_toserver"`
	Alerted       bool   `json:"alerted"`
	End           string `json:"end"`
	PktsToserver  int64  `json:"pkts_toserver"`
	PktsToclient  int64  `json:"pkts_toclient"`
	Age           int64  `json:"age"`
	BytesToclient int64  `json:"bytes_toclient"`
	State         string `json:"state"`
}
