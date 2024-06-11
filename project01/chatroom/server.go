package main
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	go h.run()
	r.HandleFunc("/ws", myws)
	if err:= http.ListenAndServe("127.0.0.1:8080", r); err!= nil {
		fmt.Println(err)
	}
}