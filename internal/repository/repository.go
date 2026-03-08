package repository

import "kt_project/internal/models"

type Storage interface {
	Save(stat models.Stat) error
}
