CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),   -- Уникальный идентификатор элемента заказа
    order_id UUID NOT NULL,                           -- Идентификатор заказа (внешний ключ)
    book_id UUID NOT NULL,                            -- Идентификатор книги (внешний ключ)
    quantity INT NOT NULL CHECK (quantity > 0),       -- Количество товара
    price DECIMAL(10, 2) NOT NULL,                    -- Цена товара (например, 10 цифр, из них 2 после запятой)

    CONSTRAINT fk_order
        FOREIGN KEY (order_id) REFERENCES orders(id),  -- Внешний ключ на таблицу заказов
    CONSTRAINT fk_book
        FOREIGN KEY (book_id) REFERENCES books(id)     -- Внешний ключ на таблицу книг
);
