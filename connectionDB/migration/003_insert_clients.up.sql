BEGIN;

INSERT INTO account(
	name, cpf, balance_account)	VALUES 
	('Naruto Uzumaki', 85296916018, 100), 
	('Sasuke Uchiha', 20059671068, 200),
	('Sakura Haruno', 73244203035, 300),
	('Kakashi Hatake', 71558629084, 400),
	('Jaraiya', 55928441070, 500),
	('Orochimaru', 52427950009, 600),
	('Minato Namikaze', 65633544080, 700),
	('Itachi Uchiha', 50219501009, 800),
	('Obito Uchiha', 07233100098, 900);

COMMIT;