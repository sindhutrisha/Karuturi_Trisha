package daos

import (
	"database/sql"
	"errors"
	"github.com/sindhutrisha/Karuturi_Trisha/karuturi_trisha/pkg/rest/server/daos/clients/sqls"
	"github.com/sindhutrisha/Karuturi_Trisha/karuturi_trisha/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type KaruturiTrishaDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateKaruturiTrishas(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS karuturiTrishas(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Sindhu TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewKaruturiTrishaDao() (*KaruturiTrishaDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateKaruturiTrishas(sqlClient)
	if err != nil {
		return nil, err
	}
	return &KaruturiTrishaDao{
		sqlClient,
	}, nil
}

func (karuturiTrishaDao *KaruturiTrishaDao) CreateKaruturiTrisha(m *models.KaruturiTrisha) (*models.KaruturiTrisha, error) {
	insertQuery := "INSERT INTO karuturiTrishas(Sindhu)values(?)"
	res, err := karuturiTrishaDao.sqlClient.DB.Exec(insertQuery, m.Sindhu)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("karuturiTrisha created")
	return m, nil
}

func (karuturiTrishaDao *KaruturiTrishaDao) UpdateKaruturiTrisha(id int64, m *models.KaruturiTrisha) (*models.KaruturiTrisha, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	karuturiTrisha, err := karuturiTrishaDao.GetKaruturiTrisha(id)
	if err != nil {
		return nil, err
	}
	if karuturiTrisha == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE karuturiTrishas SET Sindhu = ? WHERE Id = ?"
	res, err := karuturiTrishaDao.sqlClient.DB.Exec(updateQuery, m.Sindhu, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("karuturiTrisha updated")
	return m, nil
}

func (karuturiTrishaDao *KaruturiTrishaDao) DeleteKaruturiTrisha(id int64) error {
	deleteQuery := "DELETE FROM karuturiTrishas WHERE Id = ?"
	res, err := karuturiTrishaDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("karuturiTrisha deleted")
	return nil
}

func (karuturiTrishaDao *KaruturiTrishaDao) ListKaruturiTrishas() ([]*models.KaruturiTrisha, error) {
	selectQuery := "SELECT * FROM karuturiTrishas"
	rows, err := karuturiTrishaDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var karuturiTrishas []*models.KaruturiTrisha
	for rows.Next() {
		m := models.KaruturiTrisha{}
		if err = rows.Scan(&m.Id, &m.Sindhu); err != nil {
			return nil, err
		}
		karuturiTrishas = append(karuturiTrishas, &m)
	}
	if karuturiTrishas == nil {
		karuturiTrishas = []*models.KaruturiTrisha{}
	}

	log.Debugf("karuturiTrisha listed")
	return karuturiTrishas, nil
}

func (karuturiTrishaDao *KaruturiTrishaDao) GetKaruturiTrisha(id int64) (*models.KaruturiTrisha, error) {
	selectQuery := "SELECT * FROM karuturiTrishas WHERE Id = ?"
	row := karuturiTrishaDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.KaruturiTrisha{}
	if err := row.Scan(&m.Id, &m.Sindhu); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("karuturiTrisha retrieved")
	return &m, nil
}
