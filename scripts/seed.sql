-- =====================================
-- LINKME SEED SCRIPT
-- =====================================

-- Clear existing data
TRUNCATE TABLE links RESTART IDENTITY CASCADE;
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

-- Insert test users
INSERT INTO users (email, password, created_at, updated_at)
VALUES
('test1@example.com', '$2a$14$dummyhashvalue1234567890123456789012345678901234567890', NOW(), NOW()),
('test2@example.com', '$2a$14$dummyhashvalue1234567890123456789012345678901234567890', NOW(), NOW());

-- Insert test links
INSERT INTO links (title, url, user_id, created_at, updated_at)
VALUES
('Google', 'https://google.com', 1, NOW(), NOW()),
('GitHub', 'https://github.com', 1, NOW(), NOW()),
('YouTube', 'https://youtube.com', 2, NOW(), NOW());
