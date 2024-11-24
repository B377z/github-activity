# GitHub Activity CLI

A simple command-line tool that fetches and displays the recent public activity of a GitHub user.

## Features

- Fetches recent public events for any GitHub user.
- Displays event types such as:
  - Pushes to repositories
  - Pull requests
  - Issues opened
  - Stars added
  - Forks created
- Outputs events in a human-readable format, including the repository name and timestamp.

## Installation

1. Clone the repository:

   git clone https://github.com/<your-username>/<repository-name>.git

2. Navigate to the project directory:

   cd <repository-name>

3. Build the CLI:

   go build -o github-activities


## Usage

Run the CLI with the GitHub username as an argument:

./github-activities <username>

**Output:**

Recent activity for user octocat:
- Pushed to octocat/Hello-World on Nov 23, 2024 at 14:30
- Opened an issue in octocat/Hello-World on Nov 22, 2024 at 10:15
- Starred octocat/Spoon-Knife on Nov 21, 2024 at 16:45


## Error Handling

- If the username is invalid or does not exist:

  Error fetching activity: HTTP status code 404

- If the user has no recent activity:

  No activity found for user <username>


## Dependencies

This project uses Go's standard library:
- `net/http` for HTTP requests.
- `encoding/json` for parsing JSON responses.
- `io` for reading HTTP response bodies.

## Contribution

Contributions are welcome! Feel free to submit a pull request or open an issue for any bugs or feature requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Project URL
https://roadmap.sh/projects/github-user-activity

