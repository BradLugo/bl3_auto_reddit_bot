package bot

import (
	"fmt"
)

type Code struct {
	CodeId   string
	CodeType  string
	Code string
}

func (db *DB) AllCodes() ([]*Code, error) {
	rows, err := db.Query("SELECT * FROM Codes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	codes := make([]*Code, 0)
	for rows.Next() {
		c := new(Code)
		err := rows.Scan(&c.CodeId, &c.CodeType, &c.Code)
		if err != nil {
			return nil, err
		}
		codes = append(codes, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return codes, nil
}

func (db *DB) FindCodeByCode(code string) (*Code, error) {
	row, err := db.Query(fmt.Sprintf("SELECT * FROM Codes WHERE Code=%s", code))
	if err != nil {
		return nil, err
	}
	defer row.Close()

	c := new(Code)
	err = row.Scan(&c.CodeId, &c.CodeType, &c.Code)
	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return c, nil
}

func (db *DB) InsertCode(code string) error {
	rows, err := db.Query("SELECT * FROM Codes")
	if err != nil {
		return err
	}
	defer rows.Close()

	codes := make([]*Code, 0)
	for rows.Next() {
		c := new(Code)
		err := rows.Scan(&c.CodeId, &c.CodeType, &c.Code)
		if err != nil {
			return nil, err
		}
		codes = append(codes, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return codes, nil
}