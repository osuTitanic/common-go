package repositories

import (
	"github.com/osuTitanic/common-go/schemas"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) Create(report *schemas.Report) error {
	return r.db.Create(report).Error
}

func (r *ReportRepository) Save(report *schemas.Report) error {
	return r.db.Save(report).Error
}

func (r *ReportRepository) Delete(report *schemas.Report) error {
	return r.db.Delete(report).Error
}

func (r *ReportRepository) ById(id int) (*schemas.Report, error) {
	var report schemas.Report
	err := r.db.Where("id = ?", id).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *ReportRepository) ManyByTargetId(targetId int) ([]*schemas.Report, error) {
	var reports []*schemas.Report
	err := r.db.Where("target_id = ?", targetId).Order("time DESC").Find(&reports).Error
	return reports, err
}

func (r *ReportRepository) ManyBySenderId(senderId int) ([]*schemas.Report, error) {
	var reports []*schemas.Report
	err := r.db.Where("sender_id = ?", senderId).Order("time DESC").Find(&reports).Error
	return reports, err
}
