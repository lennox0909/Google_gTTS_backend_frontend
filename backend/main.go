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
    Text     string  `json:"text"`
    Language string  `json:"language"`
    Speed    float64 `json:"speed"`
    Voice    string  `json:"voice"`
}

func textToSpeechHandler(w http.ResponseWriter, r *http.Request) {
    var req TTSRequest
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // 確保專案根目錄下的 voice_output 資料夾存在
    outputDir := "./voice_output"
    if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
        http.Error(w, "Failed to create voice_output directory", http.StatusInternalServerError)
        return
    }

    // 生成音檔名
    originalFile := filepath.Join(outputDir, "output.mp3")
    processedFile := filepath.Join(outputDir, "output_processed.mp3")

    // 使用 gTTS 生成音檔
    cmd := exec.Command("gtts-cli", req.Text, "--lang", req.Language, "--output", originalFile)
    if req.Speed == 0.5 {
        cmd.Args = append(cmd.Args, "--slow")
    }

    if err := cmd.Run(); err != nil {
        http.Error(w, "Failed to generate speech", http.StatusInternalServerError)
        return
    }

    // 調整音檔播放速度
    if req.Speed != 1 {
        cmd = exec.Command("ffmpeg", "-i", originalFile, "-filter:a", fmt.Sprintf("atempo=%.1f", req.Speed), "-y", processedFile)
        if err := cmd.Run(); err != nil {
            http.Error(w, "Failed to process audio speed", http.StatusInternalServerError)
            return
        }
    } else {
        processedFile = originalFile
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Speech generated", "file": processedFile})
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    filePath := r.URL.Query().Get("file")
    if filePath == "" {
        http.Error(w, "Missing file parameter", http.StatusBadRequest)
        return
    }
    http.ServeFile(w, r, filePath)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/generate", textToSpeechHandler)
    mux.HandleFunc("/download", downloadHandler)

    // 啟用 CORS
    handler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    }).Handler(mux)

    fmt.Println("Server is running on port 8081...")
    http.ListenAndServe(":8081", handler)
}
