# Ensure your Cloudflare domain always points to the machine running this service

## Overcomes problem for a host machine with a dynamic IP address

### How it works

1. The service Pings an IP service regularly to find your public IP address.
2. It Checks if it is different to your Cloudflare A record(s) for the target domain.
3. It will update your records if the public IP has changed
