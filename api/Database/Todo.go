package Database

import (
	"auth-api/Form"
	"auth-api/Utils"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func CreateTodo(title, description, color, user, path string) (string, error) {
	query := "INSERT INTO todo (title, description, color, token, user, path) VALUES (?, ?, ?, ?, ?, ?)"
	_, get := GetColumn(path)
	if get == nil {
		Token := Utils.Token(16)
		_, err := db.Exec(query, title, description, color, Token, user, path)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return Token, nil
	} else {
		return "", get
	}
}

func GetTodo(user string) ([]Form.Todo, error) {
	query := "SELECT title, description, color, token, user, path FROM todo WHERE user = ?"
	rows, err := db.Query(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Form.Todo

	for rows.Next() {
		var todo Form.Todo
		err := rows.Scan(&todo.Title, &todo.Description, &todo.Color, &todo.Token, &todo.User, &todo.Path)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func DeleteTodo(token string) error {
	query := "DELETE FROM todo WHERE token = ? "
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
		return fmt.Errorf("todo not found or already deleted")
	}

	return nil
}
