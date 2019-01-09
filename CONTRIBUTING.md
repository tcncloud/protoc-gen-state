# Contributing Guidelines
Thank you for considering contributing!

## Pull Request Checklist
Before sending your pull requests, make sure you followed this list.

- Check if my changes are consistent with the [guidelines](https://github.com/tcncloud/protoc-gen-state/blob/master/CONTRIBUTING.md#general-guidelines-and-philosophy-for-contribution).
- Run unit tests.
- Maintainers will look to review the PR as soon as possible. If there is no traction for some time, you're welcome to bump the thread.
- All PRs require at least one reviewer.


#### General guidelines and philosophy for contribution

* If you provide a new output include an end to end example.
* Include unit tests when you contribute new features, as they help to
a) prove that your code works correctly, and b) guard against future breaking
changes to lower the maintenance cost.
* Bug fixes also generally require unit tests, because the presence of bugs
usually indicates insufficient test coverage.
* Keep API compatibility in mind when you change code,
  * When you contribute a new feature to `protoc-gen-state`, the maintenance burden is (by
      default) transferred to the `protoc-gen-state` team. This means that benefit of the
  contribution must be compared against the cost of maintaining the feature.

## Releasing
  Your changes will be released with the next version release.

