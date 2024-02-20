package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceyMock struct {
	mock.Mock
}

func (r *serviceyMock) Create(newCampaing contract.NewCampaing) (string, error) {
	args := r.Called(newCampaing)
	return args.String(0), args.Error(1)
}

func Test_CampaingPost_should_save_new_campaing(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceyMock)
	body := contract.NewCampaing{
		Name:    "Teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaing) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("34121341423423", nil)
	handler := Handler{CampaingService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, status, err := handler.CampaingPost(res, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_CampaingPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceyMock)
	body := contract.NewCampaing{
		Name:    "Teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaingService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaingPost(res, req)

	assert.NotNil(err)

}
