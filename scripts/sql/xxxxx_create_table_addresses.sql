CREATE TABLE addresses ( 
    id TEXT PRIMARY KEY,
    street TEXT,
    locality TEXT,
    region TEXT,
    postal_code TEXT,
    country TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);