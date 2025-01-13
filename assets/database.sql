CREATE EXTENSION pgcrypto;

CREATE SCHEMA marisa;
-- User table
-- Metadata may contain address, ssn etc.
CREATE TABLE IF NOT EXISTS marisa.user (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    "name" VARCHAR(255),
    surname VARCHAR(255),
    birth DATE,
    "role" VARCHAR(50), -- e.g., 'user', 'admin', 'mod'
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Parents table
CREATE TABLE IF NOT EXISTS marisa.parent (
    id SERIAL PRIMARY KEY,
    parent_id INT REFERENCES marisa.user(id) ON DELETE CASCADE,
    child_id INT REFERENCES marisa.user(id) ON DELETE CASCADE,
    UNIQUE (parent_id, child_id)
);

-- Groups table
CREATE TABLE IF NOT EXISTS marisa.group (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL, -- helpy, seconda elementare, terza elementare, etc.
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Group member association table
CREATE TABLE IF NOT EXISTS marisa.group_member (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES marisa.user(id) ON DELETE CASCADE,
    group_id INT REFERENCES marisa.group(id) ON DELETE CASCADE,
    "role" VARCHAR(50), -- e.g., 'member', 'staff'
    UNIQUE (user_id, group_id)
);

-- Events table
CREATE TABLE IF NOT EXISTS marisa.event (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS marisa.event_item (
    id SERIAL PRIMARY KEY,
    event_id UUID REFERENCES marisa.event(id) ON DELETE CASCADE,
    "name" VARCHAR(255) NOT NULL, -- e.g., 'deposit', 'final'
    payment_link VARCHAR(255),
    document_template UUID,
    kind VARCHAR(50), -- e.g., 'payment', 'document'
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS marisa.user_item (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES marisa.user(id) ON DELETE CASCADE,
    payment_id VARCHAR(255), -- Stores Stripe's payment intent
    document_id UUID,
    event_item INT REFERENCES marisa.event_item(id) ON DELETE SET NULL,
    amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS marisa.event_group (
    id SERIAL PRIMARY KEY,
    event_id UUID REFERENCES marisa.event(id) ON DELETE CASCADE,
    group_id INT REFERENCES marisa.group(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
