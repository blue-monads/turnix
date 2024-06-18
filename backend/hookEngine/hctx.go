package hookengine

type EventContext struct {
	ProjectId   int64             `json:"project_id"`
	EventId     int64             `json:"event_id"`
	RunasUserID int64             `json:"run_as_user"`
	ExtraMeta   map[string]string `json:"envs"`
}
