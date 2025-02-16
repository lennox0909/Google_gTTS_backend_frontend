# 使用 Node.js 作為基礎映像
FROM node:18

# 設定工作目錄
WORKDIR /app

# 複製 package.json 和 package-lock.json，先安裝依賴
COPY package.json package-lock.json ./
RUN npm install

# 複製前端程式碼
COPY . .

# 編譯 Next.js 應用
RUN npm run build

# 啟動 Next.js 應用
CMD ["npm", "run", "start"]
