package domain

import (
	"github.com/asaskevich/govalidator"
)

type IMarket interface {
	IsValid() (bool, error)
	GetID() string
	GetLongitude() string
	GetLatitude() string
	GetCensusSector() string
	GetWeightingArea() string
	GetTownship() string
	GetTownshipCode() string
	GetSubPrefectureCode() string
	GetSubPrefecture() string
	GetRegion5() string
	GetRegion8() string
	GetName() string
	GetRegistry() string
	GetStreet() string
	GetNumber() string
	GetDistrict() string
	GetReference() string
}

type IMarketService interface {
	ICreateMarketService
	IFindMarketService
	IUpdateMarketService
	IDeleteMarketService
}

type ICreateMarketService interface {
	Create(market IMarket) (IMarket, error)
}

type IFindMarketService interface {
	GetByID(id int) (IMarket, error)
	GetAll() ([]IMarket, error)
	Get(township string, region5 string, name string, district string) ([]IMarket, error)
}

type IUpdateMarketService interface {
	Update(id int, market IMarket) (IMarket, error)
}

type IDeleteMarketService interface {
	DeleteByRegistry(registry string) error
}

type IMarketReaderPersistence interface {
	FindByID(id int) (IMarket, error)
	FindAll() ([]IMarket, error)
	Find(query map[string]string) ([]IMarket, error)
}

type IMarketWriterPersistence interface {
	Save(market IMarket) (IMarket, error)
	Update(id int, market IMarket) (IMarket, error)
	DeleteByRegistry(registry string) (int64, error)
}

type Market struct {
	ID                int
	Longitude         string `valid:"required"`
	Latitude          string `valid:"required"`
	CensusSector      string `valid:"required"`
	WeightingArea     string `valid:"required"`
	Township          string `valid:"required"`
	TownshipCode      string `valid:"required"`
	SubPrefectureCode string `valid:"required"`
	SubPrefecture     string `valid:"required"`
	Region5           string `valid:"required"`
	Region8           string `valid:"required"`
	Name              string `valid:"required"`
	Registry          string `valid:"required"`
	Street            string `valid:"required"`
	Number            string
	District          string `valid:"required"`
	Reference         string
}

func (m *Market) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *Market) GetID() int {
	return m.ID
}

func (m *Market) GetLongitude() string {
	return m.Longitude
}

func (m *Market) GetLatitude() string {
	return m.Latitude
}

func (m *Market) GetCensusSector() string {
	return m.CensusSector
}

func (m *Market) GetWeightingArea() string {
	return m.WeightingArea
}

func (m *Market) GetTownship() string {
	return m.Township
}

func (m *Market) GetTownshipCode() string {
	return m.TownshipCode
}

func (m *Market) GetSubPrefectureCode() string {
	return m.SubPrefectureCode
}

func (m *Market) GetSubPrefecture() string {
	return m.SubPrefecture
}

func (m *Market) GetRegion5() string {
	return m.Region5
}

func (m *Market) GetRegion8() string {
	return m.Region8
}

func (m *Market) GetName() string {
	return m.Name
}

func (m *Market) GetRegistry() string {
	return m.Registry
}

func (m *Market) GetStreet() string {
	return m.Street
}

func (m *Market) GetNumber() string {
	return m.Number
}

func (m *Market) GetDistrict() string {
	return m.District
}

func (m *Market) GetReference() string {
	return m.Reference
}

func NewMarket() *Market {
	return &Market{}
}
