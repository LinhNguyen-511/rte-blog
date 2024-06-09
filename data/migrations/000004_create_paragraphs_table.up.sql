CREATE TABLE IF NOT EXISTS paragraphs (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    content_id INT REFERENCES posts_contents(content_id) ON DELETE CASCADE, 
    value TEXT DEFAULT ''
);