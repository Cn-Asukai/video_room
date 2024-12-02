-- auto-generated definition
create schema video collate utf8mb4_0900_ai_ci;
use video
create table video_comments
(
    id          int      null,
    video_id    int      null,
    content     text     null,
    commentator int      null,
    created_at  datetime null,
    updated_at  datetime null,
    deleted_at  datetime null
);

create table video_tag
(
    id       int null,
    video_id int null,
    tag_id   int null
);

create table videos
(
    id         int         null,
    created_at datetime    null,
    updated_at timestamp   null,
    deleted_at datetime    null,
    name       varchar(20) null,
    url        text        null,
    cover_url  text        null,
    state      int         null
);

