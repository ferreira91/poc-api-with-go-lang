package domain

import (
	"github.com/asaskevich/govalidator"
)

type IMarket interface {
	IsValid() (bool, error)
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
	GetByID(id string) (IMarket, error)
	GetAll() ([]IMarket, error)
	Get(township string, region5 string, name string, district string) ([]IMarket, error)
}

type IUpdateMarketService interface {
	Update(id string, market IMarket) (IMarket, error)
}

type IDeleteMarketService interface {
	DeleteByRegistry(registry string) error
}

type IMarketReaderPersistence interface {
	FindByID(id string) (IMarket, error)
	FindAll() ([]IMarket, error)
	Find(query map[string]string) ([]IMarket, error)
}

type IMarketWriterPersistence interface {
	Save(market IMarket) (IMarket, error)
	Update(id string, market IMarket) (IMarket, error)
	DeleteByRegistry(registry string) (int64, error)
}

type Market struct {
	ID                string
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

func NewMarket() *Market {
	return &Market{}
}
