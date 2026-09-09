-- ============================================================
-- Seed 1 data tindakan untuk testing (status: menunggu)
-- Idempotent: hanya insert jika mr_number belum ada
-- ============================================================
INSERT INTO tindakans (
    mr_number, visit_date, user_username, patient_name, gender, birth_date,
    division, diagnosis_label, procedure_code, plan_procedure, activity,
    procedure_date, room, role, kemandirian, clinical_note,
    supervisor_name, status, feedback
)
SELECT
    'RM-SEED-001',
    CURRENT_DATE - 1,
    COALESCE((SELECT name FROM users WHERE username = 'residen01'), 'dr. Ratna Puspita'),
    'Tn. Ahmad Hidayat',
    'L'::gender_enum,
    '1992-05-15'::DATE,
    'Bedah Umum',
    'Apendisitis Akut',
    'APX-001',
    'Apendektomi',
    'Bedah Sentral (OK)',
    CURRENT_DATE,
    'OK 1 Central',
    'Operator Utama',
    'mandiri'::kemandirian_enum,
    'Apendektomi elektif, pasien stabil, kondisi post-op baik',
    COALESCE((SELECT name FROM users WHERE username = 'supervisor01'), 'dr. Budi Santoso, Sp.B'),
    'draft'::status_enum,
    NULL
WHERE NOT EXISTS (
    SELECT 1 FROM tindakans WHERE mr_number = 'RM-SEED-001'
);
