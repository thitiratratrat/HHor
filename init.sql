CREATE DATABASE hhor;

CREATE TYPE dorm_type AS ENUM (
    'mixed',
    'female',
    'male');

CREATE EXTENSION citext