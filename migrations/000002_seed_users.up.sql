INSERT INTO users (username, name, email, password, role, jabatan) VALUES
('residen01', 'dr. Ratna Puspita', 'residen01@ppds.id', '$2a$10$qKelQ9ZR9gAO2TpamvZj6uZNDYl99wnZL9aSZX/E6snosCDoGAyNW', 'residen', 'Residen Bedah Th.2'),
('supervisor01', 'dr. Budi Santoso, Sp.B', 'supervisor01@ppds.id', '$2a$10$/TMOxcKcXg/wCi3TPGC5/O5DbNH87yVDbwX3gv5UKYJhgE/11boxO', 'supervisor', 'DPJP Bedah Umum'),
('admin01', 'Siti Amalia, S.KM', 'admin01@ppds.id', '$2a$10$r34Z/hlp7ajPXp6FEl25UugXxfdlDoywpRisxjotyHFq5xwuL0sNm', 'admin', 'Admin Komkordik')
ON CONFLICT (username) DO UPDATE SET 
    name = EXCLUDED.name,
    email = EXCLUDED.email,
    password = EXCLUDED.password, 
    role = EXCLUDED.role,
    jabatan = EXCLUDED.jabatan,
    created_at = CURRENT_TIMESTAMP;