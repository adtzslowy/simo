DELETE FROM roles
WHERE code IN (
    'master',
    'organization',
    'member'
);