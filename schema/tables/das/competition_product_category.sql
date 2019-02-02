CREATE TABLE IF NOT EXISTS DAS.COMPETITION_PRODUCT_CATEGORY (
  ID SERIAL NOT NULL  PRIMARY KEY ,
  NAME VARCHAR(32) NOT NULL UNIQUE,
  DESCRIPTION TEXT,
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO DAS.COMPETITION_PRODUCT_CATEGORY(NAME, DESCRIPTION) VALUES ('Athlete Competition Pass', 'Ticket that can be used to compete at any events at a competition');
INSERT INTO DAS.COMPETITION_PRODUCT_CATEGORY(NAME, DESCRIPTION) VALUES ('Athlete Event Pass', 'Ticket that can only be used to compete a particular event');
INSERT INTO DAS.COMPETITION_PRODUCT_CATEGORY(NAME, DESCRIPTION) VALUES ('Spectator Competition Pass', 'Ticket that can be used to spectate any event throughout the competition');
INSERT INTO DAS.COMPETITION_PRODUCT_CATEGORY(NAME, DESCRIPTION) VALUES ('Spectator Event Pass', 'Ticket that can only be used to spectate a particular event');
INSERT INTO DAS.COMPETITION_PRODUCT_CATEGORY(NAME, DESCRIPTION) VALUES ('Spectator Session Pass', 'Ticket that can only be used to spectate events during a particular session');
INSERT INTO DAS.COMPETITION_PRODUCT_CATEGORY(NAME, DESCRIPTION) VALUES ('Private Lesson', 'Ticket that can be used to take one session of private lesson');
INSERT INTO DAS.COMPETITION_PRODUCT_CATEGORY(NAME, DESCRIPTION) VALUES ('Workshop', 'Ticket that can be used to attend one session of workshop');