-- Relasi 1 residen -> 1 supervisor:
-- supervisor_name menyimpan nama supervisor (users.name) pembimbing residen.
ALTER TABLE users ADD COLUMN IF NOT EXISTS supervisor_name VARCHAR(255) NULL;
