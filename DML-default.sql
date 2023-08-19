/*
DML insert Defaul data dummy
*/
-- data instrument
INSERT INTO instruments (instrumen_name)
VALUES
	('drum kit'),
	('guitar'),
	('bass guitar'),
	('keyboard');

-- data room
INSERT INTO rooms (name)
VALUES
	('room A'),
	('room B');

-- data schedules
INSERT INTO schedule_types (name)
VALUES
	('pagi'),
	('siang'),
	('sore'),
	('malam');

INSERT INTO schedule_days (name)
VALUES
	('senin'),
	('selasa'),
	('rabu'),
	('kamis'),
	('jumat'),
	('sabtu'),
	('minggu');

INSERT INTO schedule_times (schedule_type_id, clock_start, clock_end)
VALUES
	(1, '08:00:00', '09:00:00'),
	(1, '09:05:00', '10:05:00'),
	(1, '10:10:00', '11:10:00'),
	(1, '11:15:00', '12:15:00'),
	(2, '13:00:00', '14:00:00'),
	(2, '14:05:00', '15:05:00'),
	(3, '15:10:00', '16:10:00'),
	(3, '16:15:00', '17:15:00'),
	(3, '17:20:00', '18:20:00'),
	(4, '19:05:00', '20:05:00'),
	(4, '20:10:00', '21:10:00'),
	(4, '21:15:00', '22:15:00'),
	(4, '22:20:00', '23:20:00');
