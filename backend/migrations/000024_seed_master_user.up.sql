INSERT INTO users (
    name,
    email,
    password_hash,
    is_active
)
VALUES (
    'Aditya Prasetyo',
    'aditya.prasetyooo25@gmail.com',
    '$2a$10$11T00wLxap8Ai6ZH7YKjx.de0.SJaemedBk66NvRFdrANYwlh4GS2',
    TRUE
)
ON CONFLICT (email) DO NOTHING;

INSERT INTO user_roles (
    user_id,
    role_id,
    organization_id
)
SELECT
    u.id,
    r.id,
    NULL
FROM users u
CROSS JOIN roles r
WHERE u.email = 'aditya.prasetyooo25@gmail.com'
  AND r.code = 'master'
  AND NOT EXISTS (
      SELECT 1
      FROM user_roles ur
      WHERE ur.user_id = u.id
        AND ur.role_id = r.id
        AND ur.organization_id IS NULL
  );