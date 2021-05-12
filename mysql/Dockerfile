FROM mysql:latest

COPY init.sql /docker-entrypoint-initdb.d
COPY my.cnf /etc
EXPOSE 3306