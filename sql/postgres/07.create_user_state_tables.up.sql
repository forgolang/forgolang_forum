CREATE TYPE state AS ENUM ('active', 'banned', 'wait_for_confirmation');

CREATE TABLE IF NOT EXISTS user_states (
   id BIGSERIAL NOT NULL PRIMARY KEY,
   user_id bigint not null,
   state state default 'wait_for_confirmation',
   source_user_id bigint null,
   inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

   CONSTRAINT fk_user_states_user_id FOREIGN KEY (user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE cascade,
   CONSTRAINT fk_user_states_source_user_id FOREIGN KEY (source_user_id)
       REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);