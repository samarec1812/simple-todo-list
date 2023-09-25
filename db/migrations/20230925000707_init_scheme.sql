-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (username)
);

CREATE TABLE todo_lists
(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(255),

    PRIMARY KEY (id)
);

CREATE TABLE users_lists
(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id BIGINT REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE todo_items
(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    deadline timestamp NOT NULL,
    done boolean NOT NULL DEFAULT FALSE,

    PRIMARY KEY (id)
);

CREATE TABLE lists_items
(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    item_id BIGINT REFERENCES todo_items(id) ON DELETE CASCADE NOT NULL,
    list_id BIGINT REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL,

    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE lists_items;

DROP TABLE users_lists;

DROP TABLE todo_lists;

DROP TABLE users;

DROP TABLE todo_items;
-- +goose StatementEnd
