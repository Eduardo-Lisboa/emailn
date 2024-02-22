package database

import "emailn/internal/domain/campaing"

type CampaingRepository struct {
	campaings []campaing.Campaing
}

func (c *CampaingRepository) Save(campaing *campaing.Campaing) error {
	c.campaings = append(c.campaings, *campaing)
	return nil
}

func (c *CampaingRepository) Get() ([]campaing.Campaing, error) {
	return c.campaings, nil
}

func (c *CampaingRepository) GetBy(id string) (*campaing.Campaing, error) {
	return nil, nil
}
