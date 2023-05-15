package services

import (
	"final-project-2/dto"
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/social_media_repository"
)

type SocialMediaService interface {
	CreateSocialMedia(payload *dto.NewSocialMediaRequest, userId uint) (*dto.NewSocialMediaResponse, errs.MessageErr)
	GetAllSocialMedias() (*dto.AllSocialMediasResponse, errs.MessageErr)
	UpdateSocialMedia(payload *dto.NewSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(id uint, sm_id uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr)
}

type socialMediaService struct {
	socialMediaRepo social_media_repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepo social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepo: socialMediaRepo}
}

func (sm *socialMediaService) CreateSocialMedia(payload *dto.NewSocialMediaRequest, userId uint) (*dto.NewSocialMediaResponse, errs.MessageErr) {
	newSM := payload.NewSocialMediaRequestToModel()
	newSM.UserId = userId

	createdSM, err := sm.socialMediaRepo.CreateSocialMedia(newSM)
	if err != nil {
		return nil, err
	}

	response := &dto.NewSocialMediaResponse{
		ID:             createdSM.ID,
		Name:           createdSM.Name,
		SocialMediaUrl: createdSM.SocialMediaUrl,
		UserId:         createdSM.UserId,
		CreatedAt:      createdSM.CreatedAt,
	}

	return response, nil
}

func (sm *socialMediaService) GetAllSocialMedias() (*dto.AllSocialMediasResponse, errs.MessageErr) {
	allSM, totalCount, err := sm.socialMediaRepo.GetAllSocialMedias()
	if err != nil {
		return nil, err
	}

	var smListResponse []dto.SocialMediaResponse = make([]dto.SocialMediaResponse, 0, totalCount)

	var userOfSM models.User
	var userOfSMResponse dto.UserOfSocialMediaResponse
	var smResponse dto.SocialMediaResponse

	for _, sm := range *allSM {
		userOfSM = sm.User
		userOfSMResponse = dto.UserOfSocialMediaResponse{
			ID:              userOfSM.ID,
			Username:        userOfSM.Username,
			ProfileImageUrl: "diisi apa?", // TODO: diisi dengan photo_url apa?
		}
		smResponse = dto.SocialMediaResponse{
			ID:             sm.ID,
			Name:           sm.Name,
			SocialMediaUrl: sm.SocialMediaUrl,
			UserId:         sm.UserId,
			CreatedAt:      sm.CreatedAt,
			UpdatedAt:      sm.UpdatedAt,
			User:           userOfSMResponse,
		}
		smListResponse = append(smListResponse, smResponse)
	}

	response := &dto.AllSocialMediasResponse{
		SocialMedias: smListResponse,
	}

	return response, nil
}

func (sm *socialMediaService) UpdateSocialMedia(payload *dto.NewSocialMediaRequest) (*dto.UpdateSocialMediaResponse, errs.MessageErr) {
	return nil, nil
}

func (sm *socialMediaService) DeleteSocialMedia(id uint, sm_id uint) (*dto.DeleteSocialMediaResponse, errs.MessageErr) {
	return nil, nil
}
