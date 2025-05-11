
-- +migrate Up
INSERT INTO courier (id, name, code, created_at, created_by, updated_at, updated_by)
VALUES ('411cc8ce-df6a-4b60-ad5c-28f60934d0b3', 'Illham Express', 'ILX', now(), '11111111-1111-1111-1111-111111111111', now(), '11111111-1111-1111-1111-111111111111'),
       ('0196be3c-4d04-7409-b54a-7f55dc9765e7', 'Kubik Cargo', 'KCG', now(), '11111111-1111-1111-1111-111111111111', now(), '11111111-1111-1111-1111-111111111111'),
       ('bbe1da31-0d85-4240-9d1f-8ad8f2bb21d7', 'Siantar Aja', 'SAA', now(), '11111111-1111-1111-1111-111111111111', now(), '11111111-1111-1111-1111-111111111111');

-- +migrate Down
DELETE FROM courier where id IN (
    '411cc8ce-df6a-4b60-ad5c-28f60934d0b3',
    '0196be3c-4d04-7409-b54a-7f55dc9765e7',
    'bbe1da31-0d85-4240-9d1f-8ad8f2bb21d7'
);