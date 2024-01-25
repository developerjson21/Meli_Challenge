package repository

import (
	"database/sql"
	"fmt"
	"log"
	"meli_golang_gin_basic_app/cmd/internal/domain"
	"meli_golang_gin_basic_app/cmd/internal/dto"
	"time"
)
type IStorageDB interface {
	FindByJobId(jobId string) ([]domain.Entity, error)
	Create(entity dto.RepositoryDTO) (string, error)
}

type StorageDB struct {
	DB *sql.DB
}

func NewStorageDB(DB *sql.DB ) *StorageDB {
	return &StorageDB{
		DB: DB,
	}
}
// ************** FUNCION PARA INSERTAR REGISTRO DEL SCANEO DE VULNERABILIDADES ************	|
func (s *StorageDB) Create(entity dto.RepositoryDTO) (string, error){
	query := "INSERT INTO classification (REPOSITORY_URL, FILE, INFORMATION_TYPE, NUMBER_OF_LINES, DETECTION_DATE, JOB_ID, REPOSITORY_OWNER, AMOUNT_DETECTIONS) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return "", err
	}

	res, err := stmt.Exec(entity.RepositoryUrl, entity.File, entity.InformationType, entity.LinesNumber, entity.DetectionDate,entity.JobId, entity.RepositoryOwner, entity.AmountDetections)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return "", err
	}
	
	return entity.JobId, nil
}

// ************** FUNCION PARA TRAER LOS REGISTROS SEGUN EL JOB_ID ************	|
func (s *StorageDB) FindByJobId(jobId string) ([]domain.Entity, error) {
	var res []domain.Entity

	query := fmt.Sprintf("SELECT * FROM classification WHERE JOB_ID = %s;", jobId)
	rows, err := s.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	
	for rows.Next() {
		var entityDB domain.Entity
		var detectionDate []uint8
		err := rows.Scan(
			&entityDB.ID,
			&entityDB.RepositoryUrl,
			&entityDB.File,
			&entityDB.InformationType,
			&entityDB.LinesNumber,
			&detectionDate,
			&entityDB.JobId,
			&entityDB.RepositoryOwner,
			&entityDB.AmountDetections,
		)
		if err != nil {
			log.Fatal(err)
		}

		//* Conversion del tipo []unit8 (MySQL para fecha y Hora) --> time.Time (domain.Entity)
		entityDB.DetectionDate, err = time.Parse("2006-01-02 15:04:05", string(detectionDate))
		res = append(res, entityDB)
	}

	return res, nil
}