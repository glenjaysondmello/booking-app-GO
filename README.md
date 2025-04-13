# 🎟️ Go Conference Booking App

A simple command-line application in Go for booking tickets to the "Go Conference". The application demonstrates Go fundamentals such as functions, packages, structs, slices, goroutines, WaitGroups, and input validation.

---

## 📌 Features

- Book tickets for a conference using CLI
- User input validation (name, email, ticket count)
- Concurrency using goroutines and WaitGroup to simulate email ticket sending
- Real-time ticket tracking
- Organized modular structure with helper functions

---

## 🛠️ Technologies Used

- **Go (Golang)** `v1.23.5`
- Standard Library (`fmt`, `time`, `sync`, `strings`)

---

## 📁 Project Structure

```
booking-app/
│
├── main.go           # Main application logic
├── helper/
│   └── helper.go     # Input validation logic
├── go.mod            # Module definition
```

---

## 🚀 Getting Started

### Prerequisites

- [Go installed](https://golang.org/dl/) (version >= 1.23.5)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/booking-app.git
   cd booking-app
   ```

2. Initialize the Go module:

   ```bash
   go mod tidy
   ```

---

## ⚙️ Usage

Run the application:

```bash
go run main.go
```

Follow the on-screen instructions to:

- Enter your name, email, and ticket count
- Receive a (simulated) email confirmation via goroutine
- View booking summary

---

## ✅ Sample Output

```
Welcome to Go Conference booking application
We have total of 50 tickets and are 50 still available
Get your tickets here to attend

Enter the firstname:
Glen
Enter the lastname:
Dmello
Enter the email:
glen@example.com
Enter the number of tickets:
2

Thank you Glen Dmello for booking 2 tickets. You will receive a confirmation email at glen@example.com
48 tickets are remaining
These are all our bookings: [Glen]
...
```

---

## 🔐 Input Validations

- **Name** must be longer than 2 characters
- **Email** must contain an `@` symbol
- **Tickets** must be a positive number and less than or equal to the remaining tickets

---

## 📬 Simulated Email Feature

The application simulates sending tickets via email by printing confirmation after a 10-second delay using a goroutine and `time.Sleep`.

---

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

