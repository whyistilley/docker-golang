-- CreateTable
CREATE TABLE IF NOT EXISTS users (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "email" TEXT NOT NULL,

    PRIMARY KEY ("id")
);

-- Seed
INSERT INTO users (id, name, email)
VALUES
    ('000', 'Alice', 'alice@example.com'),
    ('001', 'Bob', 'bob@example.com'),
    ('002', 'Charlie', 'charlie@example.com');