package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/rs/cors"
)

type TTSRequest struct {
	Text     string `json:"text"`
	Language string `json:"language"`
	Speed    string `json:"speed"`
	Voice    string `json:"voice"`
}

func textToSpeechHandler(w http.ResponseWriter, r *http.Request) {
	var req TTSRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 確保 data 資料夾存在
	outputDir := "data"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create data directory", http.StatusInternalServerError)
		return
	}

	// 生成音檔名
	fileName := filepath.Join(outputDir, "output.mp3")

	// 使用 gTTS 生成音檔
	cmd := exec.Command("gtts-cli", req.Text, "--lang", req.Language, "--output", fileName)
	if req.Speed == "slow" {
		cmd.Args = append(cmd.Args, "--slow")
	}

	if err := cmd.Run(); err != nil {
		http.Error(w, "Failed to generate speech", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Speech generated", "file": fileName})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/generate", textToSpeechHandler)

	// 啟用 CORS，允許前端存取 API
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(mux)

	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", handler)
}
