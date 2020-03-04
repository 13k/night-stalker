ALTER TABLE heroes
  ADD COLUMN aliases character varying(255)[],
  ADD COLUMN roles bigint[],
  ADD COLUMN role_levels bigint[],
  ADD COLUMN complexity integer,
  ADD COLUMN legs integer,
  ADD COLUMN attribute_primary bigint,
  ADD COLUMN attack_capabilities bigint;

