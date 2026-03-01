package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type InfringementRepository struct {
	db *gorm.DB
}

func NewInfringementRepository(db *gorm.DB) *InfringementRepository {
	return &InfringementRepository{db: db}
}

func (r *InfringementRepository) Create(infringement *schemas.Infringement) error {
	return r.db.Create(infringement).Error
}

func (r *InfringementRepository) Save(infringement *schemas.Infringement) error {
	return r.db.Save(infringement).Error
}

func (r *InfringementRepository) Delete(infringement *schemas.Infringement) error {
	return r.db.Delete(infringement).Error
}

func (r *InfringementRepository) ById(id int) (*schemas.Infringement, error) {
	var infringement schemas.Infringement
	err := r.db.Where("id = ?", id).First(&infringement).Error
	if err != nil {
		return nil, err
	}
	return &infringement, nil
}

func (r *InfringementRepository) ManyByUserId(userId int) ([]*schemas.Infringement, error) {
	var infringements []*schemas.Infringement
	err := r.db.Where("user_id = ?", userId).Order("time DESC").Find(&infringements).Error
	return infringements, err
}
