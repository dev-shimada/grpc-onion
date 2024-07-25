CREATE TABLE entries (
    id VARCHAR(32) PRIMARY KEY,
    user VARCHAR(32),
    status VARCHAR(32),
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO entries (id, user, status, created_at, deleted_at, updated_at) 
VALUES 
('1', 'User1', 'active', '2022-01-01 00:00:00', null, '2022-01-01 00:00:00'),
('2', 'User2', 'deleted', '2022-01-02 00:00:00', '2022-01-02 00:00:00', '2022-01-01 00:00:00'),
('3', 'User3', 'active', '2022-01-03 00:00:00', null, '2022-01-01 00:00:00');
