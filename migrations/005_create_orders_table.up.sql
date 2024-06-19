CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- ID заказа (уникальный идентификатор)
    user_id UUID NOT NULL,                           -- Идентификатор пользователя (внешний ключ)
    order_date TIMESTAMP NOT NULL,                   -- Дата заказа
    total_amount DECIMAL(10, 2) NOT NULL,            -- Общая сумма заказа (например, 10 цифр, из них 2 после запятой)

    CONSTRAINT fk_user
        FOREIGN KEY (user_id) REFERENCES users(id)  -- Внешний ключ на таблицу пользователей
);