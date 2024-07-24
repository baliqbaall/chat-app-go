package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files (HTML, CSS, JS)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Endpoint untuk API chat
	http.HandleFunc("/api/messages", messagesHandler)

	// Start server
	log.Println("Server berjalan di: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Dummy handler for chat messages
func messagesHandler(w http.ResponseWriter, r *http.Request) {
	// Here you would interact with a database or data source
	// For now, we will return some static JSON
	jsonResponse := `{
		"results": {
			"room": {
				"name": "Product A",
				"id": 12456,
				"image_url": "https://picsum.photos/id/237/200/300",
				"participant": [
					{"id": "admin@mail.com", "name": "Admin", "role": 0},
					{"id": "agent@mail.com", "name": "Agent A", "role": 1},
					{"id": "customer@mail.com", "name": "king customer", "role": 2}
				]
			},
			"comments": [
				{"id": 885512, "type": "text", "message": "Selamat malam", "sender": "customer@mail.com"},
				{"id": 885513, "type": "text", "message": "Malam", "sender": "agent@mail.com"},
				{"id": 885514, "type": "text", "message": "Ada yang bisa saya bantu?", "sender": "agent@mail.com"},
				{"id": 885515, "type": "text", "message": "Saya ingin mengirimkan bukti pembayaran, karena diaplikasi selalu gagal", "sender": "customer@mail.com"},
				{"id": 885516, "type": "text", "message": "Baik, silahkan kirimkan lampiran bukti pembayarannya", "sender": "agent@mail.com"},
				{"id": 885517, "type": "image", "message": "Bukti pembayaran", "sender": "customer@mail.com", "image_url": "https://www.w3schools.com/html/pic_trulli.jpg"},
				{"id": 885518, "type": "video", "message": "Video panduan", "sender": "agent@mail.com", "video_url": "https://www.w3schools.com/html/mov_bbb.mp4"},
				{"id": 885519, "type": "pdf", "message": "Lampiran PDF", "sender": "agent@mail.com", "pdf_url": "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"}
			]
		}
	}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
