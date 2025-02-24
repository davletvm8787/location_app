# Location-App: A Go Adventure with Maps and Routes! 🚀

Welcome to **Location-App** — a REST API for managing locations, crafted in pure Go! Here we add points, draw routes, and calculate distances "as the crow flies" (Haversine, say hi!). This project is my ticket to the Golang dev world, and I’ve poured my heart, code, and a bit of Gopher magic into it.

<p align="center">
  <img src="https://media.giphy.com/media/8UGoRT66PnY6c/giphy.gif" alt="Gopher Dance" width="300"/>
</p>

---

## What’s This Thing? 🗺️

This app is part of the **Golang Developer Assessment**. It can:
- Add locations with coordinates, names, and marker colors.
- List all saved locations.
- Show details of a single location.
- Edit location data.
- Find the nearest point and "route" to it (well, at least calculate the distance!).

All wrapped in a layered architecture, Dockerized, and spiced with some Go humor.

---

## Tech Stack (or "Who’s in the Crew?") 💻

- **Go**: Because "less is exponentially more" (thanks, Rob Pike!).
- **Fiber**: Fast as a Gopher on caffeine.
- **MySQL**: Stores our points while we sip coffee.
- **GORM**: ORM, so we don’t write SQL by hand (it’s not 2010 anymore).
- **Validator**: Ensures you’re not sending me Moon coordinates.
- **Docker**: Containers are life.
- **Rate Limiting**: Keeps the server safe from your enthusiasm.

<p align="center">
  <img src="https://media.giphy.com/media/l0IylSBOt4oEmXIs8/giphy.gif" alt="Go Tools" width="250"/>
</p>

---

## Architecture (or "How Does It Stand?") 🏗️

Layers, like a tasty cake:
- **Handlers**: Hey, HTTP! Parsing requests and sending responses.
- **Services**: Business logic, including Haversine magic for distances.
- **Repository**: Chats with the database like an old buddy.
- **Models**: Just structs. Pure as a Gopher’s tear.
