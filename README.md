# Switch-Nexus

Switch-Nexus is an admin panel and management dashboard for managing IT infrastructure switches. The application is built using [Wails](https://wails.io) with Go and Vue.js, employing an onion architecture to ensure clean separation of concerns.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Features](#features)
3. [Architecture](#architecture)
4. [Getting Started](#getting-started)
5. [Configuration](#configuration)
6. [Usage](#usage)
7. [Contributing](#contributing)
8. [License](#license)
9. [Contact](#contact)

## Project Overview

Switch-Nexus provides an intuitive interface for IT administrators to manage and monitor switches within a network. It allows for seamless SSH connections, real-time command execution, and efficient infrastructure management.

## Features

- **SSH Connection**: Securely connect and interact with network switches.
- **Command Execution**: Run commands and scripts directly from the dashboard.
- **Real-time Monitoring**: View switch status and logs in real-time.
- **User-Friendly Interface**: Built with a responsive and intuitive design using Vue.js.
- **Logging**: Comprehensive logging to disk and standard output for troubleshooting.

## Architecture

Switch-Nexus employs an onion architecture for clean separation of concerns. The main components include:

- **Config**: Manages configuration loading and validation.
- **Middleware**: Provides logging and error-handling utilities.
- **Services**: Contains business logic for switch management.
- **Utils**: SSH connection utilities for managing sessions.
- **Interfaces**: Exposes functions to the Wails frontend.

### Folder Structure

```
Switch-Nexus
├── src
│   ├── config
│   │   └── config.go
│   ├── interfaces
│   │   └── wails.go
│   ├── middleware
│   │   ├── logging.go
│   │   └── error.go
│   ├── services
│   │   └── switch_services.go
│   └── utils
│       └── switch_connector.go
├── go.mod
├── go.sum
├── main.go
├── app.go
└── README.md
```

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.19 or higher)
- [Node.js](https://nodejs.org/) (for Vue.js frontend)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/Switch-Nexus.git
   cd Switch-Nexus
   ```

2. **Install Go dependencies:**

   ```bash
   go mod tidy
   ```

3. **Set up the Vue.js frontend:**

   ```bash
   cd frontend
   npm install
   npm run build
   cd ..
   ```

4. **Build the application:**

   ```bash
   wails build
   ```

5. **Run the application:**

   ```bash
   ./Switch-Nexus
   ```

## Configuration

The application uses a `.env` file for configuration. Create a `.env` file in the root directory and add your SSH credentials:

```env
SSH_HOST=your_ssh_host
SSH_USER=your_ssh_user
SSH_PASSWORD=your_ssh_password
```

## Usage

- **SSH Connection:** Open the application and enter your network switch credentials to connect.
- **Execute Commands:** Use the command panel to send commands to your switches.
- **View Logs:** Access real-time logs from the switch management panel.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add your feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

## License

This project is licensed under the AGPL-3.0 License - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or issues, please contact:

**Paula Selim**  
- Email: [s.paulakhalil@icloud.com](mailto:s.paulakhalil@icloud.com)
- LinkedIn: [Paula Selim](https://www.linkedin.com/in/paula-selim-572507305/)
