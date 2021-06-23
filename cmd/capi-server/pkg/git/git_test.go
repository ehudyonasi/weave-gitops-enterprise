// +build integration

package git_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/fluxcd/go-git-providers/gitprovider"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"

	"github.com/weaveworks/wks/cmd/capi-server/pkg/git"
)

const (
	TestRepositoryNamePrefix = "capi-server-test-repo"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestCreatePullRequestInGitHubOrganization(t *testing.T) {
	// Create a client
	ctx := context.Background()
	client := github.NewClient(
		oauth2.NewClient(ctx,
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
			),
		),
	)
	// Create a repository using a name that doesn't exist already
	repoName := fmt.Sprintf("%s-%03d", TestRepositoryNamePrefix, rand.Intn(1000))
	repos, _, err := client.Repositories.ListByOrg(ctx, os.Getenv("GITHUB_ORG"), nil)
	assert.NoError(t, err)
	for findGitHubRepo(repos, repoName) != nil {
		repoName = fmt.Sprintf("%s-%03d", TestRepositoryNamePrefix, rand.Intn(1000))
	}
	repo, _, err := client.Repositories.Create(ctx, os.Getenv("GITHUB_ORG"), &github.Repository{
		Name:     github.String(repoName),
		Private:  github.Bool(true),
		AutoInit: github.Bool(true),
	})
	require.NoError(t, err)
	defer func() {
		_, err = client.Repositories.Delete(ctx, os.Getenv("GITHUB_ORG"), repo.GetName())
		require.NoError(t, err)
	}()

	s := git.NewGitProviderService()
	path := "management/cluster-01.yaml"
	content := "---\n"
	res, err := s.WriteFilesToBranchAndCreatePullRequest(ctx, git.WriteFilesToBranchAndCreatePullRequestRequest{
		GitProvider: git.GitProvider{
			Token:    os.Getenv("GITHUB_TOKEN"),
			Type:     "github",
			Hostname: "github.com",
		},
		RepositoryURL: repo.GetCloneURL(),
		HeadBranch:    "feature-01",
		BaseBranch:    repo.GetDefaultBranch(),
		Title:         "New cluster",
		Description:   "Creates a cluster through a CAPI template",
		CommitMessage: "Add cluster manifest",
		Files: []gitprovider.CommitFile{
			gitprovider.CommitFile{
				Path:    &path,
				Content: &content,
			},
		},
	})
	require.NoError(t, err)

	pr, _, err := client.PullRequests.Get(ctx, os.Getenv("GITHUB_ORG"), repo.GetName(), 1) // #PR is 1 because it is a new repo
	require.NoError(t, err)
	assert.Equal(t, pr.GetHTMLURL(), res.WebURL)
	assert.Equal(t, pr.GetTitle(), "New cluster")
	assert.Equal(t, pr.GetBody(), "Creates a cluster through a CAPI template")
	assert.Equal(t, pr.GetChangedFiles(), 1)
}

