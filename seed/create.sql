-- # executes SQL in create.sql, and then prompts for additional 
-- # SQL statements provided interactively. note that when specifying
-- # an init flag, the ~/.duckdbrc file is not read
-- duckdb -init create.sql

-- # executes SQL in create.sql and then immediately exits
-- # note that we're specifying a database name so that we
-- # can access the created data later. note that when specifying
-- # an init flag, the ~/.duckdbrc file is not read
-- duckdb -init create.sql -no-stdin mydb.ddb

-- duckdb < create.sql portof.ddb (pipe sql file to duckdb cli and exit)

CREATE SEQUENCE id_users START 1;
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('id_users'), 
    username VARCHAR NOT NULL,
    passw VARCHAR NOT NULL,
    is_verified BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edited_at TIMESTAMP
);

CREATE SEQUENCE id_page START 1;
CREATE TABLE pages (
    id INTEGER PRIMARY KEY DEFAULT nextval('id_page'),
    title VARCHAR NOT NULL,
    parent_id INTEGER,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edited_at TIMESTAMP
);

CREATE SEQUENCE id_category START 1;
CREATE TABLE categories (
    id INTEGER PRIMARY KEY DEFAULT nextval('id_category'),
    page_id INTEGER NOT NULL,
    title VARCHAR NOT NULL,
    parent_id INTEGER,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edited_at TIMESTAMP
);

CREATE SEQUENCE id_post START 1;
CREATE TABLE posts (
    id INTEGER PRIMARY KEY DEFAULT nextval('id_post'),
    page_id INTEGER NOT NULL,
    title VARCHAR NOT NULL,
    short_desc VARCHAR,
    the_body VARCHAR,
    display_picture_url VARCHAR,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edited_at TIMESTAMP
);