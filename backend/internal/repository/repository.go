package repository

import "backend/internal/models"

type Storage interface {
	Save(stat models.Stat) error
	GetStat() ([]models.Stat, error)
}
