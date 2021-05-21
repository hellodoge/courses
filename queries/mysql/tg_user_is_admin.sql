SELECT EXISTS(
               SELECT *
               FROM tokens
               WHERE tokens.role = 'admin'
                 AND tokens.token = (
                   SELECT tg_users.token
                   FROM tg_users
                   WHERE tg_users.id = ?
               )
           );