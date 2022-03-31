package meetup

import (
	"github.com/UpMeetApp/server/pkg/domain"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type meetupRepository struct {
	db *gorm.DB
}

// NewMeetupRepository creates a new meetup repository instance.
func NewMeetupRepository(db *gorm.DB) domain.MeetupRepository {
	return &meetupRepository{
		db: db,
	}
}

func (r *meetupRepository) CreateMeetup(m *domain.Meetup) error {
	err := r.db.Create(m).Error
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("failed to create meetup", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return nil
}

func (r *meetupRepository) GetMeetupByID(id string) (*domain.Meetup, error) {
	m := &domain.Meetup{}
	err := r.db.Where("id = ?", id).First(m).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.ErrNotFound
		}
		sentry.CaptureException(err)
		zap.L().Error("failed to get meetup by id", zap.Error(err))
		return nil, fiber.ErrInternalServerError
	}
	return m, nil
}

func (r *meetupRepository) UpdateMeetup(m *domain.Meetup) error {
	err := r.db.Save(m).Error
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("failed to update meetup", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return nil
}

func (r *meetupRepository) DeleteMeetup(id string) error {
	err := r.db.Delete(&domain.Meetup{}, "id = ?", id).Error
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("failed to delete meetup", zap.Error(err))
		return fiber.ErrInternalServerError
	}
	return nil
}
