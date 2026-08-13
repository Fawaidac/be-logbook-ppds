CREATE TABLE IF NOT EXISTS pendidikan_rotasi (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    stase VARCHAR(255) NOT NULL,
    lokasi VARCHAR(255) NULL,
    periode VARCHAR(255) NULL,
    pembimbing VARCHAR(255) NULL,
    kehadiran VARCHAR(255) DEFAULT '100%',
    nilai VARCHAR(255) NULL,
    status VARCHAR(255) DEFAULT 'Berlangsung',
    tanggal VARCHAR(255) NULL,
    catatan TEXT NULL,
    revisi_catatan TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_pendidikan_rotasi_updated_at ON pendidikan_rotasi;
CREATE TRIGGER trg_pendidikan_rotasi_updated_at
BEFORE UPDATE ON pendidikan_rotasi
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();