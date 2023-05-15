package social_media_pg

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/social_media_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type socialMediaPG struct {
	db *gorm.DB
}

func NewSocialMediaPG(db *gorm.DB) social_media_repository.SocialMediaRepository {
	return &socialMediaPG{db: db}
}

func (sm *socialMediaPG) CreateSocialMedia(newSm *models.SocialMedia) (*models.SocialMedia, errs.MessageErr) {
	// Ga perlu check if user exist.
	// Karna User nya pasti ada, kan ngambil userId nya dari token.

	if err := sm.db.Create(newSm).Error; err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed to create a new social media with name %s", newSm.Name)
		error := errs.NewInternalServerError(message)
		return nil, error
	}

	return newSm, nil
}

func (sm *socialMediaPG) GetAllSocialMedias() (*[]models.SocialMedia, uint, errs.MessageErr) {
	var allSM *[]models.SocialMedia
	result := sm.db.Find(&allSM)

	if err := result.Error; err != nil {
		log.Println(err.Error())
		error := errs.NewInternalServerError("Something is error when fetching all Social Media datas")
		return nil, 0, error
	}

	totalCount := result.RowsAffected

	// create new slice for storing the user of social media too
	var newAllSM []models.SocialMedia = make([]models.SocialMedia, 0, totalCount)
	var user *models.User

	for _, smObject := range *allSM {
		user = &models.User{}

		err := sm.db.Where("id = ?", smObject.UserId).Take(&user).Error
		if err != nil {
			error := errs.NewInternalServerError("Data Inconsistency: User of Social Media not found")
			return nil, 0, error
		}
		smObject.User = *user                 // kalau gini doang, smObject nya ga bener2 ke update, jd pas di akses di service, nilai user tetep nil
		newAllSM = append(newAllSM, smObject) // agar objek User nya beneran ke update
	}

	return &newAllSM, uint(totalCount), nil
}

func (sm *socialMediaPG) UpdateSocialMedia(updatedSm *models.SocialMedia) (*models.SocialMedia, errs.MessageErr) {
	initialSocialMedia := &models.SocialMedia{}
	err := sm.db.Where("id = ?", updatedSm.ID).Take(&initialSocialMedia).Error

	if err != nil {
		message := fmt.Sprintf("Social Media with ID %v not found", updatedSm.ID)
		err2 := errs.NewNotFound(message)
		return nil, err2
	}

	err = sm.db.Model(updatedSm).Updates(updatedSm).Error

	if err != nil {
		err2 := errs.NewBadRequest(err.Error())
		return nil, err2
	}

	updatedSm.UserId = initialSocialMedia.UserId

	return updatedSm, nil
}

func (sm *socialMediaPG) DeleteSocialMedia(sm_id uint) errs.MessageErr {
	err := sm.db.Delete(&models.SocialMedia{ID: sm_id}).Error

	if err != nil {
		err3 := errs.NewInternalServerError(err.Error())
		return err3
	}

	return nil
}
