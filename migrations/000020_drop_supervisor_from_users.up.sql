-- Hapus kolom supervisor_name dari users karena tidak lagi digunakan.
ALTER TABLE users DROP COLUMN IF EXISTS supervisor_name;
