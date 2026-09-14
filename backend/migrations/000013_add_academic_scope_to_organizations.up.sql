ALTER TABLE organizations
ADD COLUMN institution_id UUID,
ADD COLUMN academic_department_id UUID;

ALTER TABLE organizations
ADD CONSTRAINT fk_organizations_institution
    FOREIGN KEY (institution_id)
    REFERENCES institutions(id)
    ON DELETE CASCADE;

ALTER TABLE organizations
ADD CONSTRAINT fk_organizations_academic_department
    FOREIGN KEY (academic_department_id)
    REFERENCES academic_departments(id)
    ON DELETE SET NULL;