# Legal Translator 📜✨

## A modern approach to translating legal jargon into plain language

![Docker](https://img.shields.io/badge/Docker-1D63ED?style=for-the-badge&logo=docker&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-00A9D1?style=for-the-badge&logo=go&logoColor=white)
![Docker Model Runner](https://img.shields.io/badge/Docker_Model_Runner-1D63ED?style=for-the-badge&logo=docker&logoColor=white)

Legal Translator is a powerful application that leverages Docker Model Runner's AI capabilities to transform complex legal language into clear, everyday Portuguese. With its simple interface and robust backend, Legal Translator makes legal documents accessible to everyone.

## 🚀 Features

- **100% Local Processing** - All translations happen on your machine, ensuring complete privacy
- **Lightning Fast** - Get translations instantly without external API latency
- **Cost Effective** - No subscription fees or API costs
- **User-friendly Interface** - Clean, responsive design with both vertical and horizontal layouts
- **Docker Model Runner Integration** - Leverages Docker's AI capabilities for high-quality translations

## 📋 Prerequisites

- Docker Desktop 4.40+ with Docker Model Runner enabled
- Go 1.24+
- Internet connection (for initial model download only)

## 🔧 Installation

### 1. Clone the repository

```bash
git clone https://github.com/rflpazini/articles/translator.git
cd translator
```

### 2. Enable Docker Model Runner

1. Open Docker Desktop
2. Go to Settings > Features in development > Beta
3. Check "Enable Docker Model Runner"
4. Enable "Host-side TCP support" (optional, but recommended)
5. Apply and restart Docker Desktop

### 3. Start the application with Docker Compose

```bash
docker-compose up -d
```

This will start both the frontend and backend services. The application will be available at http://localhost:8080.

## 🏗️ Project Structure

```
translator/
├── backend/             # Go backend using Echo framework
│   ├── cmd/             # Application entry points
│   ├── internal/        # Internal packages
│   │   ├── api/         # API layer
│   │   ├── config/      # Configuration
│   │   ├── model/       # Data models
│   │   └── service/     # Business logic
│   ├── Dockerfile       # Backend Docker configuration
│   └── go.mod           # Go module definition
├── frontend/            # Web interface
│   ├── index.html       # Main HTML file
│   ├── Dockerfile       # Frontend Docker configuration
│   └── nginx.conf       # Nginx configuration
└── compose.yaml         # Docker Compose configuration
```

## 🔌 API Reference

### Translate Text

```
POST /api/translate
```

**Request Body:**

```json
{
  "text": "Fica o réu condenado a arcar com o ônus da sucumbência."
}
```

**Response:**

```json
{
  "translation": "O réu deve pagar as custas do processo."
}
```

### Health Check

```
GET /api/health
```

**Response:**

```
OK
```

## 🧩 Technology Stack

- **Frontend**: HTML, CSS, JavaScript, Bootstrap 5
- **Backend**: Go 1.24+, Echo Framework
- **AI**: Docker Model Runner with llama.cpp
- **Containerization**: Docker, Docker Compose
- **Web Server**: Nginx

## 🛣️ Roadmap

- [ ] Support for more languages
- [ ] Custom model configuration
- [ ] Document upload feature
- [ ] Translation history
- [ ] User accounts and saved translations
- [ ] Performance optimizations for larger documents

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://rflpazini.mit-license.org/) file for details.

## 🙏 Acknowledgments

- Docker team for creating Docker Model Runner
- The Go community for the excellent Echo framework
- All contributors who have helped shape this project

---

Made with ❤️ using Docker Model Runner