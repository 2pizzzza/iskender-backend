CREATE TABLE IF NOT EXISTS Language (
    code VARCHAR(10) PRIMARY KEY,
    name VARCHAR(10)
);


CREATE TABLE IF NOT EXISTS Users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NOT,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Category (
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS CategoryTranslation (
    category_id INT REFERENCES Category(id),
    language_code VARCHAR(10) REFERENCES Language(code),
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (category_id, language_code)
);

CREATE TABLE IF NOT EXISTS Product(
    id SERIAL PRIMARY KEY,C
    type_product VARCHAR(20) CHECK (type IN('item', 'collection')) NOT NULL,
    ref_id INT NOT NULL,
    category_id REFERENCES Category(id),
    name TEXT NOT NULL,
    price DECIMAL NOT NULL DEFAULT 0,
    is_producer BOOL NOT NULL DEFAULT FALSE,
    is_painted BOOL NOT NULL DEFAULT FALSE,
    is_popular BOOL NOT NULL DEFAULT FALSE,
    is_new BOOL NOT NULL DEFAULT FALSE,
    is_garant BOLL NOT NULL DEFAULT FALSE,
    isAqua BOOL NOT NULL DEFAULT FALSE
);


CREATE TABLE IF NOT EXISTS ProductTranslation (
    product_id INT REFERENCES Product(id),
    language_code VARCHAR(10) REFERENCES Language(code),
    description TEXT NOT NULL,
    PRIMARY KEY (product_id, language_code)
);


CREATE TABLE IF NOT EXISTS Photo (
    id SERIAL PRIMARY KEY,
    url VARCHAR(255),
    is_main BOOL DEFAULT FALSE,
    hash_color VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS ProductPhoto(
    product_id INT REFERENCES Product(id),
    photo_id INT REFERENCES Photo(id),
    PRIMARY KEY (product_id, photo_id)
);


CREATE TABLE IF NOT EXISTS Brand (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Review (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    text TEXT NOT NULL,
    is_show BOLL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS Discount (
    id SERIAL PRIMARY KEY, 
    product_id INT REFERENCES Product(id) ON DELETE CASCADE,
    discount_percentage DECIMAL(5, 2) NOT NULL CHECK (discount_percentage >= 0 AND discount_percentage <= 100),
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS Vacancy(
    id SERIAL PRIMARY KEY,
    isActive BOOLEAN DEFAULT true,
    salary DECIMAL NOT NULL
);

CREATE TABLE IF NOT EXISTS VacancyTranslation (
    vacancy_id INT REFERENCES Vacancy(id),
    language_code VARCHAR(10) REFERENCES Language(code),
    title VARCHAR(255) NOT NULL,
    requirements TEXT[] NOT NULL,
    responsibilities TEXT[] NOT NULL,
    conditions TEXT[] NOT NULL,
    information TEXT[] NOT NULL
);


