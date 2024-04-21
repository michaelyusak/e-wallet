drop database if exists e_wallet_db;
create database e_wallet_db;

\c e_wallet_db;

drop table if exists users, wallets, transactions, reset_password_tokens, gacha_prizes, user_gachas;
drop sequence if exists wallet_number_seq cascade;

create sequence wallet_number_seq;

create table users (
	user_id BIGSERIAL primary key,
	user_email VARCHAR not null,
	user_name VARCHAR not null,
	user_profile_picture_name VARCHAR,
	user_password VARCHAR not null,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null default NOW(),
	deleted_at TIMESTAMP
);

create table wallets (
	wallet_id BIGSERIAL primary key,
	user_id BIGINT not null,
	wallet_number VARCHAR(13) not null DEFAULT CONCAT('100', lpad(nextval('wallet_number_seq')::text, 10, '0')),
	balance DECIMAL not null default 0,
	gacha_trial INT not null,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null default NOW(),
	deleted_at TIMESTAMP
);

create table transactions (
	transaction_id BIGSERIAL primary key,
	sender_wallet_number VARCHAR(13) not null,
	recepient_wallet_number VARCHAR(13) not null,
	amount DECIMAL not null,
	source_of_fund VARCHAR not null,
	description text not null,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null default NOW(),
	deleted_at TIMESTAMP
);

create table reset_password_tokens (
	token_id BIGSERIAL primary key,
	user_id BIGINT not null,
	reset_password_token VARCHAR not null,
	expired_at TIMESTAMP not null,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null default NOW(),
	deleted_at TIMESTAMP
);

create table gacha_prizes (
	prize_id BIGSERIAL primary key,
	amount DECIMAL not null,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null default NOW(),
	deleted_at TIMESTAMP
);

create table user_gachas (
	user_gacha_id BIGSERIAL primary key,
	wallet_id BIGINT,
	prize_id BIGINT,
	created_at TIMESTAMP not null,
	updated_at TIMESTAMP not null default NOW(),
	deleted_at TIMESTAMP
);

select * from transactions t;

select * from reset_password_tokens rpt ;

select * from users;
select * from wallets;
select * from transactions t ;

update users set user_profile_picture_path = '' where user_id = 4;

select * from transactions t
where (sender_wallet_number = '1000000000004'
	or recepient_wallet_number = '1000000000004')
order by created_at
desc ;

select sender_wallet_number, recepient_wallet_number, u1.user_name as sender_name , u2.user_name as recepient_name from transactions t
JOIN wallets w1 ON t.sender_wallet_number = w1.wallet_number
JOIN users u1 ON w1.user_id = u1.user_id
JOIN wallets w2 ON t.recepient_wallet_number = w2.wallet_number
join users u2 on w2.user_id = u2.user_id ;

SELECT wallet_id, wallet_number, balance, sum(t1.amount) as income, sum(t2.amount) as expense, gacha_trial
		FROM wallets w
		join transactions t1 on w.wallet_number = t1.recepient_wallet_number
		join transactions t2 on w.wallet_number = t2.sender_wallet_number 
		WHERE user_id = 4
		AND w.deleted_at
			IS null
		group by wallet_id ;

select transaction_id, created_at, count(*) over()
from transactions t 
where description 
ilike '%%'
and (sender_wallet_number = '1000000000004'
or recepient_wallet_number = '1000000000004')
and created_at between '2024-03-11' and '2024-03-18 23:59'
limit 15 offset 0;

SELECT COUNT(*) OVER(), transaction_id, sender_wallet_number, u1.user_name, recepient_wallet_number, u2.user_name, amount, source_of_fund, description, t.created_at 
FROM transactions t 
	JOIN wallets w1 
		ON t.sender_wallet_number = w1.wallet_number 
	JOIN users u1 
		ON w1.user_id = u1.user_id 
	JOIN wallets w2 
		ON t.recepient_wallet_number = w2.wallet_number 
	JOIN users u2 
		ON w2.user_id = u2.user_id 
WHERE description 
	ILIKE '%%' 
		AND (sender_wallet_number = '1000000000004' 
			OR recepient_wallet_number = '1000000000004') 
		AND t.created_at 
			BETWEEN '2019-03-11' 
				AND '2024-03-18 23:59' 
ORDER 
	BY created_at 
		desc 
LIMIT 21 
OFFSET 0;

 SELECT
 w.wallet_id,
 w.user_id,
 w.wallet_number,
 w.balance,
 COALESCE(income.total_income, 0) AS income,
 COALESCE(expense.total_expense, 0) AS expense
FROM
 wallets w
LEFT JOIN (
 SELECT
 recepient_wallet_number,
 SUM(amount) AS total_income
 FROM
 transactions
 WHERE
 recepient_wallet_number = (select wallet_number from wallets w2 where user_id = 4)
 GROUP BY
 recepient_wallet_number
) AS income ON w.wallet_number = income.recepient_wallet_number
LEFT JOIN (
 SELECT
 sender_wallet_number,
 SUM(amount) AS total_expense
 FROM
 transactions
 WHERE
 sender_wallet_number = (select wallet_number from wallets w3 where user_id = 4)
 GROUP BY
 sender_wallet_number
) AS expense ON w.wallet_number = expense.sender_wallet_number
WHERE
 w.user_id = 4;

SELECT w.wallet_id, w.user_id, w.wallet_number, w.balance, COALESCE(income.total_income, 0) AS income, COALESCE(expense.total_expense, 0) AS expense
                FROM wallets w
                JOIN (
                SELECT recepient_wallet_number, SUM(amount) AS total_income
                        FROM transactions
                        GROUP 
                                        BY recepient_wallet_number) 
                        AS income 
                                ON w.wallet_number = income.recepient_wallet_number
                JOIN (
                SELECT sender_wallet_number, SUM(amount) AS total_expense
                        FROM transactions
                        GROUP 
                                        BY sender_wallet_number)
                        AS expense 
                                ON w.wallet_number = expense.sender_wallet_number
                JOIN users u 
                        ON w.user_id = u.user_id
                                WHERE u.user_id = 9

SELECT recepient_wallet_number, SUM(amount) AS total_income
                        FROM transactions
                        GROUP 
                                        BY recepient_wallet_number

                                        
SELECT COALESCE(income.total_income, 0) AS income, COALESCE(expense.total_expense, 0) AS expense
		FROM wallets w
		left JOIN (
    		SELECT recepient_wallet_number, SUM(amount) AS total_income
    			FROM transactions
    			GROUP 
					BY recepient_wallet_number) 
			AS income 
				ON w.wallet_number = income.recepient_wallet_number
		left JOIN (
    		SELECT sender_wallet_number, SUM(amount) AS total_expense
    			FROM transactions
    			GROUP 
					BY sender_wallet_number)
			AS expense 
				ON w.wallet_number = expense.sender_wallet_number
		WHERE w.wallet_id = 9;
                                        
         
	SELECT recepient_wallet_number, SUM(amount) AS total_income
    			FROM transactions
    			GROUP 
					BY recepient_wallet_number;
                                        
					SELECT sender_wallet_number, SUM(amount) AS total_expense
    			FROM transactions
    			GROUP 
					BY sender_wallet_number;
                                        
                                        
                                        
                                        
                                        

