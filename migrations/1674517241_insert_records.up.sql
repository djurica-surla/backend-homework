INSERT INTO question (body)
VALUES ('Which is the fastest land animal?'),
       ('How many continents are there?'),
       ('What is the tallest building in the world?');  

INSERT INTO answer (body, correct, question_id)
VALUES ('Monkey', 0, 1),
       ('Cow', 0, 1),
       ('Cheetah', 1, 1),
       ('4', 0, 2),
       ('11', 0, 2),
       ('7', 1, 2),
       ('Burj Khalifa', 1, 3),
       ('Shanghai Tower', 0, 3),
       ('One World Trade Center', 0, 3);