# Investing.com

Get the economic calendar from investing.com

Run it

    go run investingcom.go

## Notes

Minimal example

    curl 'http://m.investing.com/economic-calendar/services/filter/' \
        -H 'X-Requested-With: XMLHttpRequest' \
        --data 'timeZone=8&country[]=25&country[]=32&country[]=6&country[]=37&country[]=72&country[]=22&country[]=17&country[]=39&country[]=14&country[]=10&country[]=35&country[]=43&country[]=36&country[]=110&country[]=11&country[]=26&country[]=12&country[]=4&country[]=5&country[]=56&currentTab=today'

## License

The MIT License (MIT)

Copyright (c) 2016 Scott Barr

See [LICENCE.md](LICENCE.md)
