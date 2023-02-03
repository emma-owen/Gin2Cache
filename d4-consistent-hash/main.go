package main
import (
	"fmt"
	"gin2cache"
	"log"
	"net/http"
)

var db = map[string]string {
	"Tom": "630",
	"Jack": "589",
	"Sam": "567",
}

func main(){
	gin2cache.NewGroup("scores",2 << 10, gin2cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key",key)
			if v, ok := db[key]; ok {
				return []byte(v),nil
			}
			return nil,fmt.Errorf("%s not exist",key)
		}))

	addr := "localhost:9999"
	peers := gin2cache.NewHTTPPool(addr)
	log.Println("gin2cache is running at ",addr)
	log.Fatal(http.ListenAndServe(addr,peers))
}