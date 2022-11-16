USE GoBlog;

DELIMITER //

DROP PROCEDURE IF EXISTS InsertNewBlogPost;

CREATE PROCEDURE InsertNewBlogPost(IN p_author TEXT,
                                   IN p_datePosted TEXT,
                                   IN p_title TEXT,
                                   IN p_body TEXT)
BEGIN
  INSERT INTO BlogPosts (Author, DatePosted, Title, Body)
  VALUES (p_author, p_datePosted, p_title, p_body);
  
  SELECT LAST_INSERT_ID();
END //

DROP PROCEDURE IF EXISTS UpdateBlogPost;

CREATE PROCEDURE UpdateBlogPost(IN p_id INT,
                                IN p_dateLastUpdated TEXT,
                                IN p_title TEXT,
                                IN p_body TEXT)
BEGIN
  UPDATE BlogPosts
  SET DateLastUpdated = p_dateLastUpdated,
      Title = p_title,
      Body = p_body
  WHERE Id = p_id;
END //

DELIMITER ;