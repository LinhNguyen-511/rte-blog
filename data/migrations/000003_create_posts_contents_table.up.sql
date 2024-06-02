CREATE TYPE ContentType AS ENUM ('paragraph', 'code block', 'check list');
CREATE TABLE IF NOT EXISTS posts_contents (
    post_id INT REFERENCES posts ON DELETE CASCADE, 
    content_id SERIAL UNIQUE,
    content_type ContentType NOT NULL, 
    order_in_post INT NOT NULL CHECK (order_in_post >= 0),
    PRIMARY KEY (post_id, content_id, content_type)
)