func TestCreatePullRequestInGitHubUser(t *testing.T) {
	// Create a client
	ctx := context.Background()
	client := github.NewClient(
		oauth2.NewClient(ctx,
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
			),
		),
	)
	// Create a repository using a name that doesn't exist already
	repoName := fmt.Sprintf("%s-%03d", TestRepositoryNamePrefix, rand.Intn(1000))
	repos, _, err := client.Repositories.List(ctx, os.Getenv("GITHUB_USER"), nil)
	assert.NoError(t, err)
	for findGitHubRepo(repos, repoName) != nil {
		repoName = fmt.Sprintf("%s-%03d", TestRepositoryNamePrefix, rand.Intn(1000))
	}
	repo, _, err := client.Repositories.Create(ctx, "", &github.Repository{
		Name:     github.String(repoName),
		Private:  github.Bool(true),
		AutoInit: github.Bool(true),
	})
	require.NoError(t, err)
	defer func() {
		_, err = client.Repositories.Delete(ctx, os.Getenv("GITHUB_USER"), repo.GetName())
		require.NoError(t, err)
	}()

	s := git.NewGitProviderService()
	path := "management/cluster-01.yaml"
	content := "---\n"
	res, err := s.WriteFilesToBranchAndCreatePullRequest(ctx, git.WriteFilesToBranchAndCreatePullRequestRequest{
		GitProvider: git.GitProvider{
			Token:    os.Getenv("GITHUB_TOKEN"),
			Type:     "github",
			Hostname: "github.com",
		},
		RepositoryURL: repo.GetCloneURL(),
		HeadBranch:    "feature-01",
		BaseBranch:    repo.GetDefaultBranch(),
		Title:         "New cluster",
		Description:   "Creates a cluster through a CAPI template",
		CommitMessage: "Add cluster manifest",
		Files: []gitprovider.CommitFile{
			gitprovider.CommitFile{
				Path:    &path,
				Content: &content,
			},
		},
	})
	assert.NoError(t, err)

	pr, _, err := client.PullRequests.Get(ctx, os.Getenv("GITHUB_USER"), repo.GetName(), 1) // #PR is 1 because it is a new repo
	require.NoError(t, err)
	assert.Equal(t, pr.GetHTMLURL(), res.WebURL)
	assert.Equal(t, pr.GetTitle(), "New cluster")
	assert.Equal(t, pr.GetBody(), "Creates a cluster through a CAPI template")
	assert.Equal(t, pr.GetChangedFiles(), 1)
}

func TestCreatePullRequestInGitLab(t *testing.T) {
	// Create a client
	client, err := gitlab.NewClient(os.Getenv("GITLAB_TOKEN"))
	require.NoError(t, err)
	// Create a repository using a name that doesn't exist already
	repoName := fmt.Sprintf("%s-%03d", TestRepositoryNamePrefix, rand.Intn(1000))
	repos, _, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		Owned: gitlab.Bool(true),
	})
	assert.NoError(t, err)
	for findGitLabRepo(repos, repoName) != nil {
		repoName = fmt.Sprintf("%s-%03d", TestRepositoryNamePrefix, rand.Intn(1000))
	}
	repo, _, err := client.Projects.CreateProject(&gitlab.CreateProjectOptions{
		Name:                 gitlab.String(repoName),
		MergeRequestsEnabled: gitlab.Bool(true),
		Visibility:           gitlab.Visibility(gitlab.PrivateVisibility),
		InitializeWithReadme: gitlab.Bool(true),
	})
	require.NoError(t, err)
	defer func() {
		_, err = client.Projects.DeleteProject(repo.ID)
		require.NoError(t, err)
	}()

	s := git.NewGitProviderService()
	path := "management/cluster-01.yaml"
	content := "---\n"
	res, err := s.WriteFilesToBranchAndCreatePullRequest(context.Background(), git.WriteFilesToBranchAndCreatePullRequestRequest{
		GitProvider: git.GitProvider{
			Token:    os.Getenv("GITLAB_TOKEN"),
			Type:     "gitlab",
			Hostname: "gitlab.com",
		},
		RepositoryURL: repo.HTTPURLToRepo,
		HeadBranch:    "feature-01",
		BaseBranch:    repo.DefaultBranch,
		Title:         "New cluster",
		Description:   "Creates a cluster through a CAPI template",
		CommitMessage: "Add cluster manifest",
		Files: []gitprovider.CommitFile{
			gitprovider.CommitFile{
				Path:    &path,
				Content: &content,
			},
		},
	})
	assert.NoError(t, err)

	pr, _, err := client.MergeRequests.GetMergeRequest(repo.ID, 1, nil) // #PR is 1 because it is a new repo
	require.NoError(t, err)
	assert.Equal(t, pr.WebURL, res.WebURL)
	assert.Equal(t, pr.Title, "New cluster")
	assert.Equal(t, pr.Description, "Creates a cluster through a CAPI template")
}

func findGitHubRepo(repos []*github.Repository, name string) *github.Repository {
	if name == "" {
		return nil
	}
	for _, repo := range repos {
		if repo.GetName() == name {
			return repo
		}
	}
	return nil
}

func findGitLabRepo(repos []*gitlab.Project, name string) *gitlab.Project {
	if name == "" {
		return nil
	}
	for _, repo := range repos {
		if repo.Name == name {
			return repo
		}
	}
	return nil
}