# How to contribute

We'd love to accept your patches and contributions to this project. There are
a just a few small guidelines you need to follow.

## Reporting issues

Bugs, feature requests, and development-related questions should be directed to
our [GitHub issue tracker](https://github.com/hatch-ed-com/ri-sdk-go/issues). If
reporting a bug, please try and provide as much context as possible such as
your operating system, Go version, and anything else that might be relevant to
the bug. For feature requests, please explain what you're trying to do, and
how the requested feature would help you do that.

## Submitting a patch

1. It's generally best to start by opening a new issue describing the bug or
   feature you're intending to fix. Even if you think it's relatively minor,
   it's helpful to know what people are working on. Mention in the initial issue
   that you are planning to work on that bug or feature so that it can be
   assigned to you.

2. Follow the normal process of [forking][1] the project, and set up a new branch
   to work in. It's important that each group of changes be done in separate
   branches in order to ensure that a pull request only includes the commits
   related to that bug or feature.

3. Any significant changes should almost always be accompanied by tests. The
   project already has good test coverage, so look at some of the existing tests
   if you're unsure how to go about it.

4. Run `script/fmt.sh`, `script/test.sh` to format your code and
   check that it passes all tests.

5. Do your best to have [well-formed commit messages][2] for each change. This
   provides consistency throughout the project, and ensures that commit messages
   are able to be formatted properly by various git tools.

6. Finally, push the commits to your fork and submit a [pull request][3].
   **NOTE:** Please do not use force-push on PRs in this repo, as it makes it
   more difficult for reviewers to see what has changed since the last code
   review. We always perform "squash and merge" actions on PRs in this repo, so it doesn't
   matter how many commits your PR has, as they will end up being a single commit after merging.
   This is done to make a much cleaner `git log` history and helps to find regressions in the code
   using existing tools such as `git bisect`.

[1]: https://help.github.com/articles/fork-a-repo
[2]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
[3]: https://help.github.com/articles/creating-a-pull-request

## Code Comments

Every exported method needs to have code comments that follow
[Go Doc Comments][1]. A typical method's comments will look like this:

```go
// Retrieves metadata for files within the Connect files
// module and logs. This does NOT retrieve the file contents
// only the metadata as shown in the GetConnectFilesOutput
//
//meta:operation GET /admin/connect/files/{path}
func (c *Client) GetConnectFiles(getConnectFilesInput GetConnectFilesInput) (*GetConnectFilesOutput, error) {
var output GetConnectFilesOutput

url := fmt.Sprintf("%s/admin/connect/files/%s?project=%s", c.baseEndpoint, getConnectFilesInput.Path, getConnectFilesInput.Project)
req, err := c.GenerateRequest("GET", url, nil)
...
}
```

The first paragraph is a summary of what the method does, and any special
callouts that a typical user would not be aware of.

The `//meta:operation` comment maps the method to the RapidIdentity OpenAPI
Specification endpoint.

[1]: https://go.dev/doc/comment

## Scripts

The `script` directory has shell scripts that help with common development
tasks.

**script/fmt.sh** formats all go code in the repository.

**script/test.sh** runs tests on all modules.

## Other notes on code organization

Currently, everything is defined in the main `rapididentity` package. Code is organized in files
based on the RapidIdentity API Swagger documentation

## Maintainer's Guide

(These notes are mostly only for people merging in pull requests.)

**When creating a release, don't forget to update the `Version` constant in `RapidIdentity.go`.** This is used to
send the version in the `User-Agent` header to identify clients to the RapidIdentity API.
