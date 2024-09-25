# An eCommerce MVP

An eCommerce project working towards preventing spam reviewers and vendors.

## Getting Started
### Prerequisites
- [Golang](https://go.dev/doc/install)
- Node.js on [Linux & Mac](https://github.com/nvm-sh/nvm) | [Windows](https://nodejs.org/en)
- [Docker Compose](https://docs.docker.com/compose/)

### Configuration
In the `.env` file, you can find the global environmental variables I use.
- To configure your email server add your email address to `EMAIL_USER`, email password to `EMAIL_PASS`, and your chosen email provider to `EMAIL_HOST`. Add whatever subdomain you want to use for your email auth service like 'email+auth@domain.com' to `EMAIL_AUTH`.
- Ports are dynamically set in `compose.yaml` using `.env`, so you can change them directly in `.env`.
## Usage
Checkout the `Makefile` for my shortcuts if you'd like.
1. Build the app image and pull the images for redis, mongo, and mongo-express:
```bash
docker-compose build
```
2. Create the containers using the images and the environment variables preset in `.env`
```bash
docker-compose up
```
  - To run in detach mode you can add `--detach` or `-d` for short:
```bash
docker-compose up -d
```
## Development
Docs are coming soon.

### Setup
Install [Templ](https://templ.guide/quick-start/installation) by running:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```
Verify installation was successful by running:
```bash
templ version
```
Install TailwindCSS and other npm dependencies you can run:
```bash
npm install
```
To install TailwindCSS for your own project [:link:](https://tailwindcss.com/docs/installation).
## Usage

To generate or regenerate templ files:

```bash
templ generate
```

For other commands checkout my Makefile.

#### Troubleshooting
- If you run into the issue `command not found` when running `templ` you need to find the path of the go package. If you use Linux, try running:
  ```bash
  which templ
  ```
  If that doesn't work look for it in your home directory. Mine was at:
  ```bash
  /home/<username>/go/bin
  ```
  Once you find the dir, you can add it to your PATH by running:
  ```bash
  export PATH=$PATH:/path/to/your/go/bin
  ```
  The `go/bin` directory is where Go installs command-line tools and executables when you use the go install command. By adding this directory to your PATH, you can run these tools directly from the command line without specifying their full path.

  Verify installation was successful by running:
  ```bash
  templ version
  ```

### Running the tests
Currently there are only unit tests:
```bash
go test ./...
```
Integration tests and CI/CD are in the works.

### Logging
Currently logging uses structured json logs with a request id. My next steps are to store the logs in a directory and export it to a log server for analysis.