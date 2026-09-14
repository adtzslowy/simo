-- MASTER
-- Master mendapatkan seluruh permission.

INSERT INTO role_permissions (role_id, permission_id)
SELECT
    r.id,
    p.id
FROM roles r
CROSS JOIN permissions p
WHERE r.code = 'master'
  AND NOT EXISTS (
      SELECT 1
      FROM role_permissions rp
      WHERE rp.role_id = r.id
        AND rp.permission_id = p.id
  );


-- ORGANIZATION
-- Organization hanya mendapatkan permission yang berhubungan
-- dengan pengelolaan organization.

INSERT INTO role_permissions (role_id, permission_id)
SELECT
    r.id,
    p.id
FROM roles r
CROSS JOIN permissions p
WHERE r.code = 'organization'
  AND p.code IN (
      'organization.read',
      'organization.update',

      'period.read',
      'period.create',
      'period.update',
      'period.delete',

      'department.read',
      'department.create',
      'department.update',
      'department.delete',

      'position.read',
      'position.create',
      'position.update',
      'position.delete',

      'member.read',
      'member.create',
      'member.update',
      'member.delete'
  )
  AND NOT EXISTS (
      SELECT 1
      FROM role_permissions rp
      WHERE rp.role_id = r.id
        AND rp.permission_id = p.id
  );


-- MEMBER
-- Member hanya mendapatkan akses membaca data organisasi.

INSERT INTO role_permissions (role_id, permission_id)
SELECT
    r.id,
    p.id
FROM roles r
CROSS JOIN permissions p
WHERE r.code = 'member'
  AND p.code IN (
      'organization.read',
      'period.read',
      'department.read',
      'position.read',
      'member.read'
  )
  AND NOT EXISTS (
      SELECT 1
      FROM role_permissions rp
      WHERE rp.role_id = r.id
        AND rp.permission_id = p.id
  );