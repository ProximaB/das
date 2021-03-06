-- table for event's dances: one competitive event can map to multiple dances
CREATE TABLE IF NOT EXISTS DAS.EVENT_DANCES (
  ID SERIAL PRIMARY KEY NOT NULL,
  EVENT_ID INTEGER NOT NULL REFERENCES DAS.EVENT (ID),
  DANCE_ID INTEGER NOT NULL REFERENCES DAS.DANCE(ID),
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (EVENT_ID, DANCE_ID) -- a event should not have two identical dances
);

CREATE INDEX ON DAS.EVENT_DANCES (EVENT_ID);
CREATE INDEX ON DAS.EVENT_DANCES (DANCE_ID);
CREATE INDEX ON DAS.EVENT_DANCES (CREATE_USER_ID);