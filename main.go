package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


type WebhookPayload struct {
	Event string `json:"event"`
	Build struct {
		ID           string      `json:"id"`
		GraphqlID    string      `json:"graphql_id"`
		URL          string      `json:"url"`
		WebURL       string      `json:"web_url"`
		Number       int         `json:"number"`
		State        string      `json:"state"`
		Blocked      bool        `json:"blocked"`
		BlockedState string      `json:"blocked_state"`
		Message      string      `json:"message"`
		Commit       string      `json:"commit"`
		Branch       string      `json:"branch"`
		Tag          interface{} `json:"tag"`
		Source       string      `json:"source"`
		Author       struct {
			Username string `json:"username"`
			Name     string `json:"name"`
			Email    string `json:"email"`
		} `json:"author"`
		Creator     interface{} `json:"creator"`
		CreatedAt   time.Time   `json:"created_at"`
		ScheduledAt time.Time   `json:"scheduled_at"`
		StartedAt   time.Time   `json:"started_at"`
		FinishedAt  time.Time   `json:"finished_at"`
		MetaData    struct {
			BuildkiteGitCommit string `json:"buildkite:git:commit"`
		} `json:"meta_data"`
		PullRequest interface{} `json:"pull_request"`
		RebuiltFrom interface{} `json:"rebuilt_from"`
	} `json:"build"`
	Pipeline struct {
		ID                              string      `json:"id"`
		GraphqlID                       string      `json:"graphql_id"`
		URL                             string      `json:"url"`
		WebURL                          string      `json:"web_url"`
		Name                            string      `json:"name"`
		Description                     string      `json:"description"`
		Slug                            string      `json:"slug"`
		Repository                      string      `json:"repository"`
		ClusterID                       interface{} `json:"cluster_id"`
		BranchConfiguration             interface{} `json:"branch_configuration"`
		DefaultBranch                   string      `json:"default_branch"`
		SkipQueuedBranchBuilds          bool        `json:"skip_queued_branch_builds"`
		SkipQueuedBranchBuildsFilter    interface{} `json:"skip_queued_branch_builds_filter"`
		CancelRunningBranchBuilds       bool        `json:"cancel_running_branch_builds"`
		CancelRunningBranchBuildsFilter interface{} `json:"cancel_running_branch_builds_filter"`
		AllowRebuilds                   bool        `json:"allow_rebuilds"`
		Provider                        struct {
			ID       string `json:"id"`
			Settings struct {
				TriggerMode                             string `json:"trigger_mode"`
				BuildPullRequests                       bool   `json:"build_pull_requests"`
				PullRequestBranchFilterEnabled          bool   `json:"pull_request_branch_filter_enabled"`
				SkipBuildsForExistingCommits            bool   `json:"skip_builds_for_existing_commits"`
				SkipPullRequestBuildsForExistingCommits bool   `json:"skip_pull_request_builds_for_existing_commits"`
				BuildPullRequestReadyForReview          bool   `json:"build_pull_request_ready_for_review"`
				BuildPullRequestLabelsChanged           bool   `json:"build_pull_request_labels_changed"`
				BuildPullRequestForks                   bool   `json:"build_pull_request_forks"`
				PrefixPullRequestForkBranchNames        bool   `json:"prefix_pull_request_fork_branch_names"`
				BuildBranches                           bool   `json:"build_branches"`
				BuildTags                               bool   `json:"build_tags"`
				CancelDeletedBranchBuilds               bool   `json:"cancel_deleted_branch_builds"`
				PublishCommitStatus                     bool   `json:"publish_commit_status"`
				PublishCommitStatusPerStep              bool   `json:"publish_commit_status_per_step"`
				SeparatePullRequestStatuses             bool   `json:"separate_pull_request_statuses"`
				PublishBlockedAsPending                 bool   `json:"publish_blocked_as_pending"`
				UseStepKeyAsCommitStatus                bool   `json:"use_step_key_as_commit_status"`
				FilterEnabled                           bool   `json:"filter_enabled"`
				Repository                              string `json:"repository"`
				PullRequestBranchFilterConfiguration    string `json:"pull_request_branch_filter_configuration"`
				FilterCondition                         string `json:"filter_condition"`
			} `json:"settings"`
			WebhookURL string `json:"webhook_url"`
		} `json:"provider"`
		BuildsURL string `json:"builds_url"`
		BadgeURL  string `json:"badge_url"`
		CreatedBy struct {
			ID        string    `json:"id"`
			GraphqlID string    `json:"graphql_id"`
			Name      string    `json:"name"`
			Email     string    `json:"email"`
			AvatarURL string    `json:"avatar_url"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"created_by"`
		CreatedAt            time.Time   `json:"created_at"`
		ArchivedAt           interface{} `json:"archived_at"`
		Env                  interface{} `json:"env"`
		ScheduledBuildsCount int         `json:"scheduled_builds_count"`
		RunningBuildsCount   int         `json:"running_builds_count"`
		ScheduledJobsCount   int         `json:"scheduled_jobs_count"`
		RunningJobsCount     int         `json:"running_jobs_count"`
		WaitingJobsCount     int         `json:"waiting_jobs_count"`
		Visibility           string      `json:"visibility"`
		Tags                 []string    `json:"tags"`
		Configuration        string      `json:"configuration"`
		Steps                []struct {
			Type                string      `json:"type"`
			Name                string      `json:"name"`
			Command             string      `json:"command"`
			ArtifactPaths       interface{} `json:"artifact_paths"`
			BranchConfiguration interface{} `json:"branch_configuration"`
			Env                 struct {
			} `json:"env"`
			TimeoutInMinutes interface{}   `json:"timeout_in_minutes"`
			AgentQueryRules  []interface{} `json:"agent_query_rules"`
			Concurrency      interface{}   `json:"concurrency"`
			Parallelism      interface{}   `json:"parallelism"`
		} `json:"steps"`
	} `json:"pipeline"`
	Sender interface{} `json:"sender"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Must set $PORT")
	}
	r := gin.Default()
	r.POST("/build-status", func(c *gin.Context) {
		b := new(bytes.Buffer)
		b.ReadFrom(c.Request.Body)

		x := b.Bytes()
		req := x
		var buildkite WebhookPayload
		if err := json.Unmarshal(req, &buildkite); err != nil {
			fmt.Println("Could not unmarshal JSON")
		}
		if buildkite.Build.Blocked {
			log.Printf("!! Build #%v is currently blocked! Unblock it here: %v", buildkite.Build.Number, buildkite.Build.WebURL)
		}
		log.Printf("Received event %v for pipeline %v and build %v with state %v", buildkite.Event, buildkite.Pipeline.Name, buildkite.Build.Number, buildkite.Build.State)
	})
	r.POST("/timeout", func(c *gin.Context) {
		fmt.Printf("Waiting for 80 seconds")
		time.Sleep(80 * time.Second)
	})

	r.Run(":" + port)
}
