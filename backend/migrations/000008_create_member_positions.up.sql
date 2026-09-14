CREATE TABLE member_positions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    member_id UUID NOT NULL,
    position_id UUID NOT NULL,
    department_id UUID,

    start_date DATE,
    end_date DATE,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_member_positions_member
        FOREIGN KEY (member_id)
        REFERENCES organization_members(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_member_positions_position
        FOREIGN KEY (position_id)
        REFERENCES positions(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_member_positions_department
        FOREIGN KEY (department_id)
        REFERENCES departments(id)
        ON DELETE SET NULL,

    CONSTRAINT check_assignment_dates
        CHECK (
            end_date IS NULL
            OR start_date IS NULL
            OR end_date >= start_date
        )
);