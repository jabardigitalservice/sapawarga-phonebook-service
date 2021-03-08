package repository

import (
	"context"

	"github.com/sapawarga/phonebook-service/model"
)

// PhoneBookI ...
type PhoneBookI interface {
	GetListPhoneBook(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error)
	GetMetaDataPhoneBook(ctx context.Context, params *model.GetListRequest) (int64, error)
	GetPhonebook(ctx context.Context, params interface{}) (interface{}, error)
	Insert(ctx context.Context, params interface{}) error
}
