package meetup

import "github.com/UpMeetApp/server/pkg/domain"

type meetupService struct {
	meetupRepository domain.MeetupRepository
}

// NewMeetupService creates a new meetup service instance.
func NewMeetupService(meetupRepository domain.MeetupRepository) domain.MeetupService {
	return &meetupService{
		meetupRepository: meetupRepository,
	}
}

func (s *meetupService) GetMeetupByID(uid string, id string) (*domain.Meetup, error) {
	//TODO implement me
	panic("implement me")
}

func (s *meetupService) CreateMeetup(uid string, dto *domain.CreateMeetupDTO) (*domain.Meetup, error) {
	//TODO implement me
	panic("implement me")
}

func (s *meetupService) UpdateMeetup(uid string, dto *domain.UpdateMeetupDTO) (*domain.Meetup, error) {
	//TODO implement me
	panic("implement me")
}

func (s *meetupService) DeleteMeetup(uid string, id string) error {
	//TODO implement me
	panic("implement me")
}
