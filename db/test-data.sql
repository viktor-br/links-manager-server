\c test;
DELETE FROM users;
DELETE FROM sessions;
INSERT INTO users (id,username,password,created_at,role) VALUES ('2d257d0e-511a-4ff3-8c88-b239763e6121','admin','f865b53623b121fd34ee5426c792e5c33af8c227',now(),1);