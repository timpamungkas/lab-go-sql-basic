INSERT INTO clients (client_id, full_name, email, phone) 
VALUES ('475ab972-6aa8-4964-9baa-1ea1b35b5b4e', 'Alfred Pennyworth', 'alfred.pennyworth@something.com', '5551234') 
ON CONFLICT(client_id) DO NOTHING;

INSERT INTO clients (client_id, full_name, email, phone) 
VALUES ('a1d317f1-45e3-4655-8cf8-3c61fe988bfe', 'Bruce Wayne', 'bruce.wayne@something.com', '5555678') 
ON CONFLICT(client_id) DO NOTHING;

INSERT INTO clients (client_id, full_name, email, phone) 
VALUES ('0d1549b4-7edc-45fa-a019-35bbb9e39de5', 'Clark Kent', 'clark.kent@something.com', '5558765') 
ON CONFLICT(client_id) DO NOTHING;

INSERT INTO client_apartments (
    apartment_id, 
    description, 
    building_name, 
    room_number, 
    street_address, 
    city, 
    postal_code, 
    is_available_for_rent, 
    rent_price, 
    client_id
) VALUES (
    '5576b540-1f61-4876-8581-8ab6e7e65eb6',
    'Spacious 2-bedroom apartment with balcony',
    'Wayne Suites',
    '101',
    '111 Gotham Boulevard',
    'Gotham',
    '12345',
    true,
    7000,
    '475ab972-6aa8-4964-9baa-1ea1b35b5b4e'
) ON CONFLICT(apartment_id) DO NOTHING;

INSERT INTO client_apartments (
    apartment_id, 
    description, 
    building_name, 
    room_number, 
    street_address, 
    city, 
    postal_code, 
    is_available_for_rent, 
    rent_price, 
    client_id
) VALUES (
    '4eaded3d-7a06-4f72-9538-2cdb2169c4be',
    'Modern penthouse near downtown',
    'Wayne Suites',
    '3700',
    '111 Gotham Boulevard',
    'Gotham',
    '12345',
    true,
    65000,
    'a1d317f1-45e3-4655-8cf8-3c61fe988bfe'
) ON CONFLICT(apartment_id) DO NOTHING;

INSERT INTO client_apartments (
    apartment_id, 
    description, 
    building_name, 
    room_number, 
    street_address, 
    city, 
    postal_code, 
    is_available_for_rent, 
    rent_price, 
    client_id
) VALUES (
    'e048c3f4-1bb0-4b46-b9bf-e3397ff7996b',
    'Cozy 2-bedroom apartment in a quiet neighborhood',
    'Sunset Villas',
    '1405',
    '789 Pine Avenue',
    'Metropolis',
    '98765',
    true,
    4000,
    '0d1549b4-7edc-45fa-a019-35bbb9e39de5'
) ON CONFLICT(apartment_id) DO NOTHING;
