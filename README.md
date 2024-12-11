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

## Run Tests

```bash
make test
```

## Build and Run

```bash
make build
make run
```

## Database 
Run migrations after app is running
```bash
make db-migrate
```

Exec into the db container
```bash
make db-exec
```

Example data
```sql
    ip     | port  | service |         data         | timestamp  |          created_at           |          updated_at           
-----------+-------+---------+----------------------+------------+-------------------------------+-------------------------------
 1.1.1.170 | 33234 | SSH     | service response: 36 | 1733868149 | 2024-12-10 22:02:29.271887+00 | 2024-12-10 22:02:29.270444+00
 1.1.1.198 |   514 | SSH     | service response: 51 | 1733868150 | 2024-12-10 22:02:30.26541+00  | 2024-12-10 22:02:30.264492+00
 1.1.1.61  | 27374 | SSH     | service response: 41 | 1733868151 | 2024-12-10 22:02:31.251701+00 | 2024-12-10 22:02:31.250794+00
 1.1.1.32  | 25635 | HTTP    | service response: 60 | 1733868152 | 2024-12-10 22:02:32.243768+00 | 2024-12-10 22:02:32.242632+00
 1.1.1.138 | 37493 | DNS     | service response: 51 | 1733868153 | 2024-12-10 22:02:33.252666+00 | 2024-12-10 22:02:33.252266+00
```

## To improve
- [ ] Use worker pool and channel to process messages
- [ ] Add a DAO layer
- [ ] Add integration tests
- [ ] Add end to end tests
- [ ] Add audit table to track scan changes
- [ ] Add RESTful API to retrive and update data
- [ ] Add Metrics to track and alert on processing errors, latency, health, etc.
- [ ] Add structured logging
- [ ] Add graceful shutdown
- [ ] Add healthcheck endpoint
- [ ] Add retry logic for processing messages
  - [ ] Add backoff logic
  - [ ] Add max retries
