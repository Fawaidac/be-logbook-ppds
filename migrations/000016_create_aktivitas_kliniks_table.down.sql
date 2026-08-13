DROP TRIGGER IF EXISTS trg_aktivitas_kliniks_updated_at ON aktivitas_kliniks;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS aktivitas_kliniks;