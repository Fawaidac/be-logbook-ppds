CREATE TABLE IF NOT EXISTS pendidikan_cbd (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    pasien VARCHAR(255) NULL,
    topik VARCHAR(255) NOT NULL,
    kategori VARCHAR(255) NULL,
    pembimbing VARCHAR(255) NULL,
    kompleksitas VARCHAR(255) NULL,
    skor VARCHAR(255) NULL,
    status VARCHAR(255) DEFAULT 'Menunggu Validasi',
    tanggal VARCHAR(255) NULL,
    catatan TEXT NULL,
    revisi_catatan TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_pendidikan_cbd_updated_at ON pendidikan_cbd;
CREATE TRIGGER trg_pendidikan_cbd_updated_at
BEFORE UPDATE ON pendidikan_cbd
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();