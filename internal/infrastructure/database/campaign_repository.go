package database

import campaign "emailn/internal/domain/campaign"

type CampaignRepository struct {
	campaings []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaings = append(c.campaings, *campaign)
	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.campaings, nil
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	return nil, nil
}
