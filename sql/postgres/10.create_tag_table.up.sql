CREATE TABLE IF NOT EXISTS tags (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name varchar(32) not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc')
);

CREATE TABLE IF NOT EXISTS post_tags (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    tag_id bigint not null,
    source_user_id bigint null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_tags_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_tags_tag_id FOREIGN KEY (tag_id)
        REFERENCES tags(id) ON UPDATE cascade ON DELETE cascade
);

CREATE UNIQUE INDEX post_tags_post_tag_unique ON post_tags USING btree(post_id, tag_id);