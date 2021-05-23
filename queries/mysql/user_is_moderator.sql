SELECT EXISTS(
               SELECT *
               FROM tokens
               WHERE token = ?
                 AND role IN ('admin', 'moderator')
           )