DROP TABLE IF EXISTS deck_card;
DROP TABLE IF EXISTS deck;
DROP TABLE IF EXISTS "card";

CREATE TABLE IF NOT EXISTS deck (
  id serial PRIMARY KEY,
  "name" text NOT NULL,
  format text NOT NULL,
  card_cnt int NOT NULL,
  last_edit timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "card" (
  id serial PRIMARY KEY,
  "name" text NOT NULL,
  wiz_id text NOT NULL,
  "set" text NOT NULL,
  oracle text,
  "type" text NOT NULL,
  white boolean NOT NULL DEFAULT FALSE,
  blue boolean NOT NULL DEFAULT FALSE,
  black boolean NOT NULL DEFAULT FALSE,
  red boolean NOT NULL DEFAULT FALSE,
  green boolean NOT NULL DEFAULT FALSE,
  mana_value decimal NOT NULL,
  mana_cost text NOT NULL,
  power text,
  toughness text,
  loyalty text,
  img jsonb NOT NULL,
  quantity int NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS deck_card (
  id serial PRIMARY KEY,
  deck_id int REFERENCES deck(id),
  card_id int REFERENCES "card"(id)
);
