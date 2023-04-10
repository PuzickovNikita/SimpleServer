package server

import (
	"SimpleServer/internal/service/psqlService/internal/psqlDriver/jsons"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func decodeJson(w http.ResponseWriter, r *http.Request, req any) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Print("Faild to decode json\n")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func getTableKey(r *http.Request) (string, int, error) {
	table := r.FormValue("table")
	key, err := strconv.Atoi(r.FormValue("key"))
	if table == "" || key <= 0 || err != nil {
		return "", 0, fmt.Errorf("wrong requst")
	}
	return table, key, nil
}

func (s *Server) PsqlHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/PSQL/JSON" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		table, key, err := getTableKey(r)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		req := jsons.GetReq{
			Table: table,
			Key:   key,
		}
		resp, err := s.GetJson(&req)
		if err != nil {
			fmt.Printf("Unable to select %v: \n %v", req, err)
			w.WriteHeader(500)
			return
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("Failed to decode %v\n%v\n", resp, err)
			http.Error(w, fmt.Sprintf("Server unable to encode data: %v", resp), 500)
		}
	case "POST":
		var req jsons.PostReq
		if err := decodeJson(w, r, &req); err != nil {
			return
		}
		key, err := s.InsertJson(&req)
		if err != nil {
			fmt.Printf("Unable to insert %v:\n%v", req, err)
			w.WriteHeader(500)
			return
		}
		if err := json.NewEncoder(w).Encode(jsons.PostResp{Key: key}); err != nil {
			fmt.Print("Failde to encode json\n")
			http.Error(w, fmt.Sprintf("Server is unable to encode response. Added key is %d", key), 500)
		}
	case "DELETE":
		table, key, err := getTableKey(r)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		res, err := s.Delete(table, key)
		if err != nil {
			fmt.Printf("Unable to delete %s:%d\n%v\n", table, key, err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		fmt.Fprintf(w, fmt.Sprintf("Rows deleted: %d", res))
	default:
		fmt.Print(w, "SMTH ELSE")
	}
}
