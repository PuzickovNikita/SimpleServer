package server

import (
	"SimpleServer/internal/psqlDriver/tables"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	db *pgx.Conn
}

func NewServer(db *pgx.Conn) Server {
	return Server{db: db}
}

func (s *Server) InsertJson(table string, jsonBody tables.SimpleTable) (tables.SimpleTable, error) {
	var newId int
	if err := s.db.QueryRow(context.Background(),
		fmt.Sprintf("insert into %s(\"body\") values('%s') returning key",
			table, jsonBody.Body)).Scan(&newId); err != nil {
		return tables.SimpleTable{}, err
	}
	return tables.SimpleTable{Key: newId}, nil
}

func (s *Server) GetJson(table string, jsonReq tables.SimpleTable) ([]tables.SimpleTable, error) {
	var condition string
	if jsonReq.Key != 0 {
		condition = fmt.Sprintf("where key = %d", jsonReq.Key)
	}
	if jsonReq.Body != "" {
		if condition != "" {
			condition += " and "
		} else {
			condition += "where "
		}
		condition += fmt.Sprintf("body='%s'", jsonReq.Body)
	}
	rows, err := s.db.Query(context.Background(), fmt.Sprintf("select key,body from %s %s", table, condition))
	if err != nil {
		return nil, err
	}
	var result []tables.SimpleTable
	for rows.Next() {
		var newRow tables.SimpleTable
		if err := rows.Scan(&newRow.Key, &newRow.Body); err != nil {
			return nil, err
		}
		result = append(result, newRow)
	}
	return result, nil
}

func (s *Server) Delete(table string, jsonReq tables.SimpleTable) (int64, error) {
	if jsonReq.Key == 0 {
		return 0, fmt.Errorf("No key")
	}
	tag, err := s.db.Exec(context.Background(), fmt.Sprintf("delete from %s where key=%d", table, jsonReq.Key))
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
