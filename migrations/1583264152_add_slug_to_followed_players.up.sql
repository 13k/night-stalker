ALTER TABLE followed_players
  ADD COLUMN slug varchar(255);

UPDATE followed_players
  SET slug = trim(
    regexp_replace(
      replace(
        regexp_replace(
          lower(label),
          '\W',
          '-',
          'g'
        ),
        '_',
        '-'
      ),
      '-+',
      '-',
      'g'
    ),
    '-'
  );

ALTER TABLE followed_players
  ALTER COLUMN slug SET NOT NULL;
