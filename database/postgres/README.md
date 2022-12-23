## Setup postgres database for mediateq

First, you will need to install PostgreSQL on your system. Then follow the following steps to setup your mediateq database.

### 1. Connected to the PostgreSQL server

```bash
sudo -u postgres psql
```

### 2. Create a new database

```bash
CREATE DATABASE [DATABASE_NAME];
```

Here DATABASE_NAME is the name of the database. The default name in the configuration file is `mediateq`.

### 3. Create a new user for the database,

```bash
CREATE USER [DATABASE_USER] WITH PASSWORD '[DATABASE_USER_PASSWORD]';
```

### 4. Grant the new user privileges to access the database

```bash
GRANT ALL PRIVILEGES ON [DATABASE_NAME] mydatabase TO [DATABASE_USER];
```

### 5. Create database tables

```bash
psql -U [DATABASE_USER] -d [DATABASE_NAME] -f /path/to/mediateq/database/postgres/schema.sql
```
