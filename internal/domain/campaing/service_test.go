package campaing

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaing = contract.NewCampaing{
		Name:    "Teste",
		Content: "Body Hi",
		Emails:  []string{"Teste@gmial.com"},
	}
	service = ServiceImp{}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaing *Campaing) error {
	args := r.Called(campaing)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaing, error) {
	return nil, nil
}

func (r *repositoryMock) GetBy(id string) (*Campaing, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaing), nil
}

func Test_Create_Campaig(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaing)
	assert.Nil(err)
	assert.NotNil(id)

}

func Test_Create_ValidateDomanError(t *testing.T) {
	assert := assert.New(t)
	newCampaing.Name = ""
	_, err := service.Create(newCampaing)
	assert.NotNil(err)
	assert.Equal("name is required with min 5", err.Error())

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

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetByiD_ReturnCampaing(t *testing.T) {
	assert := assert.New(t)
	campaing, _ := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaing.ID
	})).Return(campaing, nil)
	service.Repository = repositoryMock
	campaingReturned, _ := service.GetBy(campaing.ID)

	assert.Equal(campaing.ID, campaingReturned.ID)
}

func Test_GetByiD_ReturnErrorWhenSomenthingWrongExist(t *testing.T) {
	assert := assert.New(t)
	campaing, _ := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("Something wrong"))
	service.Repository = repositoryMock
	_, err := service.GetBy(campaing.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
