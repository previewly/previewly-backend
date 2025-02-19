package preview

import "wsw/backend/ent"

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
	StatusPending Status = "pending"
)

type (
	Status string

	Image struct {
		ID  *int
		URL string
	}

	PreviewData struct {
		ID     int
		URL    string
		Image  Image
		Status Status
		Error  *string
		Title  *string
		IsNew  bool
		Entity *ent.Url
	}
)
