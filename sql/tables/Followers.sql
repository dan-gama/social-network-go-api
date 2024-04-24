CREATE TABLE Followers(
    userId      int not null,
    followerId  int not null,
    FOREIGN KEY (userId) REFERENCES User(Id) ON DELETE CASCADE,
    FOREIGN KEY (followerId) REFERENCES User(Id) ON DELETE CASCADE,
    PRIMARY KEY(userId, followerId)
) ENGINE=INNODB