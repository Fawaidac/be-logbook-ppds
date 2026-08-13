CREATE TABLE IF NOT EXISTS pendidikan_evaluasi (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    kategori VARCHAR(255) NOT NULL,
    kode VARCHAR(255) NULL,
    nama VARCHAR(255) NULL,
    domain VARCHAR(255) NULL,
    stase VARCHAR(255) NULL,
    lokasi VARCHAR(255) NULL,
    periode VARCHAR(255) NULL,
    pasien VARCHAR(255) NULL,
    fokus VARCHAR(255) NULL,
    kasus VARCHAR(255) NULL,
    prosedur VARCHAR(255) NULL,
    kesulitan VARCHAR(255) NULL,
    judul VARCHAR(255) NULL,
    jenis VARCHAR(255) NULL,
    topik VARCHAR(255) NULL,
    narasumber VARCHAR(255) NULL,
    evaluator VARCHAR(255) NULL,
    pembimbing VARCHAR(255) NULL,
    supervisor VARCHAR(255) NULL,
    ruang VARCHAR(255) NULL,
    kompleksitas VARCHAR(255) NULL,
    level_target VARCHAR(255) NULL,
    target_log INTEGER DEFAULT 10 NULL,
    achieved_log INTEGER DEFAULT 0 NULL,
    kehadiran VARCHAR(255) NULL,
    skor VARCHAR(255) NULL,
    nilai VARCHAR(255) NULL,
    status VARCHAR(255) DEFAULT 'Menunggu Validasi',
    tanggal VARCHAR(255) NULL,
    tgl_verifikasi VARCHAR(255) NULL,
    catatan TEXT NULL,
    deskripsi TEXT NULL,
    revisi_catatan TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Function untuk Auto-Update Timestamp (Dibuat jika belum ada)
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_pendidikan_evaluasi_updated_at ON pendidikan_evaluasi;
CREATE TRIGGER trg_pendidikan_evaluasi_updated_at
BEFORE UPDATE ON pendidikan_evaluasi
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();