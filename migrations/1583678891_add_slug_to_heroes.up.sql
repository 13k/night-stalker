ALTER TABLE heroes
  ADD COLUMN slug varchar(255);

UPDATE heroes
  SET slug = replace(regexp_replace(name, '^npc_dota_hero_', ''), '_', '-');

ALTER TABLE heroes
  ALTER COLUMN slug SET NOT NULL;

CREATE UNIQUE INDEX uix_heroes_slug ON heroes (slug);
