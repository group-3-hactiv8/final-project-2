package comment_pg

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/comment_repository"
	"log"
	"fmt"

	"gorm.io/gorm"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{db: db}
}

func (c *commentPG) CreateComment(user *models.User, comment *models.Comment) (*models.Comment, errs.MessageErr) {
	
	comment.UserId = (user.ID)
	if err := c.db.Create(comment).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new comment")
	}

	return comment, nil
}

func (c *commentPG) GetCommentById(id uint) (*models.Comment, errs.MessageErr) {
	var comment models.Comment
	result := c.db.First(&comment, id)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound(fmt.Sprintf("failed to get comment by id :", comment.ID))
		return nil, error
	}
	return &comment, nil
}

func (c *commentPG) GetCommentByUserId(userId uint)  ([]models.Comment, errs.MessageErr) {
	var comment []models.Comment
	result := c.db.Find(&comment, "user_id=?", userId)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound("Failed To Get All Comment")
		return nil, error
	}
	return comment, nil
}

func (r *commentPG) GetCommentByPhotoId(photoId uint) ([]models.Comment, errs.MessageErr) {
	var comments []models.Comment
	err := r.db.Where("photo_id = ?", photoId).Find(&comments).Error
	if err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewNotFound("Failed To Get All Comment")
		return nil, error
	}
	return comments, nil
}

func (c *commentPG) GetAllComments() ([]models.Comment, errs.MessageErr) {
    var comments []models.Comment
    if err := c.db.Find(&comments).Error; err != nil {
        return nil, errs.NewInternalServerError(err.Error())
    }
    return comments, nil
}

func (c *commentPG) UpdateComment(comment *models.Comment, cmtUpdate *models.Comment) (*models.Comment, errs.MessageErr) {
	if err := c.db.Model(comment).Where("id = ?", comment.ID).Updates(cmtUpdate).Error; err != nil {
		if err != nil {
			log.Println("Error : ",err.Error())
			error := errs.NewNotFound("Failed To Updated Comment")
			return nil, error
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return comment, nil
}

func (c *commentPG) DeleteComment(comment *models.Comment) errs.MessageErr {
	result := c.db.Delete(comment)

	if err := result.Error; err != nil {
		log.Println("Error : ",err.Error())
		error := errs.NewInternalServerError(fmt.Sprintf("Failed to delete Comment by id : %v", comment.ID))
		return error
	}
	return  nil
}