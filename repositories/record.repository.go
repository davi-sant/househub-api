package repositories

import (
	"context"
	"database/sql"
	"os"

	"github.com/davi-sant/househub-go/models"
)

type RecordRepository struct {
	DB *sql.DB
}

func NewRecordRepository(db *sql.DB) *RecordRepository {
	return &RecordRepository{DB: db}
}

func (r *RecordRepository) Create(ctx context.Context, rc models.RecordCreate) (*models.Record, error) {
	query, err := os.ReadFile("queries/record/create.sql")

	if err != nil {
		return nil, err
	}

	var newRecord models.Record
	err = r.DB.QueryRowContext(
		ctx,
		string(query),
		rc.RazaoSocial,
		rc.NomeFantasia,
		rc.CpfCnpj,
		rc.Telefone,
		rc.Email,
		rc.Endereco,
	).Scan(
		&newRecord.ID,
		&newRecord.RazaoSocial,
		&newRecord.NomeFantasia,
		&newRecord.CpfCnpj,
		&newRecord.Telefone,
		&newRecord.Email,
		&newRecord.Endereco,
		&newRecord.CriadoEm,
	)

	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}

func (r *RecordRepository) FindAll(ctx context.Context) ([]models.Record, error) {
	query, err := os.ReadFile("queries/record/find_all.sql")

	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, string(query))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var records []models.Record
	for rows.Next() {
		var record models.Record
		err := rows.Scan(
			&record.ID,
			&record.RazaoSocial,
			&record.NomeFantasia,
			&record.CpfCnpj,
			&record.Telefone,
			&record.Email,
			&record.Endereco,
			&record.CriadoEm,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (r *RecordRepository) FindById(ctx context.Context, fr models.FindRecordById) (*models.Record, error) {
	query, err := os.ReadFile("queries/record/find_by_id.sql")
	if err != nil {
		return nil, err
	}
	var record models.Record
	err = r.DB.QueryRowContext(ctx, string(query), fr.ID).Scan(
		&record.ID,
		&record.RazaoSocial,
		&record.NomeFantasia,
		&record.CpfCnpj,
		&record.Telefone,
		&record.Email,
		&record.Endereco,
		&record.CriadoEm,
	)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *RecordRepository) Update(ctx context.Context, ru *models.RecordUpdate) (*models.Record, error) {
	query, err := os.ReadFile("/queries/record/update.sql")

	if err != nil {
		return nil, err
	}

	var record models.Record

	err = r.DB.QueryRowContext(
		ctx,
		string(query),
		ru.RazaoSocial,
		ru.NomeFantasia,
		ru.Telefone,
		ru.Email,
		ru.Endereco,
	).Scan(
		&record.ID,
		&record.RazaoSocial,
		&record.NomeFantasia,
		&record.CpfCnpj,
		&record.Telefone,
		&record.Email,
		&record.Endereco,
		&record.CriadoEm,
	)
	if err != nil {
		return nil, err
	}

	return &record, nil
}
