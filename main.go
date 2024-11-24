package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
}

func main() {
	// Ensure a GitHub username is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activities <username>")
		os.Exit(1)
	}
	username := os.Args[1]

	// Fetch GitHub activity for the given username
	events, err := fetchGithubActivity(username)
	if err != nil {
		fmt.Printf("Error fetching activity: %s\n", err)
		os.Exit(1)
	}

	// Handle case where no activity is found
	if len(events) == 0 {
		fmt.Printf("No activity found for user %s\n", username)
		return
	}

	// Display the user's recent activity
	fmt.Printf("Recent activity for user %s:\n", username)
	for _, event := range events {
		displayEvent(event)
	}
}

func fetchGithubActivity(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch activity: HTTP status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var events []GitHubEvent
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return events, nil
}

func displayEvent(event GitHubEvent) {
	timestamp, err := time.Parse(time.RFC3339, event.CreatedAt)
	if err != nil {
		fmt.Println("- Invalid timestamp")
		return
	}
	formattedTime := timestamp.Format("Jan 2, 2006 at 15:04")

	switch event.Type {
	case "PushEvent":
		fmt.Printf("- Pushed to %s on %s\n", event.Repo.Name, formattedTime)
	case "PullRequestEvent":
		fmt.Printf("- Created a pull request in %s on %s\n", event.Repo.Name, formattedTime)
	case "IssuesEvent":
		fmt.Printf("- Opened an issue in %s on %s\n", event.Repo.Name, formattedTime)
	case "WatchEvent":
		fmt.Printf("- Starred %s on %s\n", event.Repo.Name, formattedTime)
	default:
		fmt.Printf("- %s on %s\n", event.Type, formattedTime)
	}
}
