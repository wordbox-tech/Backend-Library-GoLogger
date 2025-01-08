# Backend-Library-GoLogger

Logging in Go using the Uber zap library with specific configuration to create Google Cloud logs

## Execution
1. For Installing, run the next command:
    - `go get -u github.com/wordbox-tech/Backend-Library-GoLogger@vX.X.X`
2. To check if all dependencies are satisfied, run the next command:
    - `go mod tidy -v`

## Usage
import the library  
    - `github.com/wordbox-tech/Backend-Library-GoLogger`

Functions:  
- `LoggerWordbox.Info("message: %s, another %s", "hi", "world")`
- `LoggerWordbox.Error("...")`
- `LoggerWordbox.Warn("...")`
- `LoggerWordbox.Debug("message: %s, another %s", "hi", "world")`
- `LoggerWordbox.Panic("message: %s, another %s", "hi", "world")`
- `LoggerWordbox.Fatal("message: %s, another %s", "hi", "world")`
