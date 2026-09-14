DELETE FROM permissions
WHERE code IN (
    'organization.read',
    'organization.create',
    'organization.update',
    'organization.delete',

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
    'member.delete',

    'user.read',
    'user.create',
    'user.update',
    'user.delete'
);