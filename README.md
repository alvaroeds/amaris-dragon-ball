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

## ⚡ **Cómo Funciona el Sistema**

La API utiliza una **estrategia de cache inteligente** para optimizar el rendimiento:

### **🔄 Flujo de Búsqueda de Personajes:**
```
📝 Request: {"name": "Goku"}
    ↓(se manda la request)
🔍 1. Buscar en Redis Cache (clave: "character:goku")
    ↓ (si no existe)
💾 2. Buscar en Base de Datos Local (PostgreSQL)
    ↓ (si no existe)
🌐 3. Consultar API Externa de Dragon Ball
    ↓ (evalúa respuesta)
💿 4. Guardar resultado en Cache:
    • ✅ Coincidencia exacta → BD + Cache personaje
    • ⚠️ Búsqueda parcial → Cache sugerencias  
    • ❌ Sin resultados → Cache "no encontrado"
    ↓
✅ 5. Responder al cliente
```

### **💡 Ventajas del Cache:**
- **⚡ Primera consulta:** `~300ms` (API externa + BD)
- **🚀 Consultas siguientes:** `~5ms` (desde cache)
- **🛡️ Protección API externa:** Evita saturar el servicio externo
- **📊 Sugerencias inteligentes:** Cache de búsquedas parciales como en la api externa
- **🛡️ Protección anti-spam:** Cachea búsquedas fallidas para evitar consultas repetitivas a la API externa

### **🔒 Ejemplos de Protección:**
```
❌ "PersonajeInventado" → Cache 30min → No más llamadas a API externa
⚠️ "Go" → Cache 1h → Sugerencias rápidas sin re-consultar
✅ "Goku" → Cache 24h → Respuesta instantánea desde cache
```

---

## 📋 **Ejemplos de Uso**

### **🩺 Health Check**
```bash
curl http://localhost:8080/api/v1/health
```
**➜ Respuesta 200:**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z",
  "services": {
    "database": {"status": "up", "message": "PostgreSQL connected"},
    "cache": {"status": "up", "message": "Redis connected"}
  }
}
```

---

### **✅ Crear Personaje (Primera vez - Éxito)**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "Goku"}'
```
**➜ Respuesta 201 Created:**
```json
{
  "result": {
    "id": 1,
    "external_id": 1,
    "name": "Goku",
    "race": "Saiyan",
    "ki": "60.000.000",
    "description": "El protagonista de la serie, un Saiyan criado en la Tierra...",
    "image": "https://dragonball-api.com/characters/goku_normal.webp"
  }
}
```

---

### **⚠️ Personaje Ya Existe**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "Goku"}'
```
**➜ Respuesta 409 Conflict:**
```json
{
  "result": {
    "error": "Character already exists",
    "data": {
      "id": 1,
      "name": "Goku",
      "race": "Saiyan",
      "ki": "60.000.000"
    }
  }
}
```

---

### **🔍 Búsqueda Parcial (Sugerencias)**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "Go"}'
```
**➜ Respuesta 400 Bad Request:**
```json
{
  "result": {
    "error": "No exact match found for 'Go'",
    "suggestions": ["Goku", "Gohan", "Goten", "Gotenks", "Gogeta"]
  }
}
```

---

### **❌ Personaje Inexistente**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": "PersonajeInventado"}'
```
**➜ Respuesta 400 Bad Request:**
```json
{
  "result": {
    "error": "No character found for 'PersonajeInventado'",
    "suggestions": []
  }
}
```

---

### **📝 Validación de Entrada**
```bash
curl -X POST http://localhost:8080/api/v1/characters \
  -H "Content-Type: application/json" \
  -d '{"name": ""}'
```
**➜ Respuesta 400 Bad Request:**
```json
{
  "result": {
    "error": "Name is required and cannot be empty",
    "suggestions": []
  }
}
```

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