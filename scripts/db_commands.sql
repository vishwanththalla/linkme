-- =====================================
-- LINKME DATABASE UTILITY COMMANDS
-- =====================================

-- View all tables
\dt

-- ========================
-- USERS
-- ========================

-- View all users
SELECT id, email, created_at FROM users;

-- Delete all users (dangerous)
-- DELETE FROM users;

-- ========================
-- LINKS
-- ========================

-- View all links
SELECT * FROM links;

-- View links for a specific user
-- Replace 1 with actual user_id
SELECT * FROM links WHERE user_id = 1;

-- Join links with users
SELECT links.id, links.title, links.url, users.email
FROM links
JOIN users ON links.user_id = users.id;

-- Delete a specific link
-- DELETE FROM links WHERE id = 1;

-- Reset links table (clears data + resets ID counter)
-- TRUNCATE TABLE links RESTART IDENTITY;

-- Reset both tables (careful)
-- TRUNCATE TABLE links, users RESTART IDENTITY CASCADE;


psql -d linkme -f scripts/db_commands.sql
