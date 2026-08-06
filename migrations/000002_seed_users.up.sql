INSERT INTO users (username, name, email, password, role, jabatan) VALUES
('residen01', 'dr. Ratna Puspita', 'residen01@ppds.id', 'residen01', 'residen', 'Residen Bedah Th.2'),
('supervisor01', 'dr. Budi Santoso, Sp.B', 'supervisor01@ppds.id', 'supervisor01', 'supervisor', 'DPJP Bedah Umum'),
('admin01', 'Siti Amalia, S.KM', 'admin01@ppds.id', 'admin01', 'admin', 'Admin Komkordik'),
('fawaid', 'Achmad Fawaid, S.Tr.Kom', 'fawaid@ppds.id', 'fawaid', 'superadmin', 'Super Admin')
ON CONFLICT (username) DO UPDATE SET password = EXCLUDED.password, role = EXCLUDED.role;
