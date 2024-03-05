package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "CampaingX"
	content  = "body Hi!"
	contacts = []string{"test@test.com", "test.2@test.com"}
	fake     = faker.New()
)

func Test_NewCampaignSuccess(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_MustStatusStartWithPending(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(Pending, campaign.Status)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	campaign, _ := NewCampaign(name, content, contacts)
	assert.Greater(campaign.CreatedOn, now)

}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, contacts)
	assert.Equal("name is required with min 5", err.Error())

}

func Test_NewCampaigng_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)
	assert.Equal("name is required with max 24", err.Error())

}

func Test_NewCampaign_MustValidateNameContenteMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)
	assert.Equal("content is required with min 5", err.Error())

}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1050), contacts)
	assert.Equal("content is required with max 1024", err.Error())

}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)
	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContactsEmailInvalid(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"invalid_email"})
	assert.Equal("email is invalid", err.Error())
}
