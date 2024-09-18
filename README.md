# Simple WebUI for Analyzing Logs

It should allow to easily parse, vizualize and analyze logs of any kind.

This project in in an early stage.

----

## Contribute

Feel free to contribute by:
* reporting issues
* taking part in discussions

PRs welcome.

----

## Config

See: [Example Config](https://github.com/O-X-L/webui-log-analysis/blob/main/config.yml)

You can define:
* Multiple files
* Extract multiple values from every single line
* Add additional information using processors (for example GeoIP information)

----

## Extraction

### Name

The variable-name that should be used.

### Search

You can use the `search` attribute to easily search for values inside the log-lines.

It is handled like a simplified regex search.

We will extract the value represented by the `<THIS>` placeholder.

### Regex

If the `search` isn't enough for you - you can use basic regex patterns. For testing see: [regex101.com](https://regex101.com)

The first match-group `()` will be used as value. Maybe use non-greedy matching: `(.*?)`

### Type

Some type-casing - especially useful for `int` or `float` - as we need numeric values for some vizualizations.

### Clean

* **Remove**: The string will be removed after extraction.

### Multiple

Set `multiple: true` to extract an array of matched values.

### Fallback

Useful if type-casting fails. You can define a fallback value.

----

## Process

### Name/Search/Regex

See Extration

### From

The source-variable to transform.

### GeoIP

The kind of GeoIP database to use for the lookup.

You can optionally override the attribute you want to lookup.

#### IPInfo

**Country Attributes**

* StartIp
* EndIp
* Country
* CountryName
* Continent
* ContinentName
  
**ASN Attributes**

* StartIp
* EndIp
* ASN
* Name
* Domain

**Country & ASN Attributes**

* StartIp
* EndIp
* Country
* CountryName
* Continent
* ContinentName
* ASN
* ASName
* ASDomain

**City Attributes**

* StartIp
* EndIp
* JoinKey
* City
* Region
* Country
* Latitude
* Longitude
* PostalCode
* Timezone

**Privacy Attributes**

* StartIp
* EndIp
* JoinKey
* Hosting
* Proxy
* Tor
* Vpn
* Relay
* Service

#### MaxMind

**Country Attributes**

* Country.Code
* Country.Id
* Country.EuopeanUnion
* Continent.Code
* Continent.Id
* Continent.Names
  
**ASN Attributes**

* ASN
* Name

**City Attributes**

* City.Code
* City.Id
* Location.AccuracyRadius
* Location.Latitude
* Location.Longitude
* Location.Timezone
* Postal.Code
* Traits.IsAnycast
* Traits.IsAnonymousProxy

+ Country
