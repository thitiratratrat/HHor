CREATE DATABASE hhor;

CREATE TYPE dorm_type AS ENUM (
    'mixed',
    'female',
    'male');

\c hhor
CREATE EXTENSION citext
CREATE EXTENSION fuzzystrmatch
