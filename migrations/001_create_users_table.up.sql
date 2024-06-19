
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),              -- Первичный ключ
    email VARCHAR(255) NOT NULL UNIQUE, -- Email пользователя, уникальный
    password VARCHAR(255) NOT NULL,    -- Пароль (захешированный)
    role VARCHAR(50) NOT NULL          -- Роль пользователя (например, admin, user и т.д.)
);
