[![GitHub version](https://img.shields.io/github/v/release/jeffalyanak/check_godaddy)](https://github.com/jeffalyanak/check_godaddy/releases/latest)
[![License](https://img.shields.io/github/license/jeffalyanak/check_godaddy)](https://github.com/jeffalyanak/check_godaddy/blob/master/LICENSE)
[![Donate](https://img.shields.io/badge/donate--green)](https://jeff.alyanak.ca/donate)
[![Matrix](https://img.shields.io/badge/chat--green)](https://matrix.to/#/#check_godaddy:social.rights.ninja)

# GoDaddy Domain Expiry Checker

Icinga/Nagios plugin, checks the domain expiry status using the GoDaddy API.

User configurable `warning` and `critical` levels

## Installation and requirements

* Golang 1.13.8


## Usage

Requires [GoDaddy API keys](https://developer.godaddy.com/).

```bash
Usage of check_godaddy:
  -domain string
        domain to search
  -warn int
        days until warning (default 15)
  -crit int
        days until critical (default 7)
  -key string
        API Key
  -secret string
        API Secret
```

Example:

```bash
check_godaddy -domain example.com -warn 30 -crit 14 -key 1234567890 -secret 123456
```

## License

GoDaddy Domain Expiry Checker is licensed under the terms of the GNU General Public License Version 3.