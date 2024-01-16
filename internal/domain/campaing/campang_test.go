package campaing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "CampaingX"
	content  = "body"
	contacts = []string{"test@test.com", "test.2@test.com"}
)

func Test_NewCampaingSuccess(t *testing.T) {
	assert := assert.New(t)
	campaing := NewCampaing(name, content, contacts)
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}

func Test_NewCampaing_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	campaing := NewCampaing(name, content, contacts)
	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	campaing := NewCampaing(name, content, contacts)
	assert.Greater(campaing.CreatedOn, now)

}
