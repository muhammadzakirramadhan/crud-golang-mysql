package mahasiswa

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/muhammadzakirramadhan/crud-golang-mysql/config"
	"github.com/muhammadzakirramadhan/crud-golang-mysql/models"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Data
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {

	var mahasiswas []models.Mahasiswa

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mahasiswa models.Mahasiswa
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.NIM,
			&mahasiswa.Name,
			&mahasiswa.Semester,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}

	return mahasiswas, nil
}

// Insert data Mahasiswa
func Insert(ctx context.Context, mhs models.Mahasiswa) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nim, name, semester, created_at, updated_at) values(%v,'%v',%v,'%v','%v')", table,
		mhs.NIM,
		mhs.Name,
		mhs.Semester,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Data Mahasiswa
func Update(ctx context.Context, mhs models.Mahasiswa) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nim = %d, name ='%s', semester = %d, updated_at = '%v' where id = '%d'",
		table,
		mhs.NIM,
		mhs.Name,
		mhs.Semester,
		time.Now().Format(layoutDateTime),
		mhs.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Data Mahasiwa
func Delete(ctx context.Context, mhs models.Mahasiswa) error {
 
    db, err := config.MySQL()
 
    if err != nil {
        log.Fatal("Can't connect to MySQL", err)
    }
 
    queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, mhs.ID)
 
    s, err := db.ExecContext(ctx, queryText)
 
    if err != nil && err != sql.ErrNoRows {
        return err
    }
 
    check, err := s.RowsAffected()
    fmt.Println(check)
    if check == 0 {
        return errors.New("id tidak ada")
    }
 
    return nil
}