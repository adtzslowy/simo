INSERT INTO permissions (
    name,
    code,
    description
)
VALUES
-- Organization
('View Organizations', 'organization.read', 'View organization data'),
('Create Organizations', 'organization.create', 'Create a new organization'),
('Update Organizations', 'organization.update', 'Update organization data'),
('Delete Organizations', 'organization.delete', 'Delete an organization'),

-- Organization Period
('View Organization Periods', 'period.read', 'View organization periods'),
('Create Organization Periods', 'period.create', 'Create an organization period'),
('Update Organization Periods', 'period.update', 'Update an organization period'),
('Delete Organization Periods', 'period.delete', 'Delete an organization period'),

-- Department
('View Departments', 'department.read', 'View organization departments'),
('Create Departments', 'department.create', 'Create a department'),
('Update Departments', 'department.update', 'Update a department'),
('Delete Departments', 'department.delete', 'Delete a department'),

-- Position
('View Positions', 'position.read', 'View organization positions'),
('Create Positions', 'position.create', 'Create a position'),
('Update Positions', 'position.update', 'Update a position'),
('Delete Positions', 'position.delete', 'Delete a position'),

-- Member
('View Members', 'member.read', 'View organization members'),
('Create Members', 'member.create', 'Add a member to an organization'),
('Update Members', 'member.update', 'Update organization member data'),
('Delete Members', 'member.delete', 'Remove a member from an organization'),

-- User
('View Users', 'user.read', 'View system users'),
('Create Users', 'user.create', 'Create a system user'),
('Update Users', 'user.update', 'Update system user data'),
('Delete Users', 'user.delete', 'Delete a system user')
ON CONFLICT (code) DO NOTHING;