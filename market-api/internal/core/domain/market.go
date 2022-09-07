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
	GetSubprefectureCode() string
	GetSubprefecture() string
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
	Create(market IMarket) (string, error)
	GetByID(id string) (IMarket, error)
	GetAll() ([]IMarket, error)
	Get(
		longitude string,
		latitude string,
		censusSector string,
		weightingArea string,
		township string,
		townshipCode string,
		subprefectureCode string,
		subprefecture string,
		region5 string,
		region8 string,
		name string,
		registry string,
		street string,
		number string,
		district string,
		reference string,
	) ([]IMarket, error)
	Update(id string, market IMarket) (IMarket, error)
	DeleteByRegistry(registry string) error
}

type IMarketPersistence interface {
	IMarketReaderPersistence
	IMarketWriterPersistence
}

type IMarketReaderPersistence interface {
	FindByID(id string) (IMarket, error)
	FindAll() ([]IMarket, error)
	Find(
		longitude string,
		latitude string,
		censusSector string,
		weightingArea string,
		township string,
		townshipCode string,
		subprefectureCode string,
		subprefecture string,
		region5 string,
		region8 string,
		name string,
		registry string,
		street string,
		number string,
		district string,
		reference string,
	) ([]IMarket, error)
}

type IMarketWriterPersistence interface {
	Save(market IMarket) (string, error)
	Update(id string, market IMarket) (IMarket, error)
	DeleteByRegistry(registry string) error
}

type Market struct {
	ID                string
	Longitude         string `valid:"required,maxstringlength(10)"`
	Latitude          string `valid:"required,maxstringlength(10)"`
	CensusSector      string `valid:"required,maxstringlength(15)"`
	WeightingArea     string `valid:"required,maxstringlength(13)"`
	TownshipCode      string `valid:"required,maxstringlength(9)"`
	Township          string `valid:"required,maxstringlength(18)"`
	SubprefectureCode string `valid:"required,maxstringlength(2)"`
	Subprefecture     string `valid:"required,maxstringlength(25)"`
	Region5           string `valid:"required,maxstringlength(6)"`
	Region8           string `valid:"required,maxstringlength(7)"`
	Name              string `valid:"required,maxstringlength(30)"`
	Registry          string `valid:"required,maxstringlength(6)"`
	Street            string `valid:"required,maxstringlength(34)"`
	Number            string `valid:"maxstringlength(15)"`
	District          string `valid:"maxstringlength(20)"`
	Reference         string `valid:"maxstringlength(30)"`
}

func (m *Market) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *Market) GetID() string {
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

func (m *Market) GetSubprefectureCode() string {
	return m.SubprefectureCode
}

func (m *Market) GetSubprefecture() string {
	return m.Subprefecture
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
