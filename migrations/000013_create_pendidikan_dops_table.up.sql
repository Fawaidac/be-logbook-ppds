CREATE TABLE IF NOT EXISTS pendidikan_dops (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    prosedur VARCHAR(255) NULL,
    kategori VARCHAR(255) NULL,
    kesulitan VARCHAR(255) NULL,
    supervisor VARCHAR(255) NULL,
    skor VARCHAR(255) NULL,
    status VARCHAR(255) DEFAULT 'Menunggu Validasi',
    tanggal VARCHAR(255) NULL,
    catatan TEXT NULL,
    revisi_catatan TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_pendidikan_dops_updated_at ON pendidikan_dops;
CREATE TRIGGER trg_pendidikan_dops_updated_at
BEFORE UPDATE ON pendidikan_dops
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();