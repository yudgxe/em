package user

import (
	"context"
	"em/internal/model"
	"em/pkg/utils"

	"github.com/gofiber/fiber/v2/log"
)

type UserStorage interface {
	SaveUser(context.Context, *model.User) error
}

type Enricher interface {
	GetAgeByName(ctx context.Context, name string) (*int32, error)
	GetGenderByName(ctx context.Context, name string) (*string, error)
	GetNationalityByName(ctx context.Context, name string) (*string, error)
}

type UserService struct {
	userStorage UserStorage
	enricher    Enricher
}

func New(userStorage UserStorage, enricher Enricher) *UserService {
	return &UserService{
		userStorage: userStorage,
		enricher:    enricher,
	}
}

func (this *UserService) EnrichAndSaveUser(ctx context.Context, user *model.User) error {
	const op string = "services.user.EnrichAndSaveUser"

	age, err := this.enricher.GetAgeByName(ctx, user.Name)
	if err != nil {
		log.Errorf("%s: enricher.GetAgeByName: %w", op, err)
	}

	gender, err := this.enricher.GetGenderByName(ctx, user.Name)
	if err != nil {
		log.Errorf("%s: enricher.GetGenderByName: %w", op, err)
	}

	nationality, err := this.enricher.GetNationalityByName(ctx, user.Name)
	if err != nil {
		log.Errorf("%s: enricher.GetNationalityByName: %w", op, err)
	}

	user.Age = age
	user.Gender = model.NewGender(gender)
	user.Nationality = nationality

	return utils.WrapError(this.userStorage.SaveUser(ctx, user), op)
}
