# plantai

Plantai is personal house plant tracking system, with an embedded encyclopedia.

This application provides the server that will serve data for frontend clients,
made for many platforms.

## Developent

To install plantai, you must have Golang >= 1.14 installed at your environent.
Then, run the following:

```sh
$ make install
```

You can check the program usage, using `plantai help`:

```
$ plantai help
plantai targets Cachalote Wix ecommerce, managing customer queues for products

Usage:
  plantai [command]

Available Commands:
  help        Help about any command
  serve       start plantai in server mode

Flags:
  -h, --help   help for plantai

Use "plantai [command] --help" for more information about a command.
```
