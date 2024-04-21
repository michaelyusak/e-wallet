\c e_wallet_db;

insert into users (user_email, user_name, user_password, created_at)
values
	('bank@example.com', 'bank wallet', '$2a$12$sUvKnupsWbyd9h7kjtW/RusaHbT43YTj.PZvYMw2.4OSUPbgiRh3q', NOW()),
	('retail@example.com', 'retail wallet', '$2a$12$ShYgQ/rOhsrt7hGm0tNvfuYDunEEQ9b5Upg1y7VumKTPrpi.A2TMi', NOW()),
	('company@example.com', 'company wallet', '$2a$12$i2O.LGGxiYP9KqIoF.JFFuVziPx.aR.i5HFvQmjbXND1DdnZGwyiy', NOW());

insert into users (user_email, user_name, user_password, created_at)
values
	('alice@example.com', 'alice', '$2a$12$Gpd3MQU6o/36dGo5CnKno.qATJvPXyw5aq1j0p7hcZagXxfmJriy2', NOW()),
	('bob@example.com', 'bob', '$2a$12$Pn/ftEa0hnSSv.Wa/r6BS.j2HM/yivsOA4vr8oAGPh4DU8Abl4f2W', NOW()),
	('charlie@example.com', 'charlie', '$2a$12$xJydTODvQRbsH7jLIKms.exp6J99EYQa7VdrNNMOxzMPCW9Ar/uki', NOW()),
	('david@example.com', 'david', '$2a$12$Jb2WCa3NsLJSkANRiWjH4ezGQYHSjqb2O/270qyLy5E1Fx98hlzzu', NOW()),
	('eric@example.com', 'eric', '$2a$12$EvzyrdMgekO4lJtxLCRQF.wNDxAqpNMRAnFjsMAFC0ZOnMGt.vjEO', NOW());

insert into wallets (user_id, gacha_trial, created_at)
values
	(1, 0, NOW()),
	(2, 0, NOW()),
	(3, 0, NOW()),
	(4, 0, NOW()),
	(5, 0, NOW()),
	(6, 0, NOW()),
	(7, 0, NOW()),
	(8, 0, NOW());

update wallets 
set balance = 1000000000000 
where user_id 
	in (1, 2, 3);

insert into transactions (sender_wallet_number, recepient_wallet_number, amount, source_of_fund, description, created_at)
values
	('1000000000001', '1000000000004', 50000000, 'Bank Transfer', 'Top Up from Bank Transfer', '2023-02-02'),
	('1000000000002', '1000000000005', 5000000, 'Cash', 'Top Up from Cash', '2022-01-01'),
	('1000000000001', '1000000000006', 5000000, 'Bank Transfer', 'Top Up from Bank Transfer', '2022-09-30'),
	('1000000000002', '1000000000007', 5000000, 'Cash', 'Top Up from Cash', '2023-01-31'),
	('1000000000001', '1000000000008', 5000000, 'Bank Transfer', 'Top Up from Bank Transfer', '2022-12-31'),
	('1000000000008', '1000000000004', 30000, 'Wallet', 'nasi padang', '2023-01-31'),
	('1000000000004', '1000000000006', 20000, 'Wallet', 'ongkos gojek', '2023-02-28'),
	('1000000000004', '1000000000007', 5000, 'Wallet', 'ongkos angkot', '2021-12-31'),
	('1000000000005', '1000000000004', 500000, 'Wallet', 'Fancy dinner', '2022-07-01'),
	('1000000000006', '1000000000007', 500000, 'Wallet', 'Utang', '2023-01-29'),
	('1000000000001', '1000000000004', 1000000, 'Bank Transfer', 'Top Up from Bank Transfer', '2022-12-02'),
	('1000000000004', '1000000000008', 300000, 'Wallet', 'dinner', '2023-02-03'),
	('1000000000006', '1000000000007', 50000, 'Wallet', 'tiket', '2023-01-02'),
	('1000000000007', '1000000000004', 100000, 'Wallet', 'Tiket bus', '2022-12-02'),
	('1000000000008', '1000000000006', 5000000, 'Wallet', 'Tiket Pesawat', '2022-10-02'),
	('1000000000004', '1000000000007', 70000, 'Wallet', 'tiket bioskop', '2023-02-20'),
	('1000000000006', '1000000000005', 100000, 'Wallet', 'gojek makan siang', '2022-12-22'),
	('1000000000004', '1000000000005', 50000, 'Wallet', 'Tiket Bioskop', '2023-01-12'),
	('1000000000005', '1000000000007', 1000000, 'Wallet', 'bayar utang', '2023-01-02'),
	('1000000000004', '1000000000008', 50000, 'Wallet', 'morning coffee', '2022-12-02'),
	('1000000000004', '1000000000007', 500000, 'Wallet', '-', '2024-03-10'),
	('1000000000005', '1000000000004', 50000, 'Wallet', '-', '2024-03-13'),
	('1000000000004', '1000000000006', 500000, 'Wallet', '-', '2024-03-17');
	
	
update wallets 
set balance = balance - (50000000 + 5000000 + 5000000 + 1000000)
where user_id = 1;

update wallets 
set balance = balance - (5000000 + 5000000)
where user_id = 2;

update wallets 
set balance = balance + (50000000 + 30000 - 20000 - 5000 + 500000 + 1000000 - 300000 + 100000 - 70000 - 50000 - 50000 - 500000 + 50000 - 500000),
	gacha_trial = gacha_trial + 5
where user_id = 4;

update wallets 
set balance = balance + (5000000 - 500000 + 100000 + 50000 - 1000000)
where user_id = 5;

update wallets 
set balance = balance + (5000000 + 20000 - 500000 - 50000 + 5000000 - 100000)
where user_id = 6;

update wallets 
set balance = balance + (5000000 + 500000 + 50000 - 100000 + 70000 + 1000000)
where user_id = 7;

update wallets 
set balance = balance + (5000000 - 30000 + 300000 - 5000000 + 50000)
where user_id = 8;

update wallets 
set balance = 100000000
where user_id = 4;

insert into gacha_prizes (amount, created_at)
values
	(50000000, NOW()),
	(5000, NOW()),
	(0, NOW()),
	(13000000, NOW()),
	(100000, NOW()),
	(0, NOW()),
	(20000000, NOW()),
	(1000000000, NOW()),
	(0, NOW());