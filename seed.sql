-- Insert users
INSERT INTO users (username, password, email) VALUES ('John Doe', 'p@ssw0rd1', 'john@example.com');
INSERT INTO users (username, password, email) VALUES ('Jane Smith', 'p@ssw0rd1', 'jane@example.com');
INSERT INTO users (username, password, email) VALUES ('Michael Johnson', 'p@ssw0rd1', 'michael@example.com');
INSERT INTO users (username, password, email) VALUES ('Emily Davis', 'p@ssw0rd1', 'emily@example.com');
INSERT INTO users (username, password, email) VALUES ('David Brown', 'p@ssw0rd1', 'david@example.com');
INSERT INTO users (username, password, email) VALUES ('Sarah Miller', 'p@ssw0rd1', 'sarah@example.com');
INSERT INTO users (username, password, email) VALUES ('Daniel Wilson', 'p@ssw0rd1', 'daniel@example.com');
INSERT INTO users (username, password, email) VALUES ('Olivia Taylor', 'p@ssw0rd1', 'olivia@example.com');
INSERT INTO users (username, password, email) VALUES ('James Anderson', 'p@ssw0rd1', 'james@example.com');
INSERT INTO users (username, password, email) VALUES ('Sophia Thomas', 'p@ssw0rd1', 'sophia@example.com');

-- Insert securities
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('CS-1', 1, 100, 'Stock', 50.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('CS-2', 1, 200, 'Stock', 75.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('CS-3', 1, 150, 'Stock', 60.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('CS-4', 1, 300, 'Stock', 55.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('OPT-1', 2, 50, 'Option', 10.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('OPT-2', 3, 75, 'Option', 15.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('OPT-3', 3, 100, 'Option', 12.50);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('OPT-4', 4, 120, 'Option', 20.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('OPT-5', 5, 90, 'Option', 18.00);
INSERT INTO securities (name, user_id, amount, type, price) VALUES ('OPT-6', 5, 80, 'Option', 16.50);

-- Insert exercise requests
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (1, 500, '2022-01-15', 'Pending');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (1, 300, '2022-02-20', 'Completed');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (2, 200, '2022-03-10', 'Pending');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (3, 400, '2022-04-05', 'Completed');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (4, 150, '2022-05-12', 'Pending');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (5, 100, '2022-06-18', 'Completed');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (6, 250, '2022-07-25', 'Pending');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (7, 175, '2022-08-30', 'Completed');
INSERT INTO exercise_requests (security_id, amount, exercise_date, status) VALUES (8, 300, '2022-09-05', 'Pending');