
-- 1. Create a project
INSERT INTO projects (id, name)
VALUES (
    '11111111-1111-4111-8111-111111111111',
    'School Management'
)
ON CONFLICT DO NOTHING;

-- 2. Create the root collection
INSERT INTO collections (id, project_id, name)
VALUES (
    '22222222-2222-4222-8222-222222222222',
    '11111111-1111-4111-8111-111111111111',
    'students'
)
ON CONFLICT DO NOTHING;

-- 3. Create a student document
INSERT INTO documents (
    id, project_id, collection_id, document_id, data
)
VALUES (
    '33333333-3333-4333-8333-333333333333',
    '11111111-1111-4111-8111-111111111111',
    '22222222-2222-4222-8222-222222222222',
    'student_001',
    '{
      "name": "Kamal Perera",
      "age": 21,
      "email": "kamal@example.com",
      "active": true
    }'::jsonb
)
ON CONFLICT DO NOTHING;

