INSERT INTO question (body)
VALUES ('Which is the fastest land animal?'),
       ('How many continents are there?'),
       ('What is the tallest building in the world?');  

INSERT INTO question_option (body, correct, question_id)
VALUES ('Monkey', FALSE, 1),
       ('Cow', FALSE, 1),
       ('Cheetah', TRUE, 1),
       ('4', FALSE, 2),
       ('11', FALSE, 2),
       ('7', TRUE, 2),
       ('Burj Khalifa', TRUE, 3),
       ('Shanghai Tower', FALSE, 3),
       ('One World Trade Center', FALSE, 3);