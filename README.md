# EI-GO

EI-GO is an api written in golang with support of gin to handle backend logic of my app to learn new english words. App logic is easy because we have sets which we can attach to categoreis which we can create my on our own.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Optional routes](#optional-routes)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [License](#license)

---

## Features

- **Categories**: Create,delete, and retrieve categories.
- **Sets**: Create,delete, and retrieve sets.
- **Elements**: Retrieve elements.
---

## Technologies Used

- **Go**: Backend runtime.
- **Gin**: Web framework for building RESTful APIs.
- **PostgreSql**: Database for storing sets,categories and elements.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/slodkiadrianek/EI.git
   cd EI
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Create a .env file by copying .env.example:
   ```bash
   cp .env.example .env
   ```
4. Update the .env file with your environment-specific variables.
5. Start the development server:
   ```bash
   go run main.go
   ```
6. Build the project for production:
   ```bash
   go build -o ei
   ```
7. Start the production server:
   ```bash
   ./ei
   ```

## Docker

1.  Build docker image

```bash
docker build -t ei-go .
```

2. Run docker container

```bash
docker run -p 3000:3000 --env-file .env ei-go
```

## Environment Variables

- `DbLink` PostgreSql connection string.
- `PORT` Port on which app will be hosted.

## Pro structure

```bash
El-GO/
├── config/
├── controller/
├── DTO/
├── logs/
├── middleware/
├── migrations/
├── models/
├── repositories/
├── routes/
├── schema/
├── services/
├── tests/
├── utils/
├── .env
├── .env.test
├── .env.example
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```

