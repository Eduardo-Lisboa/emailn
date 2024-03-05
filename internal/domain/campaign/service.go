package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internalErrors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaing, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaing)
	if err != nil {
		return "", internalerrors.ErrInternal
	}
	return campaing.ID, nil

}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {

	campaing, err := s.Repository.GetBy(id)
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.CampaignResponse{
		ID:      campaing.ID,
		Name:    campaing.Name,
		Content: campaing.Content,
		Status:  campaing.Status,
	}, nil
}
