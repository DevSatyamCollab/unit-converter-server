# Unit Converter Project

Welcome to the Unit Converter project! This is a beginner-level web application designed to help users convert between different units of measurement.
project is tailored for backend development and is suitable for beginners looking to enhance their programming skills. It is hosted on [roadmap.sh](https://roadmap.sh/projects/unit-converter) and is designed for educational purposes.

## Overview

The Unit Converter allows users to input a value and select the units they wish to convert from and to. The application supports conversions for various types of measurements, including:

- **Length**: millimeter, centimeter, meter, kilometer, inch, foot, yard, mile
- **Weight**: milligram, gram, kilogram, ounce, pound
- **Temperature**: Celsius, Fahrenheit, Kelvin

## Project Structure

```
.
├── LICENSE
├── README.md
├── go.mod
├── main.go
└── internal/
    ├── api/
    │   ├── handler.go
    │   └── users.go
    ├── handler/
    │   └── handler.go
    ├── services/
    │   └── converter.go
    ├── static/
    │   └── style.css
    └── templates/
        ├── index.html
        ├── length.html
        ├── temperature.html
        └── weight.html
```

## Features

- User-friendly web interface for inputting values and selecting units.
- Instant conversion results displayed on the same page.
- No database required; the application processes data in real-time.

## Requirements

- Build a simple web page with sections for each type of unit conversion.
- Allow users to input values, select units for conversion, and view the results.
- The application should handle conversions without needing a backend database.

## How It Works

1. **Input Value**: Users enter the value they wish to convert.
2. **Select Units**: Users choose the units to convert from and to.
3. **Display Result**: Upon submission, the application calculates the converted value and displays it on the webpage.

The application consists of three main pages for different unit conversions: length, weight, and temperature. Each page includes a form for inputting values and selecting units.

## Getting Started

To get started with the Unit Converter project, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/DevSatyamCollab/unit-converter-server.git
```

2. Run the Server:

```bash
go run .
```

3. Access the application using one of the following options:
   - **Web Browser** — Open `index.html` directly in your browser to use the web interface.
   - **Terminal TUI Client** — Prefer staying in the terminal? You can use a separate TUI (Terminal User Interface) project that acts as a client for this server. Clone or set up the TUI project and point it at the running server to interact with the converter entirely from your terminal.

     ```bash
     git clone https://github.com/DevSatyamCollab/unit-converter-terminal-client.git
     ```

> **Note:** If you are using a TUI client, ensure the Go server is running in a separate terminal window first, as the terminal client relies on the server's API to process conversions.
