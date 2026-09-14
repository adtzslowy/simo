INSERT INTO roles (
    name,
    code,
    description
)
VALUES
(
    'Master',
    'master',
    'Full access to the entire SIMO system.'
),
(
    'Organization',
    'organization',
    'Manage an assigned organization and its resources.'
),
(
    'Member',
    'member',
    'Basic access for organization members.'
)
ON CONFLICT (code) DO NOTHING;