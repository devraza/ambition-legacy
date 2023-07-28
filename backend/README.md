# Backend Usage

In order to run the backend, you need to have [sqlite3](https://sqlite.org)
installed on your system. Once you do, make a database file named `users.db`
and initialize it with the `sql/init.sql` script:

```sh
$ cat sql/init.sql | sqlite3 users.db
```

You also need to create a `.env` file with the following variables:

- `JWT_SECRET`: Required. A cryptographically secure string used to encode
tokens.
- `PORT`: Optional. Overrides the default port of `7741`

