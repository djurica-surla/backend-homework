-- Create question table
CREATE TABLE IF NOT EXISTS question (
    id INTEGER PRIMARY KEY,
    body VARCHAR(255)
);

-- Create answer table
CREATE TABLE IF NOT EXISTS answer (
    id INTEGER PRIMARY KEY,
    body VARCHAR(255),
    correct INTEGER NOT NULL CHECK (correct IN (0, 1)),
    question_id INTEGER NOT NULL,
    CONSTRAINT fk_question
    FOREIGN KEY (question_id)
    REFERENCES question(id)
);