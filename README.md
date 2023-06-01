# Watchdog

## Introduction

Watchdog is a dynamic monitoring tool designed specifically for tracking and alerting on cryptocurrency and DeFi silos. The Watchdog system keeps a continuous watch on various crypto assets, monitoring their price movement and silo changes, and instantly notifies users about significant events through Discord alerts.

With features to monitor price movement, assess threshold values, and scan DeFi silos, Watchdog is a comprehensive tool that offers a bird's-eye view of your crypto assets' performance and health.

## Features

1. **CoinGecko Price Monitoring:** Watchdog checks the price of specified coins at regular intervals, allowing you to keep track of their market values in real time.

2. **Silo Monitoring:** Watchdog keeps an eye on specified DeFi silos, monitoring the total deposited, available to borrow, and utilization rates.

3. **Alert System:** Watchdog comes with a configurable alerting system that triggers alerts based on certain conditions. The alert system supports two types of alerts:
   - Movement Alert: Triggered when a specified percentage change occurs in the value of a crypto asset or a silo parameter.
   - Threshold Alert: Triggered when a silo parameter (like available to borrow) crosses a certain threshold.

4. **Discord Integration:** Watchdog uses Discord Webhooks to send alerts, ensuring that you get instant notifications about the events that matter to you.

5. **Flexible Configuration:** With a configuration file (config.yaml), you can easily customize Watchdog to suit your needs. You can specify the coins, silos, and alert conditions you want to monitor.

## File Structure

Watchdog's source code is organized into different packages and files, each with a specific role:

- `main.go`: The entry point to the application.
- `watchdog/alerts`: Handles the logic for checking and triggering alerts.
- `watchdog/config`: Defines the configuration structs and provides functions to retrieve configuration information.
- `watchdog/discord`: Contains the logic for sending alert messages to Discord.
- `watchdog/jobs`: Contains the jobs that run at regular intervals to check prices and monitor silos.
- `watchdog/silo`: Defines the logic for scraping token data from silos.
- `watchdog/state`: Manages the state of the system and keeps track of previous values.
- `watchdog/utils`: Contains utility functions used across the system.

## Environment Configuration

Before running the Watchdog application, you must set the `WATCHDOG_DIR` environment variable. This variable specifies the directory where Watchdog will look for the configuration file.

The configuration file should be named config.yaml and must be located in the directory specified by the `WATCHDOG_DIR` environment variable.

## Usage

To start using Watchdog, clone the repository and configure the `config.yaml` file according to your needs:

```yaml
discordWebhookUrl: YOUR_DISCORD_WEBHOOK_URL
coins:
  - name: coin_name
    displayName: COIN_DISPLAY_NAME

silos:
  - name: silo_name
    address: silo_address
    displayName: SILO_DISPLAY_NAME
    tokens:
      - name: token_name
        displayName: TOKEN_DISPLAY_NAME

alerts:
  movement:
    - name: entity_name
      change: change_percentage
  threshold:
    - name: entity_name
      amount: threshold_amount
```

Example:

```yaml
discordWebhookUrl: https://discord.com/api/webhooks/...

coins:
- name: ethereum # the CoinGecko token ID
  displayName: ETH
- name: magic
  displayName: MAGIC
- name: sui
  displayName: SUI

silos:
- name: magic
  address: 0x30c4aa967f68705ab5677ebe17b3affd0c59e71c
  displayName: Magic
  tokens:
  - name: magic
    displayName: MAGIC # this name must match the silo's name displayed in their app for scraping purposes
  - name: eth
    displayName: Ether
  - name: USDC
    displayName: USD Coin (Arb1)

alerts:
  movement:
  - name: coins.ethereum
    change: 1%
  - name: coins.magic
    change: 5%
  - name: coins.sui
    change: 0.01
  - name: silos.magic.magic.totalDeposited
    change: 10K
  - name: silos.magic.magic.utilization
    change: 5%

  threshold:
  - name: silos.magic.magic.availableToBorrow
    amount: 200K
```

Once the configuration is set, run the application using the Go command:

```bash
go run main.go
```

Watchdog will start monitoring the specified coins and silos, and it will send alerts to the specified Discord Webhook URL based on the alert conditions defined in the config file.

Enjoy the peace of mind that comes from knowing that Watchdog is keeping a vigilant eye on your crypto assets!

Please note that this is a basic overview of Watchdog. For more detailed information about configuration and usage, refer to the code.

