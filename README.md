# AdGuard Automator

This tool provides a command-line interface to manage DNS rewrites via the AdGuard API. It allows users to add or remove DNS rewrite rules easily using basic authentication.

## Prerequisites

Before you run this tool, make sure you have the following installed:

- Access to an AdGuard server with the API enabled

## Installation

### Option 1

Download the binary.

### Option 2

1. **Clone the repository**

   ```bash
   git clone https://github.com/Mikej81/adguard_automator.git
   cd adguard_automator
   go build -o ./agauto .
   ```

## Usage

### Add Rewrite

```bash
./agauto --add --url <AdGuard_URL> --username <Username> --password <Password> --domain <Domain> --value <IP_or_Domain>
```

Example

```bash
./dns-rewrite-tool --add --url http://example.com/control --username admin --password secret --domain example.com --value 192.168.1.1

```

### Remove Rewrite

```bash
./agauto --remove --url <AdGuard_URL> --username <Username> --password <Password> --domain <Domain> --value <IP_or_Domain>
```

Example

```bash
./agauto --remove --url http://example.com/control --username admin --password secret --domain example.com --value 192.168.1.1

```

## Limitations

Today the tool only handles DNS Rewrites, I will add more as I need it.
