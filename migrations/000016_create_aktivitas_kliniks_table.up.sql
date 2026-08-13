DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'enum_jenis_kelamin') THEN
        CREATE TYPE enum_jenis_kelamin AS ENUM ('L', 'P');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'enum_kemandirian') THEN
        CREATE TYPE enum_kemandirian AS ENUM ('Mandiri', 'Dibimbing', 'Observasi');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'enum_status_aktivitas') THEN
        CREATE TYPE enum_status_aktivitas AS ENUM ('draft', 'Menunggu Validasi', 'Disetujui', 'Perlu Revisi');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS aktivitas_kliniks (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    kategori VARCHAR(255) NOT NULL,
    log_id VARCHAR(255) NOT NULL UNIQUE,
    rm VARCHAR(255) NULL,
    nama_pasien VARCHAR(255) NULL,
    jenis_kelamin enum_jenis_kelamin NULL,
    umur VARCHAR(255) NULL,
    tanggal DATE NULL,
    diagnosis VARCHAR(255) NULL,
    prosedur TEXT NULL,
    tindakan VARCHAR(255) NULL,
    tindakan_key VARCHAR(255) NULL,
    supervisor VARCHAR(255) NULL,
    kemandirian enum_kemandirian DEFAULT 'Dibimbing',
    status enum_status_aktivitas DEFAULT 'Menunggu Validasi',
    catatan TEXT NULL,
    revisi_catatan TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TRIGGER IF EXISTS trg_aktivitas_kliniks_updated_at ON aktivitas_kliniks;
CREATE TRIGGER trg_aktivitas_kliniks_updated_at
BEFORE UPDATE ON aktivitas_kliniks
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();