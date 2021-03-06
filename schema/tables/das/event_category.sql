CREATE TABLE DAS.EVENT_CATEGORY (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME TEXT NOT NULL UNIQUE,
  DESCRIPTION TEXT,
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO DAS.EVENT_CATEGORY (NAME, DESCRIPTION) VALUES ('Competitive Ballroom', 'Competitive dancesport event');
INSERT INTO DAS.EVENT_CATEGORY (NAME, DESCRIPTION) VALUES ('Show Dance', 'Show dance event. Not necessarily competitive.');
INSERT INTO DAS.EVENT_CATEGORY (NAME, DESCRIPTION) VALUES ('Cabaret', 'Cabaret dance');
