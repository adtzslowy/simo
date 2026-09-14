INSERT INTO organization_periods (
    organization_id,
    name,
    start_date,
    end_date,
    is_active
)
SELECT
    o.id,
    '2026/2027',
    DATE '2026-08-01',
    DATE '2027-07-31',
    TRUE
FROM organizations o
WHERE o.is_active = TRUE
  AND NOT EXISTS (
      SELECT 1
      FROM organization_periods op
      WHERE op.organization_id = o.id
        AND op.name = '2026/2027'
  );