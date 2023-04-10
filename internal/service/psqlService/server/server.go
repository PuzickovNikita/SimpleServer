package server

import (
	"SimpleServer/internal/service/psqlService/internal/psqlDriver/jsons"
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

func (s *Server) InsertJson(jsonBody *jsons.PostReq) (int, error) {
	var newId int
	if err := s.db.QueryRow(context.Background(), fmt.Sprintf("insert into %s(\"body\") values('%s') returning key", jsonBody.Table, jsonBody.Body)).Scan(&newId); err != nil {
		return -1, err
	}
	return newId, nil
}

func (s *Server) GetJson(jsonReq *jsons.GetReq) (jsons.GetResp, error) {
	var jsonResp jsons.GetResp
	if err := s.db.QueryRow(context.Background(),
		fmt.Sprintf("select key,body from %s where key = %d", jsonReq.Table, jsonReq.Key)).Scan(&jsonResp.Key, &jsonResp.Body); err != nil {
		return jsonResp, err
	}
	return jsonResp, nil
}

func (s *Server) Delete(table string, key int) (int64, error) {
	tag, err := s.db.Exec(context.Background(), fmt.Sprintf("delete from %s where key=%d", table, key))
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
