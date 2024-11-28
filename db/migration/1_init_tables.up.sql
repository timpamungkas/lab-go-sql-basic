CREATE TABLE IF NOT EXISTS clients (
    client_id TEXT NOT NULL PRIMARY KEY, 
    full_name TEXT NOT NULL, 
    email TEXT, 
    phone TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS client_apartments (
    apartment_id TEXT NOT NULL PRIMARY KEY, 
    description TEXT,
    building_name TEXT,
    room_number TEXT,
    street_address TEXT NOT NULL,
    city TEXT NOT NULL,
    postal_code TEXT,
    is_available_for_rent BOOLEAN,
    rent_price REAL NOT NULL,
    client_id TEXT NOT NULL, 
    FOREIGN KEY (client_id) REFERENCES clients(client_id) ON DELETE CASCADE
);