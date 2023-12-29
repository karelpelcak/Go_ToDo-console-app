# ToDo App in Go

## !!Console cleaning work only on windows!!

## Overview

This is a simple command-line ToDo application written in Go. It allows users to manage their tasks with basic functionalities like viewing all events, adding new events, and exiting the application.

## Features

1. **View All Events**
   - Option 1: Displays a list of all events currently in the ToDo list.

2. **Add New Event**
   - Option 2: Allows users to add a new event to the ToDo list.

3. **Exit**
   - Option 3: Exits the ToDo app.

## Getting Started

To run the ToDo app, ensure you have Go installed on your machine. Follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/karelpelcak/Go_ToDo-console-app.git
2. Navigate to the project directory:

   ```bash
   cd todo-app-go

3. Run the application:

   ```bash
   go run main.go


## Usage

### View All Events

- Select option 1.
- If there are no events, it will inform you that you have 0 events. Otherwise, it will display the list of all events.

### Add New Event

- Select option 2.
- Enter the event details as prompted.
- The application will confirm the successful addition of the event.

### Exit

- Select option 3 to exit the ToDo app.

### Example

```bash
Welcome to ToDo app
1 - all events
2 - add event
3 - exit
your select: 2
Enter event: Complete README file
Event added successfully.

Welcome to ToDo app
1 - all events
2 - add event
3 - exit
your select: 1
All events:
Complete README file

Welcome to ToDo app
1 - all events
2 - add event
3 - exit
your select: 3
Exiting the ToDo app. Goodbye!


