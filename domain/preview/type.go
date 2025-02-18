package preview

import "wsw/backend/ent"

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
	StatusPending Status = "pending"
)

type PreviewData struct {
	ID     int
	URL    string
	Image  string
	Status Status
	Error  *string
	Title  *string
	IsNew  bool
	Entity *ent.Url
}
