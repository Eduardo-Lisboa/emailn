package endpoints

import (
	"emailn/internal/contract"
	internal_mock "emailn/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaing := contract.CampaignResponse{
		ID:      "1",
		Name:    "Teste",
		Content: "hi !",
		Status:  "Pending",
	}

	service := new(internal_mock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaing, nil)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(res, req)

	assert.Equal(200, status)
	assert.Equal(campaing.ID, response.(*contract.CampaignResponse).ID)

}

func Test_CampaignGetById_should_return_error(t *testing.T) {
	assert := assert.New(t)
	errExpected := errors.New("error")
	service := new(internal_mock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignGetById(res, req)

	assert.Equal(errExpected.Error(), err.Error())

}
