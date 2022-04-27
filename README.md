# skaypet-api

**Configuración de entorno de desarrollo**

1. Installar [taskfile](https://taskfile.dev/#/installation):
2. Establecer las variables de entorno config/ENV/dev/.env
3. Iniciar los contenedores para la base de datos y visualizador grafico
   > task db.start
   >
4. Crear tablas haciendo uso de los schema db/schema
5. Correr la aplicación
   > task run
   >
6. Verificar que la aplicación está corriendo
   > http://localhost:9000/api/v1.0/ping
   >
7. Ver documentación de la API
   > http://localhost:9000/api/v1.0/help
   >
8. Acceder a la aplicación en producción
      [http://ec2-3-238-20-158.compute-1.amazonaws.com:9000/api/v1.0/help](http://ec2-3-238-20-158.compute-1.amazonaws.com:9000/api/v1.0/help)
      [http://3.238.20.158:9000/api/v1.0/help](http://3.238.20.158:9000/api/v1.0/help)