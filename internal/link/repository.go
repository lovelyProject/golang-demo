package link

import (
	"go/adv-example/db"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)

	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) Count() (int, error) {
	var count int64
	result := repo.Database.DB.
		Table("links").
		Where("deleted_at IS NULL").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func (repo *LinkRepository) GetAll(limit, offset int) ([]Link, error) {
	var links []Link
	result := repo.Database.DB.
		Table("links").
		Where("deleted_at IS NULL").
		Order("links.id ASC").
		Offset(offset).
		Limit(limit).
		Scan(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}
