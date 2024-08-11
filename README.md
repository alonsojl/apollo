# Go RESTful API (Apolo)

Este microservicio ...

## Install
- Instalar la Interfaz de la Línea de Comandos de AWS ([AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)) para administrar los servicios de AWS.

```bash
make aws-cli
```

- Instalar la interfaz de línea de comandos del modelo de aplicación sin servidor de AWS ([AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)) para ejecutar comandos en el directorio de tu proyecto de aplicación AWS SAM y eventualmente convertirlo en tu aplicación serverless. 

```bash
make aws-sam
```

## Configuration
Para iniciar el entorno correctamente, es necesario generar el archivo `.env` utilizando `.example` como referencia y ejecutar el siguiente comando:

```bash
make up
```

## Documentation
Después de haber iniciado los servicios del contenedor, acceda al [Swagger](http://localhost:8082) en su navegador para obtener la documentación detallada de los endpoints.

## Envs
La siguiente lista detalla las variables de entorno empleadas durante la ejecución del microservicio `apolo`.

| Nombre          | Tipo      | Descripción                                              | Ejemplo  |
|-----------------|-----------|----------------------------------------------------------|----------|
| LOG_LEVEL       | STRING    | Nivel de detalle de los registros generados.             |info|