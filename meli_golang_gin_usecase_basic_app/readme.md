# Meli Challenge 
! Antes de correr el proyecto ! 
1. Crear archivo .env dentro de ./cmd/api con los campos PORT, TOKEN (AuthenticationAPI), USERDB y PASSDB
2. Run script.sql en MySQL Workbench

### HOW TO RUN
1 - Build the Docker image:
    docker build . -t my-golang-app 
2 - Run the Docker image:
    docker run -it -p 8080:8080 my-golang-app 
3 - Open the next URLs for verifying that all works fine:
    http://localhost:8080/ping
  