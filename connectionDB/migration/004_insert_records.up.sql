BEGIN;

INSERT INTO records(
	id_account_from, transaction_value, transaction_type)
	VALUES (1, 10, 'Pay in'),
	(2, 20, 'Pay in'),
	(3, 30, 'Pay in');
	
INSERT INTO records(
	id_account_from, transaction_value, transaction_type)
	VALUES (4, 40, 'Withdrawal'),
	(5, 50, 'Withdrawal'),
	(6, 60, 'Withdrawal');
	
INSERT INTO records(
	id_account_from, id_account_to, transaction_value, transaction_type)
	VALUES (7, 8, 78, 'Transfer'),
	(8, 9, 89, 'Transfer'),
	(9, 7, 97, 'Transfer');

COMMIT;