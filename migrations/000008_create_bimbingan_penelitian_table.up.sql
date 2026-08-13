CREATE TABLE IF NOT EXISTS bimbingan_penelitian (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    sesi_ke VARCHAR(50) DEFAULT 'Sesi 1',
    tahap VARCHAR(100) NOT NULL,
    topik_bimbingan VARCHAR(255) NOT NULL,
    tanggal DATE NULL,
    pembimbing VARCHAR(255) NULL,
    catatan_pembimbing TEXT NULL,
    status_acc VARCHAR(50) DEFAULT 'Dalam Proses',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_bimbingan_penelitian_updated_at ON bimbingan_penelitian;
CREATE TRIGGER trg_bimbingan_penelitian_updated_at
BEFORE UPDATE ON bimbingan_penelitian
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();