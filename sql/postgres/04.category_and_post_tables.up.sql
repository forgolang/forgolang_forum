CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title varchar(128) not null,
    description text null,
    slug varchar(200) not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc')
);

CREATE UNIQUE INDEX IF NOT EXISTS categories_slug_unique ON categories USING btree(slug);

CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    author_id bigint not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_posts_author_id FOREIGN KEY (author_id) REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE TABLE IF NOT EXISTS post_slugs (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    source_user_id bigint null,
    slug VARCHAR(220) not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_slugs_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_slugs_source_user_id FOREIGN KEY (source_user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE INDEX IF NOT EXISTS post_slugs_slug_index ON post_slugs USING btree(slug);

CREATE TABLE IF NOT EXISTS post_details (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    source_user_id bigint null,
    title varchar(200) not null,
    description text null,
    content text not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_details_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_details_source_user_id FOREIGN KEY (source_user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE TABLE IF NOT EXISTS post_category_assignments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    category_id bigint not null,
    source_user_id bigint null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_category_assignments_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_category_assignments_category_id FOREIGN KEY (category_id)
        REFERENCES categories(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_category_assignments_source_user_id FOREIGN KEY (source_user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE UNIQUE INDEX IF NOT EXISTS post_category_assignments_post_category_unique
    ON post_category_assignments USING btree(post_id, category_id);

CREATE TABLE IF NOT EXISTS post_votes_up (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    user_id bigint not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_votes_up_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_votes_up_user_id FOREIGN KEY (user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE UNIQUE INDEX IF NOT EXISTS post_votes_up_post_user_unique ON post_votes_up USING btree(post_id, user_id);

CREATE TABLE IF NOT EXISTS post_votes_down (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    user_id bigint not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_votes_down_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_votes_down_user_id FOREIGN KEY (user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE UNIQUE INDEX IF NOT EXISTS post_votes_down_post_user_unique ON post_votes_down USING btree(post_id, user_id);

CREATE TABLE IF NOT EXISTS post_comments (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    user_id bigint not null,
    parent_id bigint not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_comments_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_comments_user_id FOREIGN KEY (user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null,
    CONSTRAINT fk_post_comments_parent_id FOREIGN KEY (parent_id)
        REFERENCES post_comments(id) ON UPDATE cascade ON DELETE cascade
);

CREATE TABLE IF NOT EXISTS post_comment_details (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    comment_id bigint not null,
    comment text not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_comment_details_comment_id FOREIGN KEY (comment_id)
        REFERENCES post_comments(id) ON UPDATE cascade ON DELETE cascade
);

CREATE TABLE IF NOT EXISTS post_comment_votes_up (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    comment_id bigint not null,
    user_id bigint not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_comment_votes_up_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_comment_votes_up_comment_id FOREIGN KEY (comment_id)
        REFERENCES post_comments(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_comment_votes_up_user_id FOREIGN KEY (user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE UNIQUE INDEX IF NOT EXISTS post_comment_votes_up_post_comment_user_unique ON post_comment_votes_up USING btree(post_id, comment_id, user_id);

CREATE TABLE IF NOT EXISTS post_comment_votes_down (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    post_id bigint not null,
    comment_id bigint not null,
    user_id bigint not null,
    inserted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP at time zone 'utc'),

    CONSTRAINT fk_post_comment_votes_down_post_id FOREIGN KEY (post_id)
        REFERENCES posts(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_comment_votes_down_comment_id FOREIGN KEY (comment_id)
        REFERENCES post_comments(id) ON UPDATE cascade ON DELETE cascade,
    CONSTRAINT fk_post_comment_votes_down_user_id FOREIGN KEY (user_id)
        REFERENCES users(id) ON UPDATE cascade ON DELETE set null
);

CREATE UNIQUE INDEX IF NOT EXISTS post_comment_votes_down_post_user_unique ON post_comment_votes_down USING btree(post_id, comment_id, user_id);
