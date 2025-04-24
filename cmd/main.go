package main

import (
	"log"

	"lexigo/db"
	parser "lexigo/parser/cedict"
	"lexigo/utils"
)

func main() {
	// Bước 1: Load .env (đã chạy trong init của utils)
	// Bước 2: Kết nối DB
	db.Init()

	// Bước 3: Lấy đường dẫn file từ biến môi trường
	cedictPath := utils.GetEnv("CEDICT_PATH", "cedict_ts.u8")

	// Bước 4: Parse
	if err := parser.ParseCCCEDICT(cedictPath, db.DB); err != nil {
		log.Fatal("Parse error:", err)
	}

	// Bước 5: Đếm số lượng từ
	count, err := parser.CountCCCEDICT(cedictPath)
	if err != nil {
		log.Fatal("Count error:", err)
	}
	log.Printf("✅ Done parsing CC-CEDICT. Found %d entries.", count)
}
