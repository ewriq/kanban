package Database

import (
	"auth-api/Form"
	"auth-api/Utils"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
func CreateColumn(title, description, color, user string) (string, error) {
	query := "INSERT INTO columns (title, description, color, token, user) VALUES (?, ?, ?, ?, ?)"
	Token := Utils.Token(16)
	_, err := db.Exec(query, title, description, color, Token, user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return Token, nil
}

func GetColumn(path string) (Form.Column, error) {
	query := "SELECT * FROM columns WHERE token = ?"
	row := db.QueryRow(query, path)

	var column Form.Column
	err := row.Scan(&column.Title, &column.Description, &column.Color, &column.Token, &column.User)
	if err != nil {
		if err == sql.ErrNoRows {
			return Form.Column{}, fmt.Errorf("kullanıcı bulunamadı")
		}
		return Form.Column{}, err
	}
	return column, nil
}
func ListColumn(user string) ([]Form.Column, error) {
	query := "SELECT title, description, color, token, user FROM columns WHERE user = ?"
	rows, err := db.Query(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []Form.Column

	for rows.Next() {
		var column Form.Column
		err := rows.Scan(&column.Title, &column.Description, &column.Color, &column.Token, &column.User)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}

func DeleteColumn(token string) error {
	query := "DELETE FROM columns WHERE token = ?"
	res, err := db.Exec(query, token)
	if err != nil {
		fmt.Println("Query error:", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected error:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("column not found or already deleted")
	}

	return nil
}
