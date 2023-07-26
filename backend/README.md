# Backend Usage

In order for this to work, you need to have [sqlite3](https://sqlite.org)
installed on your system. Once you do, make a databse called `users.db`
and initialize it with the `sql/init.sql` script:

```sh
$ cat sql/init.sql | sqlite3 users.db
```

You also need to have a `.env` file in this folder, with the following options specified

- JWT_SECRET: Used to encrypt tokens for user auth. Must be provided. Should
be cryptographically secure
- PORT: Optionally replace the default port of 7741