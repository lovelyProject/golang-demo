package stat

import (
	"github.com/sirupsen/logrus"
	"go/adv-example/pkg/event"
)

type StatService struct {
	Repo     *StatRepository
	Eventbus *event.EventBus
}

func NewStatService(repo *StatRepository, bus *event.EventBus) *StatService {
	return &StatService{
		Repo:     repo,
		Eventbus: bus,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.Eventbus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				logrus.Fatalln("Bad event link visited data: ", msg.Data)
				continue
			}
			s.Repo.AddClick(id)
		}
	}
}
