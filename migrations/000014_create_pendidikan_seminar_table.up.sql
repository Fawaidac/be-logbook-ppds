CREATE TABLE IF NOT EXISTS pendidikan_seminar (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    judul VARCHAR(255) NOT NULL,
    jenis VARCHAR(255) NULL,
    narasumber VARCHAR(255) NULL,
    ruang VARCHAR(255) NULL,
    skor VARCHAR(255) NULL,
    status VARCHAR(255) DEFAULT 'Menunggu Validasi',
    tanggal VARCHAR(255) NULL,
    catatan TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_pendidikan_seminar_updated_at ON pendidikan_seminar;
CREATE TRIGGER trg_pendidikan_seminar_updated_at
BEFORE UPDATE ON pendidikan_seminar
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();