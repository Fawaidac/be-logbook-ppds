-- ============================================================
-- Seed data approval (menunggu / pending / Menunggu Validasi) untuk testing
-- ============================================================

-- 1. Tindakan (status: menunggu)
INSERT INTO tindakans (
    mr_number, visit_date, user_username, patient_name, gender, birth_date,
    division, diagnosis_label, procedure_code, plan_procedure, activity,
    procedure_date, room, role, kemandirian, clinical_note,
    supervisor_name, status
) VALUES (
    'RM-APPROVAL-001',
    CURRENT_DATE,
    'residen01',
    'Tn. Budi Santoso Test',
    'L'::gender_enum,
    '1990-01-01'::DATE,
    'Bedah Umum',
    'Katarak Senilis',
    'CAT-001',
    'Fakoemulsifikasi',
    'OK Utama',
    CURRENT_DATE,
    'OK 1',
    'Operator Utama',
    'mandiri'::kemandirian_enum,
    'Tindakan berjalan lancar',
    'dr. Budi Santoso, Sp.B',
    'menunggu'::status_enum
) ON CONFLICT DO NOTHING;

-- Update status jika record RM-SEED-001 sebelumnya berstatus draft
UPDATE tindakans SET status = 'menunggu'::status_enum WHERE mr_number = 'RM-SEED-001';

-- 2. Kegiatan Ilmiah (status: pending)
INSERT INTO kegiatan_ilmiah (
    user_username, program_studi, ppds_name, nim_nip,
    kategori, jenis_kegiatan, topik, tanggal_mulai, tanggal_selesai,
    lokasi_tipe, lokasi_detail, sebagai, pembimbing_1, deskripsi, status
) VALUES (
    'residen01',
    'Ilmu Kesehatan Mata',
    'dr. Ratna Puspita',
    '12345678',
    'simposium',
    'Simposium Nasional Oftalmologi',
    'Perkembangan Bedah Katarak Modern',
    CURRENT_DATE,
    CURRENT_DATE,
    'rsds_fk_unair',
    'RSUD dr. Soebandi',
    'Peserta',
    'dr. Budi Santoso, Sp.B',
    'Deskripsi kegiatan ilmiah testing approval',
    'pending'::status_kegiatan_enum
) ON CONFLICT DO NOTHING;

-- 3. Pendidikan Evaluasi (status: Menunggu Validasi)
INSERT INTO pendidikan_evaluasi (
    user_username, kategori, nama, jenis, stase, evaluator, supervisor, status, tanggal, catatan
) VALUES (
    'residen01',
    'Evaluasi Stase',
    'Mini-CEX Evaluasi Klinik',
    'Mini-CEX',
    'Oftalmologi Umum',
    'dr. Budi Santoso, Sp.B',
    'dr. Budi Santoso, Sp.B',
    'Menunggu Validasi',
    CURRENT_DATE::TEXT,
    'Catatan evaluasi pendidikan testing approval'
) ON CONFLICT DO NOTHING;
