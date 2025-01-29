FROM postgres:latest

ENV POSTGRES_DB=under_hex
ENV POSTGRES_USER=under_admin
ENV POSTGRES_PASSWORD=10482810

COPY script.sql /docker-entrypoint-initdb.d/

EXPOSE 5432