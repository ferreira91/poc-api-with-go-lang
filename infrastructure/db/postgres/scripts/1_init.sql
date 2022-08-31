CREATE TABLE market (
   id SERIAL PRIMARY KEY,
   longitude VARCHAR ( 10 ) NOT NULL,
   latitude VARCHAR ( 10 ) NOT NULL,
   census_sector VARCHAR ( 15 ) NOT NULL,
   weighting_area VARCHAR ( 13 ) NOT NULL,
   township_code VARCHAR ( 9 ) NOT NULL,
   township VARCHAR ( 18 ) NOT NULL,
   subprefecture_code VARCHAR ( 2 ) NOT NULL,
   subprefecture VARCHAR ( 25 ) NOT NULL,
   region_5 VARCHAR ( 6 ) NOT NULL,
   region_8 VARCHAR ( 7 ) NOT NULL,
   name VARCHAR ( 30 ) NOT NULL,
   registry VARCHAR ( 6 ) UNIQUE NOT NULL,
   street VARCHAR ( 34 ) NOT NULL,
   number VARCHAR ( 15 ),
   district VARCHAR ( 20 ),
   reference VARCHAR ( 30 )
);
