package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fileServer := http.FileServer(http.Dir("./media"))
	http.Handle("/media/", http.StripPrefix("/media/", fileServer))

	http.HandleFunc("/api/files", func(w http.ResponseWriter, r *http.Request) {
		files := []string{}
		_ = filepath.Walk("./media", func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}

			webPath := strings.Replace(path, "\\", "/", -1)
			webPath = strings.TrimPrefix(webPath, "./")
			files = append(files, "/"+webPath)
			return nil
		})
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	})

	log.Println("Backend running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
