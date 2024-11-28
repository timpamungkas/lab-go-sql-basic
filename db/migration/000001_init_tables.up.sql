CREATE TABLE IF NOT EXISTS clients (
    client_id TEXT NOT NULL PRIMARY KEY, 
    full_name TEXT NOT NULL, 
    email TEXT, 
    phone TEXT NOT NULL
);