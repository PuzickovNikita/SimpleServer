package tables

type SimpleTable struct {
	Key  int    `json:"key"`
	Body string `json:"body,omitempty"`
}
