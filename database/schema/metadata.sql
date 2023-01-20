-- Create a metadata table to store width, height, enqueue_time, hot_score
-- which has the primary serial and image_id as foreign key
CREATE TABLE metadata (
    id SERIAL PRIMARY KEY,
    image_id INTEGER NOT NULL,
    width INTEGER NOT NULL,
    height INTEGER NOT NULL,
    enqueue_time TIMESTAMP NOT NULL,
    hot_score INTEGER NOT NULL,
    FOREIGN KEY (image_id) REFERENCES image(id)
);