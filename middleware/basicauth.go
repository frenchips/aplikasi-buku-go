package middleware

import (
	"fmt"
	"net/http"
)

// Fungi Log yang berguna sebagai middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

// Fungsi GetMovie untuk mengampilkan text string di browser
func GetMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("<h1>Anda Berhasil Mengakses Fungsi GetMovie() </h1>"))
	}
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/", Auth(http.HandlerFunc(GetMovie)))

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
