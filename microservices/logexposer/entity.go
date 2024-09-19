package logexposer

type LogEntry struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Caller  string `json:"caller"`
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Data  []LogEntry `json:"data"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	Total int        `json:"total"`
}
