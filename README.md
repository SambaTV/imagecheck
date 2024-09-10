# Imagecheck

The `imagecheck` application checks a container image and its associated source
code and config artifacts for defects and vulnerabilities using multiple
scanners, optionally uploading scan summaries and output to an S3 bucket.

It is intended to be used in a CI/CD pipeline after images are built and before
they are pushed to a container registry to ensure they are safe for use.

It is also intended to be used by developers interactively during local
development and testing before changes are committed to the repository and
pushed to their upstream remote.

Its user documentation is available at https://sambatv.github.io/imagecheck.

## Organization

The `imagecheck` application is written in Golang, requiring Go 1.23 or later,
and is organized as follows:

* [`.github/workflows/release.yaml`](.github/workflows/release.yaml) - GitHub Actions release workflow
* [`app/`](app) - application library source code
* [`bin/`](bin) - application binary artifacts (ignored, populated by the Makefile `deps` and `build` targets)
* [`cli/`](cli) - application command line interface source code
* [`docs/`](docs) - project documentation hosted on GitHub Pages using [docsify](https://docsify.js.org/)
* [`.tool-versions`](.tool-versions) - [asdf](https://asdf-vm.com/)-managed toolchain versions
* [`Dockerfile`](Dockerfile) - application container image build file
* [`go.mod`](go.mod) - application Go module definition
* [`go.sum`](go.sum) - application Go module checksums
* [`Makefile`](Makefile) - project automation for developers
* [`metadata/`](metadata) - application metadata source code
* [`README.md`](README.md) - this document
* [`VERSION`](VERSION) - application version file

## Development

Project automation is provided by the [`Makefile`](Makefile).

Run the `make` command with no arguments to see the available targets:

```shell
make

Usage:
  make <target>

Info targets
  help             Show this help
  vars             Show environment variables used by this Makefile

Dependency targets
  deps             Install scanner dependencies

Application targets
  build            Build the application
  lint             Lint the application
  scan             Scan the application for defects and vulnerabilities
  test             Run the application tests
  clean            Clean application build artifacts

Image targets
  image-build      Build the container image
  image-scan       Scan the container image for defects and vulnerabilities
  image-push       Push the container image
```

The first thing that should be done is to install the project dependencies:

```shell
make deps
```

This will install the scanners invoked by the application into the `.bin/`
directory.

The next thing that should be done is to build the application:

```shell
make build
```

This will build the application binary and place it in the `bin/` directory.

The application can then be run locally:

```shell
bin/imagecheck --help
```

> It is recommended to add the `bin/` directory to your `PATH` for ease of use
> in your shell environment:
> 
> ```shell
> export PATH=$(pwd)/bin:$PATH
> ```
> 
> If you are using [https://direnv.net/], you should add that export statement
> to your `.envrc` file:
> 
> ```shell
> export PATH=$(pwd)/bin:$PATH
> ```
> 
> The `.envrc` file can also be used to set other environment variables used by
> `imagecheck`, as well as others such as AWS credentials for S3 bucket access in
> `--pipeline` mode.

Run tests with:

```shell
make test
```

Run linters with:

```shell
make lint
```

Build the container image with:

```shell
make image-build
```

Scan the container image with `imagecheck` itself:

```shell
make image-scan
```

### Release workflow

The release workflow is automated using [GitHub Actions](https://docs.github.com/en/actions.
as defined in the [`.github/workflows/release.yaml`](.github/workflows/release.yaml)
file when a new version tag on the `main` branch is pushed to the remote repository.

The development process to make a release is as follows:

1. A fix or feature branch off of main is created for the changes to be made.
2. The changes are made and committed to the fix or feature branch.
3. The changes are pushed to the remote repository.
4. A pull request is created from the fix or feature branch to main.
5. The pull request is reviewed and approved.
6. The pull request is merged into main.
7. The `VERSION` file is updated with the new version number.
8. A commit is made with the updated `VERSION` file.
9. The commit is tagged with the new version number.
10. The commit and tag are pushed to the remote repository.
11. The release workflow is triggered by the new tag.
12. The release workflow builds the application binaries and container image.
13. The release workflow uploads the application binaries and container image to
    its GitHub Releases and GitHub Packages.
14. The GitHub Pages documentation reflects any new changes to content in the
    `docs/` directory.

## License

This project is not rocket science or secret sauce. Have at it if you find it useful.

This project is licensed under the Apache License, Version 2.0. See the
[LICENSE](LICENSE) file for the full license text.
