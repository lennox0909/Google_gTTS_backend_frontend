# 語音轉換應用 (Google gTTS + Next.js + Golang)

## 📌 專案簡介
本專案是一個基於 **Google gTTS** 的 **文字轉語音** 應用，後端使用 **Golang**，前端使用 **Next.js (TypeScript)**，並通過 **Docker** 進行容器化部署。

## gTTS 並不是 Google 官方正式提供的 API
Google 會偵測異常流量，如果你的 IP 或帳號在短時間內發送大量請求（例如：
- 批量轉換大量文字為語音
- 太頻繁地呼叫 gtts
- 多個 IP 來自相同的來源

Google 可能會啟動自動防禦機制（Rate Limiting），導致請求被封鎖或返回錯誤。

## 📂 專案結構
```
project_root/
│── backend/          # 後端 Golang + gTTS
│   ├── main.go       # 後端 API 服務
│   ├── Dockerfile    # 後端 Docker 設定
│   ├── go.mod        # Golang 依賴管理
│   ├── go.sum        # Golang 依賴鎖定文件
│── frontend/         # 前端 Next.js + TypeScript
│   ├── pages/        # 前端頁面
│   ├── Dockerfile    # 前端 Docker 設定
│   ├── package.json  # 前端依賴管理
│   ├── tsconfig.json # TypeScript 設定
│── voice_output/     # 存放生成的語音檔案
│── docker-compose.yml  # Docker Compose 配置
│── README.md         # 專案說明文件
```

## 🚀 如何運行
### 1️⃣ **安裝 Docker & Docker Compose**
請確保你的系統已安裝 [Docker](https://www.docker.com/) 和 [Docker Compose](https://docs.docker.com/compose/)。

### 2️⃣ **啟動專案**
在專案根目錄執行以下指令，建立並啟動後端與前端：
```sh
docker-compose up --build
```
此命令將會：
- **構建後端 (Golang)** 並安裝 gTTS、ffmpeg
- **構建前端 (Next.js)** 並安裝所有依賴
- **建立 `voice_output/` 資料夾** 用於存放語音檔案

### 3️⃣ **訪問應用**
- **前端 UI**：[http://localhost:3000](http://localhost:3000)
- **後端 API**：[http://localhost:8081/generate](http://localhost:8081/generate)

## 🔄 API 介面說明
### **1️⃣ 生成語音**
- **API 路徑**：`POST /generate`
- **請求格式 (JSON)**：
```json
{
  "text": "你好，世界！",
  "language": "zh-TW",
  "speed": 1.0
}
```
- **回應格式 (JSON)**：
```json
{
  "message": "Speech generated",
  "file": "voice_output/output.mp3"
}
```

### **2️⃣ 下載語音檔案**
- **API 路徑**：`GET /download?file=voice_output/output.mp3`
- **下載語音檔案** 直接開啟 `output.mp3` 進行播放。

## 🛠️ 開發與測試
### **本地開發**
若要在本機執行後端與前端，可手動執行：
#### **後端 (Golang)**
```sh
cd backend
export PORT=8081
mkdir -p ../voice_output
go run main.go
```
#### **前端 (Next.js)**
```sh
cd frontend
npm install
npm run dev
```

## 📝 注意事項
- 若 `voice_output/` 沒有自動建立，可手動創建：
  ```sh
  mkdir -p voice_output
  ```
- 若後端無法安裝 `gTTS`，請確保 `Python3` 和 `pip3` 已安裝。
- 若前端無法啟動，請確保 `Node.js 18+` 和 `npm` 已安裝。

## 📌 授權 & 參考
- 本專案基於 **MIT License** 開源。
- 使用的技術：Google gTTS、Next.js、Golang、Docker。

🚀 **歡迎 PR 或提供建議！** 🎉

