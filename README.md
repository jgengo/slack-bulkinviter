# Slack Bulk Inviter

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Overview

**Slack Bulk Inviter** is a Golang-based tool designed to bulk invite all users in a Slack workspace to a specified channel. This utility is particularly useful for organizations or teams that need to regularly onboard members to new channels, ensuring that everyone is included without manual effort.

### Key Features

- **Automatic Channel Detection:** Differentiates between public, private, and converted channels to handle invitations appropriately.
- **Simple and Efficient:** Built with simplicity in mind, allowing quick execution with minimal setup.
- **Inspired by:** [slack-bulkinviter](https://github.com/robby-dermody/slack-bulkinviter) with enhancements to handle various channel types.

## Prerequisites

Before using this tool, ensure you have the following:

- **Golang installed**: [Install Golang](https://golang.org/doc/install)
- **Slack Token**: You must have a Slack token with the appropriate permissions to invite users to channels.

## Getting Started

#### 1. Clone the Repository

```sh
git clone https://github.com/jgengo/slack-bulkinviter.git
cd slack-bulkinviter
```

#### 2. Set Up Environment Variables

Export your Slack token as an environment variable:

```sh
export SLACK_TOKEN=<your-slack-token>
```

#### 3. Build the Application

Compile the Go application:

```sh
go build -o inviter main.go
```

#### 4. Run the Inviter

Execute the inviter with the channel ID you wish to invite users to:

```sh
./inviter -c <channel_id>
```

#### 5. Sit Back and Relax

The tool will handle the rest, inviting all workspace users to the specified channel.

<br>

## Usage example

To invite users to a channel with the ID C1234567890, you would run:

```sh
./inviter -c C1234567890
```

<br>

## Contributing
We welcome contributions to improve this tool! Feel free to submit pull requests or open issues for any bugs or feature requests.

To contribute:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a clear description of your changes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements

- **Inspiration:** This project was heavily inspired by [Robby Dermody's slack-bulkinviter](https://github.com/robby-d/slack-bulkinviter).
- **Special Thanks:** [Hive Helsinki](https://www.hive.fi/) for providing the use case that led to the creation of this tool.

