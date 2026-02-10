CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL
);

CREATE TABLE packs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    size INT NOT NULL
);

-- Insert default product and packs
DO $$
DECLARE
    default_product_id UUID;
BEGIN
    INSERT INTO products (id, name)
        VALUES (uuid_generate_v4(), 'Generic Shoes')
        RETURNING id INTO default_product_id;
    INSERT INTO packs (product_id, size) VALUES
        (default_product_id, 250),
        (default_product_id, 500),
        (default_product_id, 1000),
        (default_product_id, 2000),
        (default_product_id, 5000);
END$$;
