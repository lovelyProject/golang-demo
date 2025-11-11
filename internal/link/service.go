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

func (linkService *LinkService) GetAll(limit, offset int) (GetAllLinksResponse, error) {
	links, err := linkService.Repo.GetAll(limit, offset)
	if err != nil {
		return GetAllLinksResponse{}, err
	}
	total, err := linkService.Repo.Count()
	if err != nil {
		return GetAllLinksResponse{}, err
	}
	return GetAllLinksResponse{
		Total: total,
		Links: links,
	}, nil
}
