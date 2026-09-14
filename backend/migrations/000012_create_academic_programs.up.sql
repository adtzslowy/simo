CREATE TABLE academic_programs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    academic_department_id UUID NOT NULL,

    name VARCHAR(150) NOT NULL,
    code VARCHAR(50) NOT NULL,

    degree VARCHAR(20),

    description TEXT,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_academic_programs_department
        FOREIGN KEY (academic_department_id)
        REFERENCES academic_departments(id)
        ON DELETE CASCADE,

    CONSTRAINT unique_academic_program_code
        UNIQUE (academic_department_id, code)
);