# SecretBin CLI
This CLI allows for automatic secret creation in [SecretBin](https://github.com/Nihility-io/SecretBin). Note however that the CLI currently only support creating AES256-GCM secrets. XChaCha20 is currently not supported.

## Usage
```
SecretBin Command Line Interface

Usage:
  secretbin [command]

Available Commands:
  completion        Generate the autocompletion script for the specified shell
  create            Create a new secret using SecretBin
  generate-password Generates a secure password
  help              Help about any command
  info              Print information about the SecretBin server
  set-endpoint      Set where SecretBin is hosted

Flags:
  -h, --help          help for secretbin-cli
      --hide-banner   Do not show the SecretBin banner message. This is useful for scripting.
  -v, --version       version for secretbin-cli

Use "secretbin-cli [command] --help" for more information about a command.



Create a new secret using SecretBin

Usage:
  secretbin create [message] [flags]

Flags:
  -a, --attachment strings   Attachment files to the secret (can be specified multiple times)
  -b, --burn-after uint      Times the secret can be read before being deleted (0 means no burn after reading)
  -x, --expires string       Expiration time for the secret (e.g. 1hr, 1d, 1w, 1m, 1y)
  -h, --help                 help for create
  -p, --password string      Additionally encrypt the secret with a password

Global Flags:
      --hide-banner   Do not show the SecretBin banner message. This is useful for scripting.
```