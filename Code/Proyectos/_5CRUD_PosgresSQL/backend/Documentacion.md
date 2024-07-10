
1 se usa docker compose por lo tanto para conectarse con la terminal puedes usar lo siguiente.

Nombre del contenedor :       postgres1
Nombre de la base de datos:   test
Usuaro                        caleb
Contraseña                    12345
puerto                        8080


Abre una terminal o línea de comandos en el directorio raíz de tu proyecto donde se encuentra el archivo docker-compose.yml.

Ejecuta el siguiente comando para iniciar los servicios definidos en el archivo docker-compose.yml:

Copy code
docker-compose up -d

Luego, para conectarse a la terminal del contenedor de PostgreSQL, puedes ejecutar el siguiente comando:

Copy code
sudo  docker exec -it postgres1 bash

luego
psql -U caleb -d test

para usar la  base  de  datos 

CREATE TABLE persona (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    age INTEGER,
    gmail VARCHAR(255)
);

para insertar datos  

INSERT INTO persona (name, age, gmail) VALUES ('Juan', 32, 'juan@gmail.com');


para hacer un testeo rapido 

curl -X POST -H "Content-Type: application/json" -d '{"Id":7,"Name":"Carlos","Age":28,"Gmail":"carlos@gmail.com"}' http://localhost:3080/api/datos/new    

// put method
curl -X PUT -H "Content-Type: application/json" -d '{"Id":4,"Name":"Ana","Age":25,"Gmail":"ana@gmail.com"}' http://localhost:3080/api/datos/4

