-- create das.studio table
CREATE TABLE IF NOT EXISTS DAS.STUDIO (
  ID SERIAL NOT NULL PRIMARY KEY,
  NAME TEXT NOT NULL,
  ADDRESS TEXT,
  CITY_ID INTEGER NOT NULL REFERENCES DAS.CITY (ID),
  WEBSITE TEXT,
  CREATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (NAME, ADDRESS, CITY_ID)
);

INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE) 
  VALUES ('Dance Center Chicago', '3868 North Lincoln Avenue',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Chicago' AND S.ABBREVIATION = 'IL' AND T.ABBREVIATION = 'USA'), 'http://dancecenterchicago.com');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE) 
  VALUES ('Kasper Dance Studio', '3201 N Long Ave',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Chicago' AND S.ABBREVIATION = 'IL' AND T.ABBREVIATION = 'USA'), 'http://kasperdancestudio.com');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE)
  VALUES ('Ballroom Center', '118 A. N Milwaukee Ave',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Libertyville' AND S.ABBREVIATION = 'IL' AND T.ABBREVIATION = 'USA'), 'http://www.theballroomcenter.com/');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE)
  VALUES ('Interclub Academy of Dance', '7350 N Milwaukee Ave',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Niles' AND S.ABBREVIATION = 'IL' AND T.ABBREVIATION = 'USA'), 'http://www.interclubdance.com');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE)
  VALUES ('Ballroom Dance Studio', '130 Hawthorn Center',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Vernon Hills' AND S.ABBREVIATION = 'IL' AND T.ABBREVIATION = 'USA'), 'http://www.ballroomdancestudiovh.com/home.html');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE)
  VALUES ('Starlite Ballroom', '5720 Guion Rd',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Indianapolis' AND S.ABBREVIATION = 'IN' AND T.ABBREVIATION = 'USA'), 'http://www.indianapolisdancelessons.com');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE)
  VALUES ('Aurelia Dance Studio', '3198 IN-32',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Westfield' AND S.ABBREVIATION = 'IN' AND T.ABBREVIATION = 'USA'), 'https://www.aureliadancestudio.com/');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE)
  VALUES ('Manhattan Ballroom Dance', '29 W 36th St',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'New York' AND S.ABBREVIATION = 'NY' AND T.ABBREVIATION = 'USA'), 'http://manhattanballroomdance.com/');
INSERT INTO DAS.STUDIO (NAME, ADDRESS, CITY_ID, WEBSITE)
  VALUES ('Kanopy Dance', '341 State St',
          (SELECT C.ID FROM DAS.CITY C 
            JOIN DAS.STATE S on C.STATE_ID = S.ID 
            JOIN DAS.COUNTRY T on S.COUNTRY_ID = T.ID 
          WHERE C.NAME = 'Madison' AND S.ABBREVIATION = 'WI' AND T.ABBREVIATION = 'USA'), 'http://kanopydance.org/');