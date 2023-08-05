CREATE TABLE users (
  ID SERIAL PRIMARY KEY,
  Token TEXT NOT NULL UNIQUE,
  Name TEXT,
  EducationMode INT,
  EducationComplexity INT
);
create index on users(Token);

CREATE TABLE lessons (
  ID SERIAL PRIMARY KEY,
  Title TEXT NOT NULL,
  BlockingLesson INT,
  BlockingText TEXT
);

CREATE TABLE lessons_tags (
  ID SERIAL PRIMARY KEY,
  LessonID INT REFERENCES lessons(ID),
  Tag TEXT NOT NULL
);

CREATE TABLE completed_lessons (
  ID SERIAL PRIMARY KEY,
  UserID INT REFERENCES users(ID),
  LessonID INT REFERENCES lessons(ID)
);

insert into lessons (Title, BlockingLesson, BlockingText) values ('Урок 1: Как купить Ж/д билет?', 0, '');
insert into lessons (Title, BlockingLesson, BlockingText) values ('Урок 2: Как купить авиабилет?', 1, 'Откроется после 1 урока');
insert into lessons (Title, BlockingLesson, BlockingText) values ('Урок 3: Как заказать такси?', 2, 'Откроется после 2 урока');

insert into lessons_tags (LessonID, Tag) values (1, 'Инструкция');
insert into lessons_tags (LessonID, Tag) values (1, '5 слов');
insert into lessons_tags (LessonID, Tag) values (1, 'Пример диалога');
insert into lessons_tags (LessonID, Tag) values (1, 'Практика');

insert into lessons_tags (LessonID, Tag) values (2, 'Инструкция');
insert into lessons_tags (LessonID, Tag) values (2, '5 слов');
insert into lessons_tags (LessonID, Tag) values (2, 'Пример диалога');
insert into lessons_tags (LessonID, Tag) values (2, 'Практика');

insert into completed_lessons (UserID, LessonID) values (1, 1);

