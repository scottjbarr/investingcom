# Investing.com

A Go client for https://investing.com .

This client currently only implements the history endpoint.

## CLI Utilities

### investingcom-history

CLI utility to get history in CSV format.

```
$ go run cmd/investingcom-history --help
open,high,low,close,volume
1.0441,1.0489,1.0382,1.0482,83871
1.0481,1.0487,1.0366,1.0428,68085
1.0424,1.0464,1.0417,1.0421,68801
...
```

### investingcom-pairs

CLI utility to get a list of all pair ID's.

```
$ go run cmd/investingcom-pairs
1,EUR/USD
2,GBP/USD
3,USD/JPY
4,USD/CHF
...
```

## References

- [Investing.com Unofficial APIs](https://github.com/DavideViolante/investing-com-api)

## License

Copyright (c) 2022 Scott Barr

All rights reserved.
