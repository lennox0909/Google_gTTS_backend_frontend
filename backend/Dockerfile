# 使用 Golang 作為基礎映像
FROM golang:1.20

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum，先下載依賴
COPY go.mod go.sum ./
RUN go mod tidy

# 複製程式碼
COPY . .

# 安裝 Python3、gTTS 和 ffmpeg
RUN apt-get update && apt-get install -y python3 python3-pip ffmpeg && \
    python3 -m pip install --break-system-packages gtts

# 編譯應用程式
RUN go build -o main .

# 建立 voice_output 資料夾
RUN mkdir -p /app/voice_output && chmod -R 777 /app/voice_output

# 啟動應用
CMD ["/app/main"]
