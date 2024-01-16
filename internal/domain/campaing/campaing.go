package campaing

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaing struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaing(name string, content string, emails []string) *Campaing {

	contacts := make([]Contact, len(emails))
	for i, v := range emails {
		contacts[i].Email = v
	}

	return &Campaing{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
