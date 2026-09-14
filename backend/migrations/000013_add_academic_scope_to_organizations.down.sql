ALTER TABLE organizations
DROP CONSTRAINT IF EXISTS fk_organizations_academic_department;

ALTER TABLE organizations
DROP CONSTRAINT IF EXISTS fk_organizations_institution;

ALTER TABLE organizations
DROP COLUMN IF EXISTS academic_department_id;

ALTER TABLE organizations
DROP COLUMN IF EXISTS institution_id;