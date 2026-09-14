DELETE FROM academic_programs
WHERE code IN (
    'TRKJJ',
    'TP',
    'PM',
    'TPTP',
    'MA',
    'THP',
    'AGI',
    'TL',
    'TI'
);

DELETE FROM academic_departments
WHERE code IN (
    'JTSP',
    'JTM',
    'JPB',
    'JTEI'
);

DELETE FROM institutions
WHERE code = 'POLITAP';