# Yuki Kasir
Lorem

## Tech Stack
- Programming Language: [GO](https://go.dev/)
- HTTP Web Framework: [Gin-Gonic](https://gin-gonic.com/)
- Database: [PostgreSQL](https://www.postgresql.org/)
- Database Driver: [Gorm](https://gorm.io/)

## Development Setup

You need atleast using Golang v1.22.0

### Database
To set up the database for this project, you have two options: using Docker or WSL for a Linux environment. Follow the steps below to initialize the database.

#### Option 1: Using Docker
1. Install docker
2. Run this command:
```bash
docker pull postgres
docker run --name my_database -e POSTGRES_PASSWORD=mysecretpassword -d postgres
docker ps
```

#### Option 2: Using WSL for Linux Environment
1. Install WSL2 for windows (If you already using Linux as your main OS, then you don't need to download WSL)
2. Run this command:
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo service postgresql start
sudo service postgresql status
```

By following these steps, you can quickly set up the database for your Golang RESTful API project using either Docker or WSL.

### Project

#### Installation
To run this project, you need to install the required go package by run this command:
```bash
go mod tidy
```

#### env
For the environment file, simply copy `example.env` and rename it to `.env`.
- For the Server section, you can costumize the port and frontend host.
- For the Database section, you need to fill it up based on the database that you have initialize before.
- For the Services related section:
    - `JWT_SECRET_KEY`: you can generate random string for this.
    - `USER_ADMIN_EMAIL` & `USER_ADMIN_PASSWORD`: for the seeder as initial admin account.
    - `MAILER_EMAIL`: for the mailing system (login notification, reset password, etc.)
    - `MAILER_PASSWORD`: you can get your gmail App Password by follow this steps:
        - Create Gmail Account
        - Enable [2FA](https://myaccount.google.com/signinoptions/two-step-verification/enroll-welcome)
        - Create [App Password](https://myaccount.google.com/apppasswords)
    - `RESETPW_FE_ENDPOINT`: this is the endpoint from frontend for the reset password link that sent to user email.
- Lastly, for the Linux build section, this is only required when you want to build the app at linux OS. If you build using windows, you can comment this section.

#### Run the project
Doing all of the steps above, you have two options to run the app:
1. Use the [Makefile](https://gnuwin32.sourceforge.net/packages/make.htm) to start the application:
2. Run this command:
```bash
make
```

Or simply run the Go command:
```bash
go run main.go
```
