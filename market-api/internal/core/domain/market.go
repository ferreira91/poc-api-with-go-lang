package domain

import (
	"github.com/asaskevich/govalidator"
)

type IMarket interface {
	IsValid() (bool, error)
}

type Market struct {
	ID                string
	Longitude         string `valid:"required"`
	Latitude          string `valid:"required"`
	CensusSector      string `valid:"required"`
	WeightingArea     string `valid:"required"`
	Township		  string `valid:"required"`
	TownshipCode      string `valid:"required"`
	SubPrefectureCode string `valid:"required"`
	SubPrefecture     string `valid:"required"`
	Region5           string `valid:"required"`
	Region8           string `valid:"required"`
	Name              string `valid:"required"`
	Registry          string `valid:"required"`
	Street            string `valid:"required"`
	Number            string
	District          string `json:"required"`
	Reference         string
}

func (m *Market) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
