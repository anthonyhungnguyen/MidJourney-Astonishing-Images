-- Create a table to store image url, caption and keywords
CREATE TABLE image (
    id SERIAL PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    caption VARCHAR(255) NOT NULL,
    keywords VARCHAR(255) NOT NULL
);