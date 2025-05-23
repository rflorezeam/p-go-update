# Microservicio de Actualización de Libros

## Desarrollador
Ricardo Florez

## Descripción
Este microservicio es responsable de actualizar la información de libros existentes en el sistema. Forma parte de una arquitectura de microservicios para la gestión de una biblioteca digital.

## Características
- Implementado en Go 1.21
- Arquitectura limpia (Clean Architecture)
- Endpoints RESTful
- Integración con MongoDB
- Despliegue en Kubernetes

## Estructura del Proyecto
```
.
├── config/         # Configuración de la base de datos
├── models/         # Modelos de datos
├── repositories/   # Capa de acceso a datos
├── services/      # Lógica de negocio
├── k8s/           # Configuración de Kubernetes
└── tests/         # Pruebas unitarias
```

## API Endpoint
- **PUT** `/libros/{id}`
  - Puerto: 30084 (NodePort)
  - Actualiza un libro existente por su ID

### Ejemplo de Petición
```json
{
    "titulo": "Don Quijote de la Mancha",
    "autor": "Miguel de Cervantes Saavedra"
}
```

### Ejemplo de Respuesta
```json
{
    "id": "5f7b5e1b9d3e2a1b4c7d8e9f",
    "titulo": "Don Quijote de la Mancha",
    "autor": "Miguel de Cervantes Saavedra"
}
```

### Respuestas de Error
```json
{
    "error": "Libro no encontrado"
}
```

## Configuración Kubernetes
- Deployment con 3 réplicas
- Service tipo NodePort (30084)
- Conexión a MongoDB mediante Service Discovery

## Variables de Entorno
- MONGODB_URI: URI de conexión a MongoDB

## Despliegue
```bash
# Construir la imagen
docker build -t libro-update:latest .

# Desplegar en Kubernetes
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

## Pruebas
```bash
# Ejecutar pruebas unitarias
go test ./...

# Probar el endpoint
curl -X PUT -H "Content-Type: application/json" \
     -d '{"titulo":"Don Quijote de la Mancha","autor":"Miguel de Cervantes Saavedra"}' \
     http://localhost:30084/libros/5f7b5e1b9d3e2a1b4c7d8e9f
```

## Monitoreo
El servicio puede ser monitoreado mediante:
- Logs de Kubernetes
- Métricas de contenedor
- Estado del Service y Deployment

## Validaciones
- Verificación de existencia del libro
- Validación de campos requeridos
- Sanitización de datos de entrada
- Control de concurrencia 