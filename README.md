# Project Documentation

## Overview

This project is a web application built using the Go programming language. It utilizes various technologies such as HTMX for dynamic HTML updates, Tailwind CSS for styling, and SQLite as the database.

## Folder Structure

The project follows a standard Go project structure:

- `cmd`: This folder contains the main application entry point, gonote.go.
- `internals`: This folder contains internal packages used by the application.
- `config`: Contains configuration files and settings for the application.
- `db`: Handles the database connection and provides data access functions.
- `web`: Contains the web-related functionality, including routing and handling HTTP requests.
- `public`: This folder contains static files, such as HTML templates and CSS stylesheets.

## Technologies Used

- `Go`: The main programming language used for the project.
- `HTMX`: A JavaScript library used for dynamic HTML updates.
- `Tailwind` CSS: A utility-first CSS framework used for styling the application.
- `SQLite`: A lightweight and file-based relational database used for data storage.

## Getting Started

To run the project locally, follow these steps:

1. Clone the repository to your local machine.
2. Make sure you have Go installed.
3. Install the project dependencies by running `make gomod`.
4. Set up the database by running any necessary migrations or initializing the database.
5. Build and run the project using the command `make run`.
