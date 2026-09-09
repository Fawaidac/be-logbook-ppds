-- Tambah kolom profil ke users (diisi residen lewat halaman Master Data)
ALTER TABLE users
    ADD COLUMN IF NOT EXISTS nik VARCHAR(50) DEFAULT '',
    ADD COLUMN IF NOT EXISTS birth_place VARCHAR(100) DEFAULT '',
    ADD COLUMN IF NOT EXISTS birth_date DATE NULL,
    ADD COLUMN IF NOT EXISTS gender VARCHAR(50) DEFAULT '',
    ADD COLUMN IF NOT EXISTS blood_type VARCHAR(10) DEFAULT '',
    ADD COLUMN IF NOT EXISTS religion VARCHAR(50) DEFAULT '',
    ADD COLUMN IF NOT EXISTS university VARCHAR(150) DEFAULT '',
    ADD COLUMN IF NOT EXISTS citizenship VARCHAR(100) DEFAULT 'Warga Negara Indonesia (WNI)',
    ADD COLUMN IF NOT EXISTS email_personal VARCHAR(150) DEFAULT '',
    ADD COLUMN IF NOT EXISTS phone_mobile VARCHAR(50) DEFAULT '',
    ADD COLUMN IF NOT EXISTS phone_home VARCHAR(50) DEFAULT '',
    ADD COLUMN IF NOT EXISTS address_ktp TEXT DEFAULT '',
    ADD COLUMN IF NOT EXISTS address_domicile TEXT DEFAULT '',
    ADD COLUMN IF NOT EXISTS str_number VARCHAR(100) DEFAULT '',
    ADD COLUMN IF NOT EXISTS str_issued DATE NULL,
    ADD COLUMN IF NOT EXISTS str_expired DATE NULL,
    ADD COLUMN IF NOT EXISTS str_note VARCHAR(150) DEFAULT '',
    ADD COLUMN IF NOT EXISTS profession VARCHAR(150) DEFAULT '',
    ADD COLUMN IF NOT EXISTS competency VARCHAR(150) DEFAULT '',
    ADD COLUMN IF NOT EXISTS college VARCHAR(150) DEFAULT '',
    ADD COLUMN IF NOT EXISTS stage VARCHAR(150) DEFAULT '',
    ADD COLUMN IF NOT EXISTS dpjp VARCHAR(150) DEFAULT '';

-- Riwayat pekerjaan residen (tab Pekerjaan di Master Data)
CREATE TABLE IF NOT EXISTS user_work_histories (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(50) NOT NULL,
    position VARCHAR(150) NOT NULL,
    institution VARCHAR(150) NOT NULL,
    start_date DATE NULL,
    status VARCHAR(50) DEFAULT 'Aktif',
    sip VARCHAR(150) DEFAULT '',
    location VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_work_histories_username ON user_work_histories(user_username);
