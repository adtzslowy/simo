DELETE FROM user_roles
WHERE user_id = (
    SELECT id
    FROM users
    WHERE email = 'aditya.prasetyooo25@gmail.com'
)
AND role_id = (
    SELECT id
    FROM roles
    WHERE code = 'master'
)
AND organization_id IS NULL;

DELETE FROM users
WHERE email = 'aditya.prasetyooo25@gmail.com';