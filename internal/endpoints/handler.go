package endpoints

import campaing "emailn/internal/domain/campaign"

type Handler struct {
	CampaignService campaing.Service
}
