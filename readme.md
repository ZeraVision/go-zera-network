# go-zera-network
At set of Zera Network packages in Go using GRPC to communicate.

## Setup
> This is a private package which requires github authentication to fetch

How to install latest version on go project
```bash
go get -u github.com/ZeraVision/go-zera-network
```

### Troubleshooting
In the event `go get` returns a 404 error:
1. Ensure that your machine is using the correct Personal Access Token or is authenticated with Github correctly.
2. Temporarily adjust the environment variable to include the repository URL by setting `GOPRIVATE` to `github.com/ZeraVision/go-zera-network`.
   #### Powershell
   ```powershell
   $env:GOPRIVATE="github.com/*"
   ```
   #### ZSH Bash
   ```bash
   export GOPRIVATE=github.com/*
   ```

# Development

## Setup (optional: only needed to generate proto files)
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Generating proto/gRPC files
> Run from repository root.
```bash
protoc --go_out=. --go-grpc_out=. --proto_path=./grpc/proto  ./grpc/proto/*.proto
```

## Pushing New Version to Github

1. **Determine the Tag Name:**
   First, follow the [Version Naming Standards]({#naming-standards}) to decide on the appropriate tag name for the new version.

2. **Tag the Associated Commit:**
   To push a new version, you need to tag the   associated commit with the chosen tag name.   You can do this using the following command:
    ```bash
    git tag {tag name} 
    ```
3. **Push the Tag to GitHub:**
    Finally, push the newly created tag to GitHub:
    ```bash
    git push --tags 
    ```

### Version Naming Standards {#naming-standards}
- Example: `"v1.12.2"`
- Expanded Format: `"v{Major Version}.{Minor Version}.{Patch Version}"`

