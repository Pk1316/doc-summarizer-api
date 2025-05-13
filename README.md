# 📚 File Summarization API with Gemini

This is a RESTful API built with Go  that accepts text or file inputs (like PDFs) and returns a summarized version using Google's Gemini API (multimodal LLM).

---

## 🚀 Features

- Upload documents (PDF, text, markdown, images).
- Accepts file via:
  - Form Data
  - Base64 in headers
  - Base64 in JSON body
- MIME type detection and fallback.
- Gemini 1.5 Flash integration.
- Supports prompts for tailored summarization.

---

## 🛠️ Installation

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/file-processing-api.git
cd file-processing-api
```

### 2. Set up environment variables

Create a `.env` file and add your Gemini API key:

```env
GOOGLE_API_KEY=your-api-key-here
```

> ⚠️ Never hardcode API keys. Use environment variables or secrets manager.

---

### 3. Install dependencies

```bash
go mod tidy
```

---

### 4. Run the server

```bash
go run main.go
```

---

## 📬 API Endpoints

### 🧠 `POST /summarize`

#### Accepts:

- Form Data (file upload)
- Base64 file via `File-Data` header
- JSON body with base64 file

#### Form Data Example (Postman):

```
POST /summarize
Content-Type: multipart/form-data
```

| Key     | Type    | Description               |
|---------|---------|---------------------------|
| file    | File    | PDF or text document      |
| prompt  | Text    | Custom prompt (optional)  |

---

#### JSON Body Example

```json
{
  "fileData": "data:application/pdf;base64,JVBERi0xLjcK...",
  "prompt": "Summarize this document in bullet points"
}
```

---

### ✅ `GET /health`

Health check endpoint:

```bash
curl http://localhost:8080/health
```

Response:

```json
{ "status": "ok" }
```



