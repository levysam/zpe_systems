INSERT INTO roles (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', 'role_admin', '*', '(GET)|(POST)|(PUT)|(PATCH)|(DELETE)', '', '', '');
INSERT INTO roles (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', 'role_modifier', '*', '(GET)|(POST)|(PUT)|(PATCH)', '', '', '');
INSERT INTO roles (p_type, v0, v1, v2, v3, v4, v5) VALUES ('g', '01JBFMM7PCVGFTQRNABNEQF749', 'role_admin', '', '', '', '');
INSERT INTO roles (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', 'role_watcher', '/users*', 'GET', '', '', '');

INSERT INTO users (ID, Name, Email, Password) VALUES ('01JBFMM7PCVGFTQRNABNEQF749', 'admin', 'admin@admin.com', '$2a$10$Lj4P5OZQE41dKPON2zFamu2KIPPXBPAjfDJTAJzfFusNxPo6Lb8yO');

