CREATE TABLE IF NOT EXISTS tracknme.addresses (
    id UUID PRIMARY KEY,
    employee_id UUID NOT NULL,
    zip_code VARCHAR(9) NOT NULL,
    street VARCHAR(255) NOT NULL,
    complement VARCHAR(255),
    unit VARCHAR(255),
    neighborhood VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(2) NOT NULL,
    state_name VARCHAR(255),
    region VARCHAR(255),
    ibge_code VARCHAR(255),
    gia_code VARCHAR(255),
    area_code VARCHAR(3),
    siafi_code VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
