# 🐉 Dragon Ball API

> **API REST con arquitectura hexagonal para gestión de personajes de Dragon Ball**

---

## 🚀 **Inicio Rápido**

```bash
git clone https://github.com/alvaroeds/amaris-dragon-ball.git
cd amaris-dragon-ball
docker-compose up -d --build
```

✅ **La API estará disponible en** `http://localhost:8080`

---

## ⚡ **¿Cómo funciona?**

La API implementa un **flujo inteligente de búsqueda**:

```
1. 🔄 Cache (Redis)      → Respuesta instantánea
2. 💾 Base datos (PostgreSQL) → Datos persistentes  
3. 🌐 API externa        → Dragon Ball API pública
4. 💿 Guarda automáticamente el resultado
```

---

## 📋 **Ejemplos de Uso**

### **Health Check**
```bash
curl http://localhost:8080/health
```

### **Crear Personaje (Primera vez)**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "Goku"}'
```
**➜ Respuesta:** `201 Created` con datos del personaje

### **Buscar Personaje Existente**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "Goku"}'
```
**➜ Respuesta:** `409 Conflict` - "Character already exists"

### **Nombre Incorrecto**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "Gok"}'
```
**➜ Respuesta:** `400 Bad Request` con sugerencias `["Goku", "Gohan", "Goten"]`

### **Personaje Inexistente**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "PersonajeInventado"}'
```
**➜ Respuesta:** `400 Bad Request` con lista de personajes disponibles

---

## 🔗 **API Reference**

### **`POST /api/v1/characters`**
Busca o crea un personaje de Dragon Ball.

**Request Body:**
```json
{
  "name": "string"  // ✅ Requerido
}
```

**Response Codes:**
- `201 Created` - Personaje creado exitosamente
- `409 Conflict` - Personaje ya existe en la base de datos
- `400 Bad Request` - Nombre incorrecto o personaje no encontrado
- `500 Internal Server Error` - Error interno del servidor

### **`GET /health`**
Verifica el estado de los servicios.

**Response:**
```json
{
  "result": {
    "status": "healthy",
    "services": {
      "database": "ok",
      "cache": "ok"
    }
  }
}
```

---

## 🏗️ **Arquitectura**

```
📁 pkg/character/
├── 🎯 domain/          # Entidades y reglas de negocio
├── 🧠 application/     # Casos de uso (service.go)
└── 🔌 infrastructure/  # HTTP, DB, Cache, API externa

📁 internal/
├── ⚙️  config/          # Configuración de la aplicación
├── 💚 health/          # Health check endpoints
└── 🏭 infrastructure/  # Servicios compartidos (DB, HTTP, Response)
```

**Stack Tecnológico:**
- **Go 1.24** - Lenguaje principal
- **PostgreSQL 16** - Base de datos relacional
- **Redis 7** - Cache de alta performance
- **Chi Router** - HTTP routing ligero
- **Docker** - Containerización

---

## ⚙️ **Configuración Avanzada (Opcional)**

El proyecto **funciona sin configuración adicional**. Para personalizar:

```bash
cp env.example .env
# Editar .env con tus valores personalizados
docker-compose up -d --build
```

**Variables disponibles:**
- `DB_*` - Configuración PostgreSQL
- `REDIS_*` - Configuración Redis  
- `API_PORT` - Puerto del servidor (default: `8080`)
- `DRAGONBALL_API_URL` - URL de la API externa


## 💡 **Características Técnicas**

✅ **Arquitectura Hexagonal** - Separación clara entre dominio e infraestructura  
✅ **Cache Strategy** - Redis para optimización de performance  
✅ **Error Handling** - Manejo consistente de errores HTTP  
✅ **Health Checks** - Monitoreo de servicios críticos  
✅ **Graceful Shutdown** - Cierre ordenado de conexiones  
✅ **SQL Directo** - Sin ORM para mayor control y performance  

---

**🎯 Diseñado con Domain-Driven Design y principios de arquitectura limpia**