CREATE TABLE IF NOT EXISTS cart_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),               -- Первичный ключ
    cart_id UUID NOT NULL,             -- Идентификатор корзины (внешний ключ)
    book_id UUID NOT NULL,             -- Идентификатор книги (внешний ключ)
    count INT NOT NULL CHECK (count > 0),  -- Количество книг, обязательно больше 0
    CONSTRAINT fk_cart
        FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,  -- Внешний ключ на таблицу carts
    CONSTRAINT fk_book
        FOREIGN KEY (book_id) REFERENCES books(id)  -- Внешний ключ на таблицу books
);
