CREATE TABLE IF NOT EXISTS pendidikan_kompetensi (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    kode VARCHAR(255) NOT NULL,
    nama VARCHAR(255) NOT NULL,
    domain VARCHAR(255) NOT NULL,
    level_target VARCHAR(255) NOT NULL,
    target_log INT DEFAULT 10,
    achieved_log INT DEFAULT 0,
    evaluator VARCHAR(255) NULL,
    status VARCHAR(255) DEFAULT 'Dalam Proses',
    tgl_verifikasi VARCHAR(255) NULL,
    deskripsi TEXT NULL,
    revisi_catatan TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_pendidikan_kompetensi_updated_at ON pendidikan_kompetensi;
CREATE TRIGGER trg_pendidikan_kompetensi_updated_at
BEFORE UPDATE ON pendidikan_kompetensi
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();