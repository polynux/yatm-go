CREATE TABLE IF NOT EXISTS Praticiens (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    firstname TEXT NOT NULL,
    address TEXT NOT NULL,
    zip TEXT NOT NULL,
    city TEXT NOT NULL,
    description TEXT,
    profession TEXT NOT NULL
);
