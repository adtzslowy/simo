INSERT INTO institutions (
    name,
    code,
    description
)
VALUES (
    'Politeknik Negeri Ketapang',
    'POLITAP',
    'Perguruan Tinggi Negeri Vokasi di Kabupaten Ketapang, Kalimantan Barat.'
);


INSERT INTO academic_departments (
    institution_id,
    name,
    code
)
SELECT
    id,
    'Jurusan Teknik Sipil dan Pertambangan',
    'JTSP'
FROM institutions
WHERE code = 'POLITAP';

INSERT INTO academic_departments (
    institution_id,
    name,
    code
)
SELECT
    id,
    'Jurusan Teknik Mesin',
    'JTM'
FROM institutions
WHERE code = 'POLITAP';

INSERT INTO academic_departments (
    institution_id,
    name,
    code
)
SELECT
    id,
    'Jurusan Pertanian dan Bisnis',
    'JPB'
FROM institutions
WHERE code = 'POLITAP';

INSERT INTO academic_departments (
    institution_id,
    name,
    code
)
SELECT
    id,
    'Jurusan Teknik Elektro dan Teknik Informatika',
    'JTEI'
FROM institutions
WHERE code = 'POLITAP';

-- Teknik Sipil dan Pertambangan

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Teknologi Rekayasa Konstruksi Jalan dan Jembatan',
    'TRKJJ',
    'D4'
FROM academic_departments
WHERE code = 'JTSP';

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Teknologi Pertambangan',
    'TP',
    'D3'
FROM academic_departments
WHERE code = 'JTSP';


-- Teknik Mesin

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Pemeliharaan Mesin',
    'PM',
    'D3'
FROM academic_departments
WHERE code = 'JTM';


-- Pertanian dan Bisnis

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Teknologi Produksi Tanaman Perkebunan',
    'TPTP',
    'D4'
FROM academic_departments
WHERE code = 'JPB';

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Manajemen Agribisnis',
    'MA',
    'D4'
FROM academic_departments
WHERE code = 'JPB';

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Teknologi Hasil Perkebunan',
    'THP',
    'D3'
FROM academic_departments
WHERE code = 'JPB';

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Agroindustri',
    'AGI',
    'D3'
FROM academic_departments
WHERE code = 'JPB';


-- Teknik Elektro dan Teknik Informatika

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Teknologi Listrik',
    'TL',
    'D3'
FROM academic_departments
WHERE code = 'JTEI';

INSERT INTO academic_programs (
    academic_department_id,
    name,
    code,
    degree
)
SELECT
    id,
    'Teknologi Informasi',
    'TI',
    'D3'
FROM academic_departments
WHERE code = 'JTEI';

