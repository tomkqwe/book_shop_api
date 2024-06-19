CREATE TABLE IF NOT EXISTS  books (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),              -- Первичный ключ
    title VARCHAR(255) NOT NULL,       -- Название книги
    author VARCHAR(255) NOT NULL,      -- Автор книги
    year_published INTEGER NOT NULL,      -- Год публикации
    price DECIMAL(10, 2) NOT NULL,     -- Цена книги (до 10 цифр, 2 из которых после запятой)
    category VARCHAR(100) NOT NULL     -- Категория книги
);
