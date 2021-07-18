CREATE TABLE users(
  id int generated always as identity PRIMARY KEY,
  email varchar(255) not null,
  full_name varchar(255) not null,
  user_name varchar(255) unique not null,
  password varchar(255) not null,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp
);

CREATE TABLE subjects (
  id int generated always as identity PRIMARY KEY,
  type varchar(255) unique not null,
  is_delete boolean not null default 'false'
);

CREATE TABLE roles (
  id int generated always as identity PRIMARY KEY,
  type varchar(255) unique not null,
  created_at timestamptz not null default current_timestamp
);

CREATE TABLE permissions (
  id int generated always as identity PRIMARY KEY,
  type varchar(255) unique not null,
  created_at timestamptz not null default current_timestamp
);

CREATE TABLE resources (
  id int generated always as identity PRIMARY KEY,
  type varchar(255) unique not null,
  created_at timestamptz not null default current_timestamp
);

CREATE TABLE role_permissions_resource (
  role_id int not null,
  permission_id int not null,
  resource_id int not null,
  PRIMARY KEY (role_id, permission_id, resource_id)
);

CREATE TABLE user_role_course (
  user_id int not null,
  role_id int not null,
  course_id int not null,
  PRIMARY KEY (user_id, role_id, course_id)
);

CREATE TABLE user_role_class (
  user_id int not null,
  role_id int not null,
  class_id int not null,
  PRIMARY KEY (user_id, role_id, class_id)
);

CREATE TABLE role_permissions_resource (
  role_id int not null,
  permission_id int not null,
  resource_id int not null
);

CREATE TABLE subject_modules (
  id int generated always as identity PRIMARY KEY,
  subject_id int not null,
  type varchar(255) unique not null,
  is_delete boolean not null default 'false'
);

CREATE TABLE courses (
  id int generated always as identity PRIMARY KEY,
  name varchar(255) unique not null,
  description varchar(255),
  subject_id int not null,
  is_delete boolean not null default 'false',
  is_open boolean not null default 'true'
);

CREATE TABLE classes (
  id int generated always as identity PRIMARY KEY,
  name varchar(255) unique not null,
  course_id int not null,
  start_date  date not null,
  end_date date not null,
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
  is_use boolean not null default 'true',
  start_time time not null,
  end_time time not null
);

ALTER TABLE subject_modules ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);

ALTER TABLE role_permissions_resource ADD FOREIGN KEY (role_id) REFERENCES roles (id);

ALTER TABLE role_permissions_resource ADD FOREIGN KEY (permission_id) REFERENCES permissions (id);

ALTER TABLE role_permissions_resource ADD FOREIGN KEY (resource_id) REFERENCES resources (id);

ALTER TABLE user_role_course ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_role_course ADD FOREIGN KEY (role_id) REFERENCES roles (id);

ALTER TABLE user_role_course ADD FOREIGN KEY (course_id) REFERENCES courses (id);

ALTER TABLE user_role_class ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_role_class ADD FOREIGN KEY (role_id) REFERENCES roles (id);

ALTER TABLE user_role_class ADD FOREIGN KEY (class_id) REFERENCES classes (id);

ALTER TABLE classes ADD FOREIGN KEY (course_id) REFERENCES courses (id) on delete cascade;

ALTER TABLE courses ADD FOREIGN KEY (subject_id) REFERENCES subjects (id) on delete cascade;

ALTER TABLE class_session ADD FOREIGN KEY (class_id) REFERENCES classes (id) on delete cascade;

ALTER TABLE class_session ADD FOREIGN KEY (dayofweek_id) REFERENCES dayofweek (id);