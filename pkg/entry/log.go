package entry

type LogEntry struct {
	ID  string `json:"id"`
	Msg string `json:"full_line"`
}
