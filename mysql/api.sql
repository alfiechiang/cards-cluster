FROM mysql:latest
COPY api.sql /docker-entrypoint-initdb.d/api.sql