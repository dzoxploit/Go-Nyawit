CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE estates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    width INTEGER NOT NULL CHECK (width BETWEEN 1 AND 50000),
    length INTEGER NOT NULL CHECK (length BETWEEN 1 AND 50000)
);

CREATE TABLE trees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    estate_id UUID NOT NULL REFERENCES estates(id) ON DELETE CASCADE,
    x INTEGER NOT NULL,
    y INTEGER NOT NULL,
    height INTEGER NOT NULL CHECK (height BETWEEN 1 AND 30),

    UNIQUE(estate_id, x, y)
);