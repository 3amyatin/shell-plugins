# Cloudflare flarectl Plugin

Authenticate the Cloudflare CLI (`flarectl`) with 1Password.

## Requirements

- 1Password CLI installed and configured
- flarectl installed (`brew install cloudflare/cloudflare/cf-cli` or from source)
- Cloudflare account with Global API Key

## Credential Setup

1. In 1Password, create or locate your Cloudflare credential with:
   - Email field (your Cloudflare account email)
   - Field labeled "Global API Key" or similar with your API key

2. Initialize the plugin:
   ```bash
   op plugin init flarectl
   ```

3. Source your plugins file:
   ```bash
   source ~/.config/op/plugins.sh
   ```

## Usage

Once configured, use flarectl commands normally:

```bash
flarectl user info
flarectl zone list
flarectl dns list -zone example.com
```

Credentials are automatically injected via `CF_API_EMAIL` and `CF_API_KEY` environment variables.

## Resources

- [flarectl Documentation](https://github.com/cloudflare/cloudflare-go/tree/master/cmd/flarectl)
- [Cloudflare API Keys](https://developers.cloudflare.com/fundamentals/api/get-started/keys/)
