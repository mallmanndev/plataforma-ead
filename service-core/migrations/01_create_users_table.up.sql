CREATE TABLE user_types (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP,
    "deletedAt" TIMESTAMP
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(11) NOT NULL,
    password VARCHAR(255) NOT NULL,
    typeId VARCHAR(100) NOT NULL REFERENCES user_types(id),
    "createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updatedAt" TIMESTAMP,
    "deletedAt" TIMESTAMP
);

INSERT INTO user_types (id, name, description)
VALUES ('admin', 'Admin', 'A Administrator user'),
       ('student', 'Student', 'A student user');
