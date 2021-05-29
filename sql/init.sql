CREATE TABLE students(
  id int generated always as identity PRIMARY KEY,
  mail varchar(255) not null,
  full_name varchar(255) not null,
  user_name varchar(255) unique not null,
  password varchar(255)
);

CREATE TABLE subjects (
  id int generated always as identity PRIMARY KEY,
  type varchar(255) unique not null
);

CREATE TABLE subject_modules (
  id int generated always as identity PRIMARY KEY,
  subject_id int not null,
  type varchar(255) unique not null
);

CREATE TABLE teachers (
  id int generated always as identity PRIMARY KEY,
  mail varchar(255) unique not null,
  full_name varchar(255) not null,
  user_name varchar(255) unique not null,
  password varchar(255) not null,
  subject_id int not null
);

CREATE TABLE class (
  id int generated always as identity PRIMARY KEY,
  name varchar(255) unique not null,
  teacher_id int not null,
  is_use boolean not null
);

CREATE TABLE courses (
  id int generated always as identity PRIMARY KEY,
  class_id int not null,
  start_date timestamp not null,
  end_date timestamp not null,
  grade int not null
);

CREATE TABLE dayofweek (
  id int generated always as identity PRIMARY KEY,
  day varchar(10) unique not null
);

CREATE TABLE class_session (
  id int generated always as identity PRIMARY KEY,
  class_id int not null,
  dayofweek_id int not null,
  start_time time not null,
  end_time time not null
);

CREATE TABLE courses_class_session (
  courses_id int not null,
  class_session_id int not null,
  subject_module_id int not null,
  at_date date,
  PRIMARY KEY (courses_id, class_session_id)
);

CREATE TABLE students_courses (
  student_id int not null,
  course_id int not null,
  PRIMARY KEY (student_id, course_id)
);

ALTER TABLE subject_modules ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);

ALTER TABLE teachers ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);

ALTER TABLE class ADD FOREIGN KEY (teacher_id) REFERENCES teachers (id);

ALTER TABLE courses ADD FOREIGN KEY (class_id) REFERENCES class (id) on delete cascade;

ALTER TABLE class_session ADD FOREIGN KEY (class_id) REFERENCES class (id) on delete cascade;

ALTER TABLE class_session ADD FOREIGN KEY (dayofweek_id) REFERENCES dayofweek (id);

ALTER TABLE courses_class_session ADD FOREIGN KEY (courses_id) REFERENCES courses (id) on delete set null;

ALTER TABLE courses_class_session ADD FOREIGN KEY (class_session_id) REFERENCES class_session (id) on delete cascade;

ALTER TABLE courses_class_session ADD FOREIGN KEY (subject_module_id) REFERENCES subject_modules (id) on delete cascade;

ALTER TABLE students_courses ADD FOREIGN KEY (student_id) REFERENCES students (id);

ALTER TABLE students_courses ADD FOREIGN KEY (course_id) REFERENCES courses (id) on delete cascade;

