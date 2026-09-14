INSERT INTO organizations (
    organization_type_id,
    institution_id,
    academic_department_id,
    name,
    code,
    description,
    is_active
)
SELECT
    ot.id,
    i.id,
    NULL,
    'BEM POLITAP',
    'BEM-POLITAP',
    'Badan Eksekutif Mahasiswa Politeknik Negeri Ketapang',
    TRUE
FROM organization_types ot
CROSS JOIN institutions i
WHERE ot.code = 'BEM'
  AND i.code = 'POLITAP'
  AND NOT EXISTS (
      SELECT 1
      FROM organizations o
      WHERE o.code = 'BEM-POLITAP'
  );

INSERT INTO organizations (
    organization_type_id,
    institution_id,
    academic_department_id,
    name,
    code,
    description,
    is_active
)
SELECT
    ot.id,
    i.id,
    ad.id,
    'HMJ ' || ad.name,
    'HMJ-' || ad.code,
    'Himpunan Mahasiswa Jurusan ' || ad.name,
    TRUE
FROM organization_types ot
CROSS JOIN institutions i
JOIN academic_departments ad
    ON ad.institution_id = i.id
WHERE ot.code = 'HMJ'
  AND i.code = 'POLITAP'
  AND NOT EXISTS (
      SELECT 1
      FROM organizations o
      WHERE o.code = 'HMJ-' || ad.code
  );