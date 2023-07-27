# Backend Usage

In order to run the backend, you need to have [sqlite3](https://sqlite.org)
installed on your system. Once you do, make a database file named `users.db`
and initialize it with the `sql/init.sql` script:

```sh
$ cat sql/init.sql | sqlite3 users.db
```

You can optionally provide the `PORT` environment variable to override the
default port of `7741`
