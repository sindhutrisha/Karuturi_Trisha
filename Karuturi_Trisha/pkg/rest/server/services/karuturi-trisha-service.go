package services

import (
	"github.com/sindhutrisha/Karuturi_Trisha/karuturi_trisha/pkg/rest/server/daos"
	"github.com/sindhutrisha/Karuturi_Trisha/karuturi_trisha/pkg/rest/server/models"
)

type KaruturiTrishaService struct {
	karuturiTrishaDao *daos.KaruturiTrishaDao
}

func NewKaruturiTrishaService() (*KaruturiTrishaService, error) {
	karuturiTrishaDao, err := daos.NewKaruturiTrishaDao()
	if err != nil {
		return nil, err
	}
	return &KaruturiTrishaService{
		karuturiTrishaDao: karuturiTrishaDao,
	}, nil
}

func (karuturiTrishaService *KaruturiTrishaService) CreateKaruturiTrisha(karuturiTrisha *models.KaruturiTrisha) (*models.KaruturiTrisha, error) {
	return karuturiTrishaService.karuturiTrishaDao.CreateKaruturiTrisha(karuturiTrisha)
}

func (karuturiTrishaService *KaruturiTrishaService) UpdateKaruturiTrisha(id int64, karuturiTrisha *models.KaruturiTrisha) (*models.KaruturiTrisha, error) {
	return karuturiTrishaService.karuturiTrishaDao.UpdateKaruturiTrisha(id, karuturiTrisha)
}

func (karuturiTrishaService *KaruturiTrishaService) DeleteKaruturiTrisha(id int64) error {
	return karuturiTrishaService.karuturiTrishaDao.DeleteKaruturiTrisha(id)
}

func (karuturiTrishaService *KaruturiTrishaService) ListKaruturiTrishas() ([]*models.KaruturiTrisha, error) {
	return karuturiTrishaService.karuturiTrishaDao.ListKaruturiTrishas()
}

func (karuturiTrishaService *KaruturiTrishaService) GetKaruturiTrisha(id int64) (*models.KaruturiTrisha, error) {
	return karuturiTrishaService.karuturiTrishaDao.GetKaruturiTrisha(id)
}
