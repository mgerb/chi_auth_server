CREATE TABLE users(
    id BIGSERIAL NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    date_created timestamp DEFAULT NOW(),
    
    PRIMARY KEY(id)
)
