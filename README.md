# risk-management-octave-allegro

# HTMX + Go Fiber

A proof of concept HTMX app with a Go Fiber backend. This project is an application of risk management using a well-defined framework that is already popular in the field of cybersecurity, namely OCVATE ALLEGRO.

<img width="959" alt="image" src="https://github.com/AliNicholas/risk-management-octave-allegro/assets/109023698/f44876c2-d297-425b-97ed-fbaa60a5b351">

<img width="960" alt="image" src="https://github.com/AliNicholas/risk-management-octave-allegro/assets/109023698/9ffcd334-aac7-432e-b417-cb2b6fb55896">

<img width="960" alt="image" src="https://github.com/AliNicholas/risk-management-octave-allegro/assets/109023698/0f3f3702-8bb6-41db-bcde-1bb3ef4efd4d">

<img width="554" alt="image" src="https://github.com/AliNicholas/risk-management-octave-allegro/assets/109023698/b7c3734c-3e32-42d3-b327-b8a7399ec94c">


## Requirements

- Go `1.21+`

## Installation

Fetch the dependencies
```sh
go get
```

## Development
```

There are two options to start the web server:
- without hot reloading
- with hot reloading

### Without hot reloading

```sh
go run .
```

### With hot reloading

Go Fiber does not have a hot reloading feature. Install [`air`](https://github.com/cosmtrek/air#installation) to run
the app with hot reloading.

```sh
air
```

```
  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.49.0, built with Go go1.21.4

watching .
watching bin
!exclude node_modules
watching public
watching src
!exclude tmp
watching views
building...
running...

 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.51.0                   │
 │               http://127.0.0.1:3000               │
 │       (bound on host 0.0.0.0 and port 3000)       │
 │                                                   │
 │ Handlers ............. 7  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 76782 │
 └───────────────────────────────────────────────────┘
```

Then open the browser to http://localhost:3000. You should be able to search for a Risk Management Application, you can create your own project and identfy the assets and threats you may have.
