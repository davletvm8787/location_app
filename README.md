# Location-App: A Go Adventure with Maps and Routes! 🚀

Welcome to **Location-App** — a REST API for managing locations, crafted in pure Go! Here we add points, draw routes, and calculate distances "as the crow flies" (Haversine, say hi!). This project is my ticket to the Golang dev world, and I’ve poured my heart, code, and a bit of Gopher magic into it.

<p align="center">
  <img src="https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExcTg0ZWkzNnVmMGZybWR4cjg1eDRrOGswaTFwMTdyc2Jod2g4bndvNCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/en4M5qpoxaOyUcDmYU/giphy.gif" alt="Gopher Dance" width="300"/>
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
  <img src="https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExY3VrcjNkdTJkMDVrczVydjV4c3R5aGZhOXdmdjdqcTZhN3Y5ODN5OCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/26tPskse8zR2Wkwu4/giphy.gif" alt="Go Tools" width="250"/>
</p>

---

## Architecture (or "How Does It Stand?") 🏗️

Layers, like a tasty cake:
- **Handlers**: Hey, HTTP! Parsing requests and sending responses.
- **Services**: Business logic, including Haversine magic for distances.
**Repository**: Chats with the database like an old buddy.
- **Models**: Just structs. Pure as a Gopher’s tear.

  
```
Location-app/
├── cmd
│ └── app_configs.dart
│
├── internal
│ ├── handlers
│ ├── models
│ ├── repository
│ ├── services
│ └── middleware
├── Dockerfile
│ └── docker-compose.yml #Orchestration
....
```

---

## Endpoints (or "What Can I Do?") 🌐

| Method | Endpoint         | Description                      | Example Request                    |
|--------|------------------|----------------------------------|------------------------------------|
| POST   | `/locations`     | Add a location                  | `{"name": "Home", "latitude": 55.75, "longitude": 37.61, "color": "#FF5733"}` |
| GET    | `/locations`     | List all locations              | -                                  |
| GET    | `/locations/:id` | Get one location’s details      | `/locations/1`                    |
| PUT    | `/locations/:id` | Edit a location                 | `{"name": "New Home", "color": "#00FF00"}` |
| POST   | `/route`         | Find the nearest location       | `{"latitude": 55.75, "longitude": 37.61}` |

---

## How to Run It? (or "Go Run, Gopher, Go!") 🏃‍♂️

1. **Clone the repo**:
   ```bash
   git clone https://github.com/yourusername/location-app.git
   cd location-app
   docker-compose up --build
Done! API is live at http://localhost:3000.
<p align="center"> <img src="https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExMDQyZmQzMzMwYTU4MWM5Y2Q3N2U2MjVjMzM5NTVjYjI3ZjMwYzMwOCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/xT9IgzoKnwFNmISR8I/giphy.gif" alt="Running Code" width="300"/> </p>
Bonus Points (or "What I’m Proud Of?") 🎉
Clean Code: As Rob Pike said, "Clarity is better than cleverness".
Haversine: Distances without Google Maps — just math and Go!
Rate Limiting: 100 requests per minute, chill out, bro.
Dockerized: Everything in containers, like a pro DevOps.
Want tests, CI/CD, or deployment? That’s the next chapter!

Gopher Jokes 😂
Why do Go devs love simplicity? Because if err != nil is half the program!
What’s a Gopher’s favorite dance? The Concurrency Shuffle!
What did the Gopher say about my code? "Nice goroutines, bro!"
<p align="center"> <img src="https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExMjY0YzU5YzZmNGVjYzZmY2Y5ZjYyY2Q4ZmY2YzkyOWYxZmQyYzQyNCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/3o6Zt6KHxJTbXCnSso/giphy.gif" alt="Gopher Laugh" width="250"/> </p>
