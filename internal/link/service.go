package link

type LinkService struct {
	Repo *LinkRepository
}

func NewService(repo *LinkRepository) *LinkService {
	return &LinkService{Repo: repo}
}

func (linkService *LinkService) CreateUniqueHash(url string) (*Link, error) {
	link := NewLink(url)
	for {
		existedLink, _ := linkService.Repo.GetByHash(url)
		if existedLink == nil {
			break
		}
		link = NewLink(url)
	}
	createdLink, err := linkService.Repo.Create(link)
	return createdLink, err
}

func (linkService *LinkService) FindById(id uint) (*Link, error) {
	return linkService.Repo.GetById(id)
}

func (linkService *LinkService) Delete(id uint) error {
	err := linkService.Repo.Delete(id)
	return err
}
