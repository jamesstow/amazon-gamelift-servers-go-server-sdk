# Go server SDK for Amazon GameLift Servers

## Documentation

You can find the official Amazon GameLift Servers documentation [here](https://docs.aws.amazon.com/gamelift/).

## Minimum requirements:

- [Go 1.18 or newer](https://golang.org/dl/)
- [Make](https://www.gnu.org/software/make/) utility and [Docker](https://www.docker.com/) to run tests and linter

## Installation

```bash
go get github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5
```

## Installation (Local Beta)

1.  You can add the local module by using replace in your go.mod file:
    ```go
    // replace the local path with a relative path from your project root to where you unpacked the SDK
    replace github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5 => ../path/to/amazon-gamelift-servers-go-server-sdk
    ```
2.  Import it in your code:
    ```golang
    import (
        "github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/model"
        "github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/server"
    )
    ```
3.  Then run go mod tidy in your project root
    ```sh
    go mod tidy
    ```
    This will set up the proper local dependencies.

### Example code

```golang
package main

import (
	"github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/model"
	"github.com/jamesstow/amazon-gamelift-servers-go-server-sdk/v5/server"
	"log"
)

type gameProcess struct {
	// Port - port for incoming player connections
	Port int

	// Logs - set of files to upload when the game session ends.
	// Amazon GameLift Servers will upload everything specified here for the developers to fetch later.
	Logs server.LogParameters
}

func (g gameProcess) OnStartGameSession(model.GameSession) {
	// When a game session is created,
	// Amazon GameLift Servers sends an activation request to the game server and passes
	// along the game session object containing game properties and other settings.
	// Here is where a game server should take action based on the game session object.
	// Once the game server is ready to receive incoming player connections,
	// it should invoke server.ActivateGameSession()
	err := server.ActivateGameSession()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (g gameProcess) OnUpdateGameSession(model.UpdateGameSession) {
	// When a game session is updated (e.g. by FlexMatch backfill),
	// Amazon GameLift Servers sends a request to the game
	// server containing the updated game session object.
	// The game server can then examine the provided
	// MatchmakerData and handle new incoming players appropriately.
	// UpdateReason is the reason this update is being supplied.
}

func (g gameProcess) OnProcessTerminate() {
	// Amazon GameLift Servers will invoke this callback before shutting down an instance hosting this game server.
	// It gives this game server a chance to save its state,
	// communicate with services, etc., before being shut down.
	// In this case, we simply tell Amazon GameLift Servers we are indeed going to shutdown.
	server.ProcessEnding()
}

func (g gameProcess) OnHealthCheck() bool {
	// Amazon GameLift Servers will invoke this callback every HEALTHCHECK_INTERVAL times (60 sec by default, with jitter.)
	// Here, a game server might want to check the health of dependencies and such.
	// Simply return true if healthy, false otherwise.
	// The game server has HEALTHCHECK_TIMEOUT interval (60 sec by default) to respond with its health status.
	// Amazon GameLift Servers will default to 'false' if the game server doesn't respond in time.
	// In this case, we're always healthy!
	return true
}

func main() {
	err := server.InitSDK(server.ServerParameters{
		WebSocketURL: "wss://1234abcdef.execute-api.us-west-2.amazonaws.com/prod",
		ProcessID:    "myProcess",
		HostID:       "myHost",
		FleetID:      "myFleet",
		AuthToken:    "auth_token_example",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	// Make sure to call server.Destroy() when the application quits.
	// This frees the server SDK from memory.
	// As a best practice, call this method after ProcessEnding() and before terminating the process.
	defer server.Destroy()
	process := gameProcess{
		Port: 7777,
		Logs: server.LogParameters{
			LogPaths: []string{"/local/game/logs/myserver.log"},
		},
	}
	err = server.ProcessReady(server.ProcessParameters{
		OnStartGameSession:  process.OnStartGameSession,
		OnUpdateGameSession: process.OnUpdateGameSession,
		OnProcessTerminate:  process.OnProcessTerminate,
		OnHealthCheck:       process.OnHealthCheck,
		LogParameters:       process.Logs,
		Port:                process.Port,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Start handling player connections here.
}
```
