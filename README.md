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
    ip     | port  | service |              data              | timestamp  |          created_at           |          updated_at           
-----------+-------+---------+--------------------------------+------------+-------------------------------+-------------------------------
 1.1.1.252 | 25152 | DNS     | service response: 21 | 1733865528 | 2024-12-10 21:18:49.780913+00 | 2024-12-10 21:18:49.776484+00
 1.1.1.250 | 29044 | SSH     | service response: 50 | 1733865529 | 2024-12-10 21:18:49.937958+00 | 2024-12-10 21:18:49.93748+00
 1.1.1.146 |   643 | DNS     | service response: 85 | 1733865530 | 2024-12-10 21:18:50.984518+00 | 2024-12-10 21:18:50.984088+00
 1.1.1.101 | 39677 | DNS     | service response: 19 | 1733865531 | 2024-12-10 21:18:51.954938+00 | 2024-12-10 21:18:51.954251+00
 1.1.1.13  | 47331 | SSH     | service response: 99 | 1733865532 | 2024-12-10 21:18:52.96854+00  | 2024-12-10 21:18:52.967805+00
 1.1.1.92  |  9466 | DNS     | service response: 24 | 1733865533 | 2024-12-10 21:18:53.956561+00 | 2024-12-10 21:18:53.955389+00
```

## To improve
- [ ] Use worker pool and channel to process messages
- [ ] Add audit table to track scan changes

