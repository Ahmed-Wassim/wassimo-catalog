-- +goose Up
INSERT INTO restaurants (name, status) VALUES
    ('La Mama', 'active'),
    ('Trattoria Milano', 'active'),
    ('Sushi House', 'active'),
    ('Burger Palace', 'inactive'),
    ('Green Garden', 'pending');

-- +goose Down
DELETE FROM restaurants
WHERE name IN (
    'La Mama',
    'Trattoria Milano',
    'Sushi House',
    'Burger Palace',
    'Green Garden'
);