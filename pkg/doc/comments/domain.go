package comments

import (
	"backend-comee/pkg/entity/user_information_personal"
	"time"

	"github.com/asaskevich/govalidator"
)

// Comment  Model struct Comment
type Comment struct {
	ID        int                                                `json:"id" db:"id" valid:"-"`
	UserId    string                                             `json:"user_id" db:"user_id" valid:"required"`
	EntityId  int                                                `json:"entity_id" db:"entity_id" valid:"required"`
	Value     string                                             `json:"value" db:"value" valid:"required"`
	Start     int                                                `json:"start" db:"start" valid:"-"`
	IsDelete  int                                                `json:"is_delete" db:"is_delete" valid:"-"`
	User      *user_information_personal.UserInformationPersonal `json:"user"`
	CreatedAt time.Time                                          `json:"created_at" db:"created_at"`
	UpdatedAt time.Time                                          `json:"updated_at" db:"updated_at"`
}

func NewComment(id int, userId string, entityId int, value string, start int, isDelete int) *Comment {
	return &Comment{
		ID:       id,
		UserId:   userId,
		EntityId: entityId,
		Value:    value,
		Start:    start,
		IsDelete: isDelete,
	}
}

func NewCreateComment(userId string, entityId int, value string, start int, isDelete int) *Comment {
	return &Comment{
		UserId:   userId,
		EntityId: entityId,
		Value:    value,
		Start:    start,
		IsDelete: isDelete,
	}
}

func (m *Comment) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
