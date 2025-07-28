# Golang REST API Boilerplate 🇻🇳

## Table of Contents <!-- omit in toc -->

- [Descriptions](#descriptions)
- [Built With](#built-with)
- [Features](#features)
- [Installation](#installation)
- [Contact](#contact    )

## Descriptions
In this project, I built a template for the Backend part written in Go and using the Gin framework combined with many other features.

## Built With
<p align="left">
  <a href="https://golang.org" target="_blank">
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="golang" />
  </a>
  <a href="https://gin-gonic.com/" target="_blank">
    <img src="https://img.shields.io/badge/Gin-20232A?style=for-the-badge&logo=gin&logoColor=white" alt="gin" />
  </a>
  <a href="https://gorm.io/" target="_blank">
    <img src="https://img.shields.io/badge/GORM-02AEA1?style=for-the-badge&logo=gorm&logoColor=white" alt="gorm" />
  </a>
  <a href="https://www.postgresql.org/" target="_blank">
    <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="postgresql" />
  </a>
  <a href="https://swagger.io/" target="_blank">
    <img src="https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black" alt="swagger" />
  </a>
  <a href="https://www.docker.com/" target="_blank">
    <img src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white" alt="docker" />
  </a>
  <a href="https://github.com/features/actions" target="_blank">
    <img src="https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white" alt="github actions" />
  </a>
</p>

## Features

- [x] Database. Support [GORM](https://gorm.io/index.html) and [Postgres](https://github.com/lib/pq).
- [ ] Seeding.
- [ ] Mailing ([viper](https://github.com/spf13/viper)).
- [ ] Sign in and sign up via email.
- [ ] Social sign in (Apple, Facebook, Google).
- [x] Admin and User roles.
- [ ] Internationalization/Translations (I18N) ([go-i18n](https://github.com/nicksnyder/go-i18n)).
- [ ] File uploads. Support local and Amazon S3 drivers.
- [x] Swagger.
- [ ] E2E and units tests.
- [x] Docker.
- [x] CI (Github Actions).

## Installation
1. Clone the repository
```
git clone https://github.com/longnh462/go-gin-boilerplate
```
2. Install neccessary dependencies

```
go mod download
```

3. Create and copy this part into env file
```
cp env.example .env
```

## Contact
Feel free to react out through any of these channels:
<p align="left">
  <a href="mailto:longnh.uit@gmail.com" target="_blank">
    <img src="https://img.shields.io/badge/Email-D14836?style=for-the-badge&logo=gmail&logoColor=white" alt="email" />
  </a>
  <a href="https://www.linkedin.com/in/long-nguyen-hoang-1141b225b/" target="_blank">
    <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="linkedin" />
  </a>
  <a href="https://github.com/longnh462" target="_blank">
    <img src="https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white" alt="github" />
  </a>
</p>