package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	internal_mock "emailn/internal/test/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignPost_should_save_new_campaing(t *testing.T) {
	assert := assert.New(t)
	service := new(internal_mock.CampaignServiceMock)
	body := contract.NewCampaign{
		Name:    "Teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("34121341423423", nil)
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_CampaignPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	service := new(internal_mock.CampaignServiceMock)
	body := contract.NewCampaign{
		Name:    "Teste",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)

}
