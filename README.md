# Mini-Scan

Hello!

As you've heard by now, Censys scans the internet at an incredible scale. Processing the results necessitates scaling horizontally across thousands of machines. One key aspect of our architecture is the use of distributed queues to pass data between machines.

---

The `docker-compose.yml` file sets up a toy example of a scanner. It spins up a Google Pub/Sub emulator, creates a topic and subscription, and publishes scan results to the topic. It can be run via `docker compose up`.

Your job is to build the data processing side. It should:
1. Pull scan results from the subscription `scan-sub`.
2. Maintain an up-to-date record of each unique `(ip, port, service)`. This should contain when the service was last scanned and a string containing the service's response.

> **_NOTE_**
The scanner can publish data in two formats, shown below. In both of the following examples, the service response should be stored as: `"hello world"`.
> ```javascript
> {
>   // ...
>   "data_version": 1,
>   "data": {
>     "response_bytes_utf8": "aGVsbG8gd29ybGQ="
>   }
> }
>
> {
>   // ...
>   "data_version": 2,
>   "data": {
>     "response_str": "hello world"
>   }
> }
> ```

Your processing application should be able to be scaled horizontally, but this isn't something you need to actually do. The processing application should use `at-least-once` semantics where ever applicable.

You may write this in any languages you choose, but Go, Scala, or Rust would be preferred. You may use any data store of your choosing, with `sqlite` being one example.

--- 

Please upload the code to a publicly accessible GitHub, GitLab or other public code repository account.  This README file should be updated, briefly documenting your solution. Like our own code, we expect testing instructions: whether it’s an automated test framework, or simple manual steps.

To help set expectations, we believe you should aim to take no more than 4 hours on this task.

We understand that you have other responsibilities, so if you think you’ll need more than 5 business days, just let us know when you expect to send a reply.

Please don’t hesitate to ask any follow-up questions for clarification.

---

## To build and run

```bash
make build
make run
```

## To run tests

```bash
make test
```
## To do
- [ ] Add a db migrarion

## DB
```bash
make db-exec
```

After the app is running, you can run the following command to create the table
db migration
```bash
make db-migrate
```

To exec into the db container
```bash
make db-exec
```

table
```sql
select * from scans;
```

```sql
mini_scan=# select * from scan limit 5;
    ip     | port  | service |              data              | timestamp  |          created_at           |          updated_at           
-----------+-------+---------+--------------------------------+------------+-------------------------------+-------------------------------
 1.1.1.32  | 50083 | DNS     | "c2VydmljZSByZXNwb25zZTogMTc=" | 1733862724 | 2024-12-10 20:32:04.875386+00 | 2024-12-10 20:32:04.849501+00
 1.1.1.94  | 20333 | SSH     | "c2VydmljZSByZXNwb25zZTogOTc=" | 1733862725 | 2024-12-10 20:32:05.832487+00 | 2024-12-10 20:32:05.832147+00
 1.1.1.198 | 25332 | HTTP    | service response: 70           | 1733862726 | 2024-12-10 20:32:06.807043+00 | 2024-12-10 20:32:06.806723+00
 1.1.1.81  | 47600 | SSH     | service response: 92           | 1733862727 | 2024-12-10 20:32:07.84021+00  | 2024-12-10 20:32:07.839839+00
 1.1.1.114 | 48560 | SSH     | "c2VydmljZSByZXNwb25zZTogMzY=" | 1733862728 | 2024-12-10 20:32:08.814121+00 | 2024-12-10 20:32:08.813847+00
(5 rows)
```

## To improve
- [ ] Use worker pool and channel to process messages
- [ ] Add audit table to track scan changes

