package server

import (
	"SimpleServer/internal/psqlDriver/tables"
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeJson(w http.ResponseWriter, r *http.Request, req *tables.SimpleTable) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Print("Faild to decode json\n")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func (s *Server) PsqlHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/PSQL/JSON" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	var table string = r.FormValue("table")
	if table == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req = tables.SimpleTable{}
	switch r.Method {
	case "GET":
		if err := decodeJson(w, r, &req); err != nil {
			return
		}
		resp, err := s.GetJson(table, req)
		if err != nil {
			fmt.Printf("Unable to select %v: \n %v", req, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			fmt.Printf("Failed to decode %v\n%v\n", resp, err)
			http.Error(w, fmt.Sprintf("Server unable to encode data: %s", err.Error()), http.StatusInternalServerError)
		}
	case "POST":
		if err := decodeJson(w, r, &req); err != nil {
			return
		}
		res, err := s.InsertJson(table, req)
		if err != nil {
			fmt.Printf("Unable to insert %v:\n%v", req, err)
			http.Error(w, fmt.Sprintf("Server is unable to INSERT:%s", err.Error()), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Print("Failed to encode json\n")
			http.Error(w, fmt.Sprintf("Server is unable to encode response. Added res is %d", res),
				http.StatusInternalServerError)
		}
	case "DELETE":
		if err := decodeJson(w, r, &req); err != nil {
			return
		}
		res, err := s.Delete(table, req)
		if err != nil {
			fmt.Printf("Unable to delete %s", err.Error())
			http.Error(w, "Delete failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fmt.Sprintf("Rows deleted: %d", res))
	default:
		fmt.Print(w, "SMTH ELSE")
	}
}
