[![GitHub version](https://img.shields.io/github/release/jeffalyanak/check_godaddy.svg)](https://github.com/jeffalyanak/check_godaddy/releases/latest)
[![License](https://img.shields.io/github/license/jeffalyanak/check_godaddy.svg)](https://github.com/jeffalyanak/check_godaddy/blob/master/LICENSE.txt)
[![Donate](https://img.shields.io/badge/donate--green.svg)](https://jeff.alyanak.ca/donate)
[![Matrix](https://img.shields.io/matrix/check_godaddy:social.rights.ninja.svg)](https://matrix.to/#/#check_godaddy:social.rights.ninja)

# GoDaddy Domain Expiry Checker

Icinga/Nagios plugin, checks the domain expiry status using the GoDaddy API.

User configurable `warning` and `critical` levels

## Installation and requirements

* Golang 1.13.8


## Usage

```bash
Usage of ./godaddy-check:
  -crit int
        days until critical (default 7)
  -domain string
        domain to search
  -key string
        API Key
  -secret string
        API Secret
  -warn int
        days until warning (default 15)
```

## License

GoDaddy Domain Expiry Checker is licensed under the terms of the GNU General Public License Version 3.