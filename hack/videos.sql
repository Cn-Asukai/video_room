create table videos
(
    id         int auto_increment
        primary key,
    created_at datetime    null,
    updated_at timestamp   null,
    deleted_at datetime    null,
    name       varchar(20) null,
    url        text        null,
    cover_url  text        null,
    state      int         null,
    creator    int         null
);

INSERT INTO video.videos (id, created_at, updated_at, deleted_at, name, url, cover_url, state, creator) VALUES (1, '2024-12-02 08:20:00', '2024-12-02 08:20:00', '2024-12-02 09:49:33', '论间更', 'http://rdei.bo/bmsdim', 'http://kqgqz.fk/eemi', null, null);
INSERT INTO video.videos (id, created_at, updated_at, deleted_at, name, url, cover_url, state, creator) VALUES (2, '2024-12-02 08:20:06', '2024-12-02 08:20:06', '2024-12-02 09:55:52', '论间更', 'http://rdei.bo/bmsdim', 'http://kqgqz.fk/eemi', null, null);
INSERT INTO video.videos (id, created_at, updated_at, deleted_at, name, url, cover_url, state, creator) VALUES (3, '2024-12-02 08:33:19', '2024-12-02 08:33:19', null, 'test1', 'http://127.0.0.1:8000/test1.mp4', 'http://127.0.0.1:8000/test1.png', null, null);
INSERT INTO video.videos (id, created_at, updated_at, deleted_at, name, url, cover_url, state, creator) VALUES (4, '2024-12-02 08:33:31', '2024-12-02 08:33:31', null, 'test2', 'http://127.0.0.1:8000/test2.mp4', 'http://127.0.0.1:8000/test2.png', null, null);
