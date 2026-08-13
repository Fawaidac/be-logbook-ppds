DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
        CREATE TYPE gender_enum AS ENUM ('L', 'P');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'kemandirian_enum') THEN
        CREATE TYPE kemandirian_enum AS ENUM ('mandiri', 'dibimbing', 'observasi');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
        CREATE TYPE status_enum AS ENUM ('draft', 'menunggu', 'disetujui', 'ditolak');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS tindakans (
    id SERIAL PRIMARY KEY,
    mr_number VARCHAR(255) NULL, -- Dibuat NULL untuk mengakomodasi data SIMRS tanpa No. RM
    visit_date DATE NULL,
    user_username VARCHAR(255) NULL,
    patient_name VARCHAR(255) NULL,
    gender gender_enum NULL,
    birth_date DATE NULL,
    division VARCHAR(255) NULL,
    diagnosis_label VARCHAR(255) NULL,
    
    -- Detail Tindakan
    procedure_code VARCHAR(255) NULL,
    plan_procedure VARCHAR(255) NOT NULL,
    activity VARCHAR(255) NULL,
    procedure_date DATE NULL,
    room VARCHAR(255) NULL,
    role VARCHAR(255) NULL,
    kemandirian kemandirian_enum DEFAULT 'dibimbing',
    clinical_note TEXT NULL,
    
    supervisor_name VARCHAR(255) NULL,
    status status_enum DEFAULT 'menunggu',
    feedback TEXT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trigger Auto-Update Timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_tindakans_updated_at ON tindakans;
CREATE TRIGGER trg_tindakans_updated_at
BEFORE UPDATE ON tindakans
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();