CREATE TABLE IF NOT EXISTS User (
    UserID INTEGER PRIMARY KEY AUTOINCREMENT,
    Username TEXT NOT NULL UNIQUE,
    Email TEXT NOT NULL,
    PasswordHash TEXT NOT NULL,
    RegistrationDate DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS Post (
    PostID INTEGER PRIMARY KEY AUTOINCREMENT,
    UserID INTEGER,
    Title TEXT NOT NULL,
    Content TEXT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Category TEXT NOT NULL,
    LastModifiedDate DATE,
    CommentCount INTEGER DEFAULT 0,
    LikeCount INTEGER DEFAULT 0,  -- Added to track likes
    DislikeCount INTEGER DEFAULT 0,  -- Added to track dislikes
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

CREATE TABLE IF NOT EXISTS Comments (
    CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
    PostID INTEGER,
    UserID INTEGER,
    Content TEXT NOT NULL,
    CreatedAt DATETIME NOT NULL,
    LastModifiedDate DATE,
    LikeCount INTEGER DEFAULT 0,  -- Added to track likes
    DislikeCount INTEGER DEFAULT 0,  -- Added to track dislikes
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

CREATE TABLE IF NOT EXISTS LikeDislike (
    LikeDislikeID INTEGER PRIMARY KEY AUTOINCREMENT,
    UserID INTEGER,
    PostID INTEGER,
    CommentID INTEGER,
    IsLike BOOLEAN NULL,
    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (CommentID) REFERENCES Comments(CommentID),
    CHECK ((PostID IS NOT NULL AND CommentID IS NULL) OR (PostID IS NULL AND CommentID IS NOT NULL))
);


CREATE TABLE IF NOT EXISTS Session (
    SessionID TEXT PRIMARY KEY,
    UserID INTEGER,
    CreatedAt DATETIME NOT NULL,
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);
