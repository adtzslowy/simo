CREATE TABLE organization_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    organization_id UUID NOT NULL,
    user_id UUID NOT NULL,
    period_id UUID,

    status VARCHAR(30) NOT NULL DEFAULT 'active',

    joined_at DATE,
    left_at DATE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_members_organization
        FOREIGN KEY (organization_id)
        REFERENCES organizations(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_members_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_members_period
        FOREIGN KEY (period_id)
        REFERENCES organization_periods(id)
        ON DELETE SET NULL,

    CONSTRAINT check_member_status
        CHECK (status IN ('active', 'inactive', 'alumni', 'suspended'))
);