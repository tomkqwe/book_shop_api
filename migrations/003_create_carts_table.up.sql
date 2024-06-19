CREATE  TABLE IF NOT EXISTS carts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),               -- Первичный ключ
    user_id UUID NOT NULL,             -- Идентификатор пользователя (внешний ключ)
    created_at TIMESTAMP NOT NULL,     -- Дата создания корзины
    CONSTRAINT fk_user
        FOREIGN KEY (user_id) REFERENCES users(id)  -- Внешний ключ на таблицу users
);
