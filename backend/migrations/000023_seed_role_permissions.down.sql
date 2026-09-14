DELETE FROM role_permissions
WHERE role_id IN (
    SELECT id
    FROM roles
    WHERE code IN (
        'master',
        'organization',
        'member'
    )
);