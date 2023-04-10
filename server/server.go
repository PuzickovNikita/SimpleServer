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

func (s *Server) GetJson(table string, jsonReq tables.SimpleTable) (tables.SimpleTable, error) {
	var jsonResp tables.SimpleTable
	if err := s.db.QueryRow(context.Background(),
		fmt.Sprintf("select key,body from %s where key = %d", table, jsonReq.Key)).Scan(&jsonResp.Key, &jsonResp.Body); err != nil {
		return jsonResp, err
	}
	return jsonResp, nil
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
