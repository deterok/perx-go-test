## What this
This is the solution of a test task at https://github.com/hollizzy/perx-go-test.

Text of the task can be founded [there](TASK.md).

## Dependencies
- go-1.*
- docker-18.*
- make-4.2.*

## How to run
- Simple run:
    ```shell
    $ make up
    ```
    or
    ```shell
    $ make up-build
    ```

    After this command, the environment will start and you can request the server at `localhost:8080`

- Run tests coverage:

    WARNING: The Command is not optimized. Extra containers can be running.
    ```shell
    $ make test_cover
    ```
    Out:
    ```
    ./docker/run.sh core go test -timeout 30s -v -cover ./...
    Starting perx-go-test_postgres_1 ... done
    ?       perx-go-test    [no test files]
    ?       perx-go-test/lib        [no test files]
    ?       perx-go-test/models     [no test files]
    === RUN   TestNewCodesResource
    --- PASS: TestNewCodesResource (0.00s)
    === RUN   TestCodesResource_Create
    === RUN   TestCodesResource_Create/simple
    === RUN   TestCodesResource_Create/no_free_values
    --- PASS: TestCodesResource_Create (0.18s)
        --- PASS: TestCodesResource_Create/simple (0.00s)
        --- PASS: TestCodesResource_Create/no_free_values (0.18s)
    === RUN   TestCodesResource_generateString
    === RUN   TestCodesResource_generateString/one_symbol
    === RUN   TestCodesResource_generateString/multiple_symbols
    --- PASS: TestCodesResource_generateString (0.00s)
        --- PASS: TestCodesResource_generateString/one_symbol (0.00s)
        --- PASS: TestCodesResource_generateString/multiple_symbols (0.00s)
    === RUN   TestCodesResource_Check
    === RUN   TestCodesResource_Check/simple
    === RUN   TestCodesResource_Check/code_already_used
    === RUN   TestCodesResource_Check/code_not_found
    --- PASS: TestCodesResource_Check (0.00s)
        --- PASS: TestCodesResource_Check/simple (0.00s)
        --- PASS: TestCodesResource_Check/code_already_used (0.00s)
        --- PASS: TestCodesResource_Check/code_not_found (0.00s)
    PASS
    coverage: 81.1% of statements
    ok      perx-go-test/resources  0.194s  coverage: 81.1% of statements
    ?       perx-go-test/views      [no test files]
    ```

- To cleanup, run the following command:

    WARNING: It's full cleanup! It can delete important containers (like postgres)!
    ```shell
    make down
    ```
## API
You can view the API [there](https://editor.swagger.io/?url=https://raw.githubusercontent.com/deterok/perx-go-test/master/openapi.yaml).


## Known bugs
Echo server writes status code before encoding a JSON (10.05.2019, Open)
https://github.com/labstack/echo/issues/1334
