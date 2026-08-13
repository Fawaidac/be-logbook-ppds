DROP TRIGGER IF EXISTS trg_kegiatan_ilmiah_updated_at ON kegiatan_ilmiah;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS kegiatan_ilmiah;