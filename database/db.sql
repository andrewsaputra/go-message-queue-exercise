-- Run with PSQL as postgres / superuser :

DROP DATABASE IF EXISTS demo_message_queue;
DROP USER IF EXISTS demouser;

CREATE DATABASE demo_message_queue;
\c demo_message_queue;
CREATE USER demouser WITH PASSWORD 'password';

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO demouser;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO demouser;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO demouser;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO demouser;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL DEFAULT 0
);

INSERT INTO users (name, email, created_at)
VALUES 
    ('John Smith', 'john.smith@email.com', EXTRACT(EPOCH FROM NOW())*1000),
    ('Laura Jane', 'laura.jane@email.com', EXTRACT(EPOCH FROM NOW())*1000);


CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    images TEXT[] NOT NULL,
    compressed_images TEXT[] NOT NULL DEFAULT '{}',
    price NUMERIC(10,2) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL DEFAULT 0,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

/*
INSERT INTO products (user_id, name, description, images, compressed_images, price, created_at)
VALUES
    (2, 'Product 1', 'This is product 1', '{"image1.jpg", "image2.jpg"}', '{}', 9.99, EXTRACT(EPOCH FROM NOW())*1000);
*/