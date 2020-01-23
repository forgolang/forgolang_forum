CREATE TYPE tparty AS ENUM ('auth');

CREATE TABLE IF NOT EXISTS third_party (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name varchar(200) not null,
    code varchar(64) not null,
    type tparty not null,
    is_active boolean default true,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc')
);

CREATE UNIQUE INDEX third_party_code_unique ON third_party USING btree(code);

INSERT INTO third_party
    (name, code, type, is_active)
VALUES
    ('Github', 'github', 'auth', true);

CREATE TABLE IF NOT EXISTS user_comeback_apps (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id bigint not null,
    tparty_id bigint not null,
    access_token varchar(220) not null,
    refresh_token varchar(220) null,
    expire bigint null,
    data jsonb null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_user_comeback_apps_user_id FOREIGN KEY (user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_user_comeback_apps_tpary_id FOREIGN KEY (tparty_id)
        REFERENCES third_party(id) ON UPDATE cascade ON DELETE cascade
);

CREATE UNIQUE INDEX user_comeback_apps_user_tparty_unique ON user_comeback_apps USING btree(user_id, tparty_id);