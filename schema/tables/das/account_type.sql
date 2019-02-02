CREATE TABLE IF NOT EXISTS DAS.ACCOUNT_TYPE (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME VARCHAR(12) NOT NULL UNIQUE,
  DESCRIPTION TEXT,
  DATETIME_CREATED TIMESTAMP DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP DEFAULT NOW()
);
INSERT INTO DAS.ACCOUNT_TYPE(NAME, DESCRIPTION) VALUES('Athlete','An athlete who competes at ballroom dance competitions');
INSERT INTO DAS.ACCOUNT_TYPE(NAME, DESCRIPTION) VALUES('Adjudicator','Judge');
INSERT INTO DAS.ACCOUNT_TYPE(NAME, DESCRIPTION) VALUES('Scrutineer','Chair person of judge(not necessarily a scrutineer)');
INSERT INTO DAS.ACCOUNT_TYPE(NAME, DESCRIPTION) VALUES('Organizer','Organizer of the competition');
INSERT INTO DAS.ACCOUNT_TYPE(NAME, DESCRIPTION) VALUES('Deck Captain','Check in competitors on deck and finalize heat line-up');
INSERT INTO DAS.ACCOUNT_TYPE(NAME, DESCRIPTION) VALUES('Emcee','Announcer of events');
INSERT INTO DAS.ACCOUNT_TYPE(NAME, DESCRIPTION) VALUES('Admin','DAS administrator who has master access and control to almost everything');