CREATE TYPE otc AS ENUM ('2fa', 'confirmation', 'password_reset');

CREATE TABLE IF NOT EXISTS user_one_time_codes
(
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    user_id     bigint      not null,
    code        varchar(20) not null,
    type        otc                         default 'confirmation',
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_user_one_time_codes_user_id FOREIGN KEY (user_id)
        REFERENCES users (id) ON UPDATE cascade ON DELETE cascade
);

CREATE UNIQUE INDEX user_one_time_codes_code_unique ON user_one_time_codes USING btree (code);

CREATE TYPE twofactor AS ENUM ('app', 'email');

CREATE TABLE IF NOT EXISTS user_2fa
(
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    user_id     bigint not null,
    type        twofactor                   default 'email',
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_user_2fa_user_id FOREIGN KEY (user_id)
        REFERENCES users (id) ON UPDATE cascade ON DELETE cascade
);