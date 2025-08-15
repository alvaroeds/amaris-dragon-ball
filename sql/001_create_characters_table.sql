-- =====================================================
-- Dragon Ball API - Schema inicial
-- =====================================================

-- Crear tabla de personajes de Dragon Ball (solo campos esenciales)
CREATE TABLE IF NOT EXISTS characters (
    id SERIAL PRIMARY KEY,
    external_id INTEGER UNIQUE,
    name VARCHAR(100) NOT NULL UNIQUE,
    race VARCHAR(50),                    -- Dato básico
    ki VARCHAR(50),                      -- Campo adicional 1
    description TEXT,                    -- Campo adicional 2
    image VARCHAR(500),                  -- Campo adicional 3
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
