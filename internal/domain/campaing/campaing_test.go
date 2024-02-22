package campaing

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

func Test_NewCampaingSuccess(t *testing.T) {
	assert := assert.New(t)
	campaing, _ := NewCampaing(name, content, contacts)
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}

func Test_NewCampaing_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	campaing, _ := NewCampaing(name, content, contacts)
	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_MustStatusStartWithPending(t *testing.T) {
	assert := assert.New(t)
	campaing, _ := NewCampaing(name, content, contacts)

	assert.Equal(Pending, campaing.Status)
}

func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	campaing, _ := NewCampaing(name, content, contacts)
	assert.Greater(campaing.CreatedOn, now)

}

func Test_NewCampaing_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaing("", content, contacts)
	assert.Equal("name is required with min 5", err.Error())

}

func Test_NewCampaing_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(fake.Lorem().Text(30), content, contacts)
	assert.Equal("name is required with max 24", err.Error())

}

func Test_NewCampaing_MustValidateNameContenteMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", contacts)
	assert.Equal("content is required with min 5", err.Error())

}

func Test_NewCampaing_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, fake.Lorem().Text(1050), contacts)
	assert.Equal("content is required with max 1024", err.Error())

}

func Test_NewCampaing_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, nil)
	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaing_MustValidateContactsEmailInvalid(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{"invalid_email"})
	assert.Equal("email is invalid", err.Error())
}
