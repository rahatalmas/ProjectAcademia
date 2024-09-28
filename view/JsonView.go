package view

import(
	"encoding/json"
	"net/http"
)

func JsonView(w http.ResponseWriter, r* http.Request, data any){
	w.Header().Set("Content-Type","application/json");
	w.WriteHeader(http.StatusOK);

	err := json.NewEncoder(w).Encode(data);
	if err!=nil{
		http.Error(w,err.Error(), http.StatusInternalServerError);
	}
}