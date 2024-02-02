package endpoints

import "emailn/internal/domain/campaing"

type Handler struct {
	CampaingService campaing.Service
}
