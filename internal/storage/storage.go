package storage

import "mailingAPI/internal/storage/models"

type Storage interface {
	AddClient(c *models.Client) error
	UpdateClient(c *models.Client) error
	DeleteClient(c *models.Client) error
	AddMailing(m *models.Mailing) error
	GetAllMailingStat() ([]models.Mailing, error)
	GetOneMailingStat(m *models.Mailing) (models.Mailing, error)
	UpdateMailing(m *models.Mailing) error
	DeleteMailing(m *models.Mailing) error
}
