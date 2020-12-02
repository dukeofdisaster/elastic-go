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
