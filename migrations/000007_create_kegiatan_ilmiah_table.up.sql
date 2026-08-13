-- ENUM Khusus Kegiatan Ilmiah
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kategori_kegiatan_enum') THEN
        CREATE TYPE kategori_kegiatan_enum AS ENUM ('simposium', 'workshop', 'multidisiplin', 'ilmiah_lain');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'lokasi_tipe_enum') THEN
        CREATE TYPE lokasi_tipe_enum AS ENUM ('rsds_fk_unair', 'luar_rsds_fk_unair');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_kegiatan_enum') THEN
        CREATE TYPE status_kegiatan_enum AS ENUM ('pending', 'disetujui', 'ditolak');
    END IF;
END $$;

-- Tabel Utama Kegiatan Ilmiah (Input Teks Nama Dosen Pembimbing & Penguji)
CREATE TABLE IF NOT EXISTS kegiatan_ilmiah (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) NULL,
    program_studi VARCHAR(100) NULL,
    ppds_name VARCHAR(100) NULL,
    nim_nip VARCHAR(100) NULL,
    
    kategori kategori_kegiatan_enum NOT NULL,

    -- Detail Kegiatan
    jenis_kegiatan VARCHAR(255) NOT NULL,
    topik VARCHAR(255) NOT NULL,
    tanggal_mulai DATE NULL,
    tanggal_selesai DATE NULL,
    lokasi_tipe lokasi_tipe_enum DEFAULT 'rsds_fk_unair',
    lokasi_detail VARCHAR(255) NULL,
    sebagai VARCHAR(100) NULL,

    -- Nama Pembimbing (Input Teks Biasa, maks 5)
    pembimbing_1 VARCHAR(255) NULL,
    pembimbing_2 VARCHAR(255) NULL,
    pembimbing_3 VARCHAR(255) NULL,
    pembimbing_4 VARCHAR(255) NULL,
    pembimbing_5 VARCHAR(255) NULL,

    -- Nama Penguji (Input Teks Biasa, maks 5)
    penguji_1 VARCHAR(255) NULL,
    penguji_2 VARCHAR(255) NULL,
    penguji_3 VARCHAR(255) NULL,
    penguji_4 VARCHAR(255) NULL,
    penguji_5 VARCHAR(255) NULL,

    deskripsi TEXT NULL,
    lampiran_path VARCHAR(255) NULL,
    status status_kegiatan_enum DEFAULT 'pending',
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
DROP TRIGGER IF EXISTS trg_kegiatan_ilmiah_updated_at ON kegiatan_ilmiah;
CREATE TRIGGER trg_kegiatan_ilmiah_updated_at
BEFORE UPDATE ON kegiatan_ilmiah
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();