import { useState } from "react";

export default function Home() {
  const [text, setText] = useState("");
  const [language, setLanguage] = useState("zh-TW");
  const [speed, setSpeed] = useState(1);
  const [audioFile, setAudioFile] = useState("");

  const handleGenerate = async () => {
    if (!text.trim()) {
      alert("請輸入文字！");
      return;
    }

    try {
      const response = await fetch("http://localhost:8081/generate", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ text, language, speed }),
      });

      if (!response.ok) throw new Error("語音轉換失敗");

      const data = await response.json();
      setAudioFile(data.file);
      window.open(`http://localhost:8081/download?file=${data.file}`, "_blank");
    } catch (error) {
      console.error("請求錯誤:", error);
      alert("請求失敗，請檢查後端是否運行！");
    }
  };

  return (
    <div style={{ display: "flex", flexDirection: "column", alignItems: "center", width: "100%", padding: "20px" }}>
      <h1 style={{ fontSize: "2rem", marginBottom: "20px" }}>文字轉語音</h1>
      <textarea
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="請輸入文字"
        style={{ width: "60%", height: "50vh", resize: "none", padding: "10px", fontSize: "1.2rem", marginBottom: "15px" }}
      />
      <select value={language} onChange={(e) => setLanguage(e.target.value)} style={{ fontSize: "1rem", padding: "5px", marginBottom: "10px" }}>
        <option value="zh-TW">中文（台灣）</option>
        <option value="en">英文</option>
      </select>
      <input
        type="range"
        min="0.5"
        max="2"
        step="0.5"
        value={speed}
        onChange={(e) => setSpeed(parseFloat(e.target.value))}
        style={{ width: "50%", marginBottom: "10px" }}
      />
      <p style={{ fontSize: "1.2rem", marginBottom: "10px" }}>當前速度: {speed} 倍</p>
      <button
        onClick={handleGenerate}
        style={{ fontSize: "1rem", padding: "10px 20px", backgroundColor: "#007bff", color: "white", border: "none", borderRadius: "5px", cursor: "pointer" }}
      >
        生成語音並下載
      </button>
    </div>
  );
}
