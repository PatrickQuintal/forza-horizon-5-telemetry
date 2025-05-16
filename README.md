# forza-horizon-5-telemetry

## overview

GO client that listens to Forza Horizon 5's UDP stream and prints some pretty shit to your console.

_Warning, contains AI Shlop randomly. Grug brain lazy._

## How to run

It listens on 127.0.0.1 & port 9999

`go run .\client\`

Or if you dont have forza installed

`go run .\client\ -debug -debugfile "./debugpacketstream"`

will run using some sample data instead of a UDP stream.

Its 600 packets that loop.
