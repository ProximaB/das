CREATE TABLE IF NOT EXISTS DAS.COMPETITION_STATUS (
  ID SERIAL NOT NULL PRIMARY KEY ,
  NAME TEXT NOT NULL UNIQUE,
  ABBREVIATION CHARACTER(1) NOT NULL UNIQUE,
  DESCRIPTION TEXT NOT NULL,
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO DAS.COMPETITION_STATUS(NAME, DESCRIPTION, ABBREVIATION) VALUES ('Pre-Registration','Initial Status when competition is created the first time','P');
INSERT INTO DAS.COMPETITION_STATUS(NAME, DESCRIPTION, ABBREVIATION) VALUES ('Open Registration','Competition is open to eligible couples to register','O');
INSERT INTO DAS.COMPETITION_STATUS(NAME, DESCRIPTION, ABBREVIATION) VALUES ('Closed Registration','No more couples can register any events but scrutineer can modify or manually add/remove/edit','C');
INSERT INTO DAS.COMPETITION_STATUS(NAME, DESCRIPTION, ABBREVIATION) VALUES ('In Progress','Competition is running','I');
INSERT INTO DAS.COMPETITION_STATUS(NAME, DESCRIPTION, ABBREVIATION) VALUES ('Processing','Results are visible and rank/ratings are not updated','S');
INSERT INTO DAS.COMPETITION_STATUS(NAME, DESCRIPTION, ABBREVIATION) VALUES ('Closed','Results are visible and rank/ratings are updated','D');
INSERT INTO DAS.COMPETITION_STATUS(NAME, DESCRIPTION, ABBREVIATION) VALUES ('Cancelled','Competition is cancelled for any reason and no rating or rank is updated','L');