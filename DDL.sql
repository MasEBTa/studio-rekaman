/*
DDL CREATE TABLE
*/
-- TABLE BANDS
CREATE TABLE bands (
	id SERIAL PRIMARY KEY,
	name_band VARCHAR(50),
	captain VARCHAR(50),
	email VARCHAR(50) UNIQUE,
	membership BOOL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABLE INSTRUMENTS
CREATE TABLE instruments (
	id SERIAL PRIMARY KEY,
	instrumen_name VARCHAR(50), -- drum kit, guitar, bass, keyboard {bisa ditambah lagi} 
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE instrument_types (
  id SERIAL PRIMARY KEY,
  instrument_id INT REFERENCES instruments(id),
	type VARCHAR(50),
	total INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABLE INSTRUMENT OF BAND
CREATE TABLE band_instruments (
  id SERIAL PRIMARY KEY,
  band_id INT REFERENCES bands(id),
  instrument_type_id INT REFERENCES instruments(id),
  total INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABLE ROOMS
CREATE TABLE rooms (
  id SERIAL PRIMARY KEY,
	name VARCHAR(50), -- room A, room B {bisa di tambah lagi}
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABLE SCHEDULE
CREATE TABLE schedule_types (
  id SERIAL PRIMARY KEY,
	name VARCHAR(50), -- pagi, siang, sore, malam
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE schedule_days (
  id SERIAL PRIMARY KEY,
	name VARCHAR(50), -- senin, selasa, rabu, kamis, jumat, sabtu, minggu
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE schedule_times (
  id SERIAL PRIMARY KEY,
	schedule_type_id INT REFERENCES schedule_types(id),
	clock_start TIME,
	clock_end TIME,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
  /*pagi
    - 08.00-09.00
    - 09.05-10.05
    - 10.10-11.10
    - 11.15-12.15
  */
  /*siang
    - 13.00-14.00
    - 14.05-15.05
  */
  /*sore
    - 15.10-16.10
    - 16.15-17.15
    - 17.20-18.20
  */
  /*malam
    - 19.05-20.05
    - 20.10-21.10
    - 21.15-22.15
    - 22.20-23.20
  */

CREATE TABLE schedule_assigns (
  id SERIAL PRIMARY KEY,
	room_id INT REFERENCES rooms(id),
	schedule_day_id INT REFERENCES schedule_days(id),
	schedule_time_id INT REFERENCES schedule_times(id),
	band_id INT REFERENCES bands(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TABLE TRANSACTION
CREATE TABLE transaction (
  id SERIAL PRIMARY KEY,
	band_id INT REFERENCES bands(id),
	schedule_assign_id INT REFERENCES schedule_assigns(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
/*
DDL DROP TABLE
*/
/*
-- table transaction
DROP TABLE transaction;
-- table schedule
DROP TABLE schedule_assigns;
DROP TABLE schedule_times;
DROP TABLE schedule_days;
DROP TABLE schedule_types;
-- TABLE INSTRUMENT OF BAND
DROP TABLE band_instruments;
-- table bands
DROP TABLE bands;
-- table instruments
DROP TABLE instrument_types;
DROP TABLE instruments;
-- table room
DROP TABLE rooms;
*/