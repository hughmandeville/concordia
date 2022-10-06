FROM mariadb:10.2

# Environment variables
ENV MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=true
ENV MYSQL_DATABASE=concordia

# Load DB schema into the DB on startup
ADD ["concordia_db.sql", "/docker-entrypoint-initdb.d/sources.sql"]
