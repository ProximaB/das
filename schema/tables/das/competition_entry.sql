-- Competition Entry for individual Athlete
CREATE TABLE IF NOT EXISTS DAS.COMPETITION_ENTRY_ATHLETE (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  ATHLETE_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  LEAD_INDICATOR BOOLEAN NOT NULL DEFAULT FALSE,
  LEAD_TAG INTEGER NOT NULL DEFAULT 101,
  CHECKIN_IND BOOLEAN NOT NULL DEFAULT FALSE, -- if user is not checked in, this is not count as entry
  CHECKIN_DATETIME TIMESTAMP,
  ORGANIZER_TAG TEXT,
  PAYMENT_IND BOOLEAN NOT NULL DEFAULT FALSE, -- if dancer has not paid registration, show as false
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, ATHLETE_ID)
);

CREATE INDEX ON DAS.COMPETITION_ENTRY_ATHLETE (COMPETITION_ID);
CREATE INDEX ON DAS.COMPETITION_ENTRY_ATHLETE (ATHLETE_ID);
CREATE INDEX ON DAS.COMPETITION_ENTRY_ATHLETE (CHECKIN_IND);
CREATE INDEX ON DAS.COMPETITION_ENTRY_ATHLETE (PAYMENT_IND);
CREATE INDEX ON DAS.COMPETITION_ENTRY_ATHLETE (CREATE_USER_ID);
CREATE INDEX ON DAS.COMPETITION_ENTRY_ATHLETE (UPDATE_USER_ID);

-- Competition Entry for Partnership
CREATE TABLE IF NOT EXISTS DAS.COMPETITION_ENTRY_PARTNERSHIP (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  PARTNERSHIP_ID INTEGER NOT NULL REFERENCES DAS.PARTNERSHIP (ID),
  CHECKIN_IND BOOLEAN NOT NULL DEFAULT FALSE, -- if user is not checked in, this is not count as entry
  CHECKIN_DATETIME TIMESTAMP,
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, PARTNERSHIP_ID)
);

CREATE INDEX ON DAS.COMPETITION_ENTRY_PARTNERSHIP (COMPETITION_ID);
CREATE INDEX ON DAS.COMPETITION_ENTRY_PARTNERSHIP (PARTNERSHIP_ID);


-- Competition Entry for Adjudicator
CREATE TABLE IF NOT EXISTS DAS.COMPETITION_ENTRY_ADJUDICATOR (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  ADJUDICATOR_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  CHECKIN_IND BOOLEAN NOT NULL DEFAULT FALSE, -- if user is not checked in, this is not count as entry
  CHECKIN_DATETIME TIMESTAMP,
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, ADJUDICATOR_ID)
);

CREATE INDEX ON DAS.COMPETITION_ENTRY_ADJUDICATOR (COMPETITION_ID);
CREATE INDEX ON DAS.COMPETITION_ENTRY_ADJUDICATOR (ADJUDICATOR_ID);


-- TODO: Competition Entry for Scrutineer
CREATE TABLE IF NOT EXISTS DAS.COMPETITION_ENTRY_SCRUTINEER (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  SCRUTINEER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, SCRUTINEER_ID)
);

-- TODO: Competition Entry for Deck Captain
CREATE TABLE IF NOT EXISTS DAS.COMPETITION_ENTRY_DECK_CAPTAIN (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  DECK_CAPTAIN_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, DECK_CAPTAIN_ID)
);

-- TODO: Competition Entry for Emcee
CREATE TABLE IF NOT EXISTS DAS.COMPETITION_ENTRY_EMCEE (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  EMCEE_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, EMCEE_ID)
);

CREATE TABLE IF NOT EXISTS DAS.COMPETITION_LEAD_TAG (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION(ID),
  LEAD_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  TAG_NUMBER INTEGER NOT NULL DEFAULT 101,
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, LEAD_ID, TAG_NUMBER)
);

CREATE INDEX ON DAS.COMPETITION_LEAD_TAG (COMPETITION_ID);
CREATE INDEX ON DAS.COMPETITION_LEAD_TAG (LEAD_ID);
CREATE INDEX ON DAS.COMPETITION_LEAD_TAG (TAG_NUMBER);