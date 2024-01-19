package campaing

import (
	"emailn/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaing = contract.NewCampaing{
		Name:    "Teste",
		Content: "Body",
		Emails:  []string{"Teste@gmial.com"},
	}
	service = Service{}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaing *Campaing) error {
	args := r.Called(campaing)
	return args.Error(0)
}

func Test_Create_Campaig(t *testing.T) {
	assert := assert.New(t)
	id, err := service.Create(newCampaing)
	assert.Nil(err)
	assert.NotNil(id)

}

func Test_Create_ValidateDomanError(t *testing.T) {
	assert := assert.New(t)
	newCampaing.Name = ""
	_, err := service.Create(newCampaing)
	assert.NotNil(err)
	assert.Equal("name is required", err.Error())

}

func Test_Create_SaveCampaing(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaing *Campaing) bool {
		if campaing.Name != newCampaing.Name {
			return false
		} else if campaing.Content != newCampaing.Content {
			return false
		} else if len(campaing.Contacts) != len(newCampaing.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = repositoryMock
	service.Create(newCampaing)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock
	_, err := service.Create(newCampaing)

	assert.Equal("error to save on database", err.Error())
}
