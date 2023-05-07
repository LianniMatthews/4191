CREATE TABLE courses (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    code text NOT NULL,
    title text NOT NULL,
    credit int NOT NULL,
    version integer NOT NULL DEFAULT 1
);