SQLite format 3   @     1   	                                                            1 .[5� 	���F_	��
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          �!�!tableitemsitemsCREATE TABLE items(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		created_at DATETIME,
		title STRING,
		content TEXT NULL,
		price INTEGER
		, photo_url STRING NULL, category_first STRING NULL, category_second STRING NULL, category_third STRING NULL)�@�Stablemessagesmessages	CREATE TABLE messages(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id STRING,
		created_at DATETIME,
		group_id STRING,
		user_name STRING NULL)�?!!�ItablechatgroupschatgroupsCREATE TABLE chatgroups(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		created_at DATETIME,
		chat_member STRING NULL,
		chat_name STRING NULL)� �                                                                                                                                                                                                           �3�9tablesessionssessionsCREATE TABLE sessions(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id STRING,
		created_at DATETIME)/C indexsqlite_autoindex_sessions_1sessionsP++Ytablesqlite_sequencesqlite_sequenceCREATE TABLE sqlite_sequence(name,seq)�2�CtableusersusersCREATE TABLE users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		created_at DATETIME,
		name STRING,
		nick_name STRING NULL,
		email STRING,
		password STRING,
		icon_url STRING NULL,
		phone STRING NULL,
		address STRING NULL,
		birthday STRING NULL
		))= indexsqlite_autoindex_users_1users          � t�                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            �	 UM/]0069d704-1d40-11ed-b9c4-acde480011222022-08-16 17:47:04.040338+09:00user2@example.com9bc34549d565d9505b287de0cd20ac77be1d3f2c�	 UM/]d3a21790-1b91-11ed-aa85-acde480011222022-08-14 14:27:45.317801+09:00user1@example.com9bc34549d565d9505b287de0cd20ac77be1d3f2c
   � ��                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   (U0069d704-1d40-11ed-b9c4-acde48001122'U	d3a21790-1b91-11ed-aa85-acde48001122� � �����                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  	messages!	chatgroupssessions	items   	session	usersF � ���                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             \ U/M47f3079a-1d71-11ed-9c0a-acde48001122user2@example.com2022-08-16 23:39:49.397408+09:00[ U/	M56a3b1dc-1c4e-11ed-bc50-acde48001122user1@example.com2022-08-15 12:57:10.495015+09:00   ]U/	Md21ebf7a-1b9d-11ed-8463-acde48001122user1@example.com2022-08-14 15:53:36.740039+09:00[ U/	Mb6c59a4e-1d43-11ed-966a-acde48001122user1@example.com2022-08-16 18:13:38.477232+09:00
� ] ��]                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               (Ub6c59a4e-1d43-11ed-966a-acde48001122(U56a3b1dc-1c4e-11ed-bc50-acde48001122   (d21ebf7a-1b9d-11ed-8463-acde48001122(U47f3079a-1d71-11ed-9c0a-acde48001122� 
� b�
�t
�
�
�
�
�
�
�                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       �] 	M++�-	2022-08-16 23:23:56.877155+09:00ユーザー１ユーザー１'https://firebasestorage.googleapis.com/v0/b/prac-ec.appspot.com/o/images%2Fnetwork-town.jpg?alt=media&token=369d07ae-b1cb-41ce-9618-4bda5e585334�G	 M�-2022-08-16 18:16:15.885548+09:00user1user1 �Phttps://firebasestorage.googleapis.com/v0/b/prac-ec.appspot.com/o/images%2Fparttime-bar.jpg?alt=media&token=cfe29413-661a-47d9-843f-08ccc8ea3db2( �M+2022-08-16 17:54:48.386077+09:00ユーザー２ああ�(YMII2022-08-16 17:51:02.212688+09:00ユーザー２のテスト２ユーザー２の出品です'( �K=I2022-08-16 17:48:07.03477+09:00ユーザー２�k M5=�+	2022-08-17 01:39:36.345354+09:00ユーザー２　PCユーザー２の出品xhttps://firebasestorage.googleapis.com/v0/b/prac-ec.appspot.com/o/images%2Finternet-pc.jpg?alt=media&token=4e525af3-a934-4102-adf2-cb71eb3a2506�C	 M�+2022-08-16 16:14:43.619739+09:00testtest�https://firebasestorage.googleapis.com/v0/b/prac-ec.appspot.com/o/images%2Fglass-human.jpg?alt=media&token=10c26ab3-58b9-4afc-9192-4cb0e95d10b3   �M-2022-08-16 03:46:15.498957+09:00testcontent'parttime-bar.jpg   �M32022-08-16 01:56:38.802985+09:00testtest �Pinternet-tablet.jpg   GM+2022-08-16 01:52:57.420875+09:00testtest
��internet-pc.jpg   � �                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              6 M	+2022-08-17 01:21:34.429347+09:00ユーザー１   � �                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            X mM	商品が届くのを楽しみにしています2022-08-17 01:27:07.818938+09:00