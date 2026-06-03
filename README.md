# satelliteGohelper

Rocket launching simulation powered by Electron JS and Golang

## Overview

satelliteGohelper is a desktop application that provides realistic rocket launching simulations. Built with a modern tech stack combining the power of Golang for backend computation and Electron JS for cross-platform UI, this project delivers an intuitive and visually engaging experience for space enthusiasts and developers.

Using real-world data from the ISRO GSLV Mk II launch vehicle, I developed a physics-based simulation that models the rocket's ascent from launch to escape trajectory. By applying aerospace and orbital mechanics equations, I calculated and visualized key flight parameters, including velocity, acceleration, altitude, thrust, and escape velocity, at every second of the mission.

The simulation is centered around actual GSLV Mk II vehicle specifications and demonstrates how the rocket gains speed, overcomes atmospheric drag and gravity, and eventually achieves the conditions required for spaceflight. Atmospheric data used in the calculations was derived from conditions at the Satish Dhawan Space Centre, Sriharikota, providing realistic environmental inputs for the model.

Through this project, I successfully validated and visualized the rocket's flight dynamics, showcasing the relationship between propulsion, atmospheric conditions, and orbital mechanics throughout its ascent.

## Features

- **Real-time Simulation**: Accurate physics-based rocket launch simulations
- **Cross-Platform**: Works seamlessly on Windows, macOS, and Linux
- **Interactive UI**: User-friendly interface built with Electron JS
- **High Performance**: Leverages Golang for fast computational backend
- **Trajectory Visualization**: Real-time visualization of rocket trajectories
- **Multiple Rocket Models**: Support for various satellite and rocket configurations

## Tech Stack

- **Frontend**: Electron JS - Create powerful desktop applications
- **Backend**: Golang - High-performance computation engine
- **Visualization**: WebGL-powered graphics rendering

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- Golang (v1.16 or higher)
- npm or yarn

### Installation

```bash
# Clone the repository
git clone https://github.com/ananthunairj/satelliteGohelper.git
cd satelliteGohelper

# Install dependencies
npm install

# Build the Golang backend
go build ./...
```

### Running the Application

```bash
# Start the development environment
npm start
```

## Project Structure

```
satelliteGohelper/
├── electronApp/          # Electron application frontend code
├── goService/            # Golang backend services and physics engine
├── release/              # Build and release artifacts
├── .vscode/              # VS Code workspace settings
├── LICENSE               # License file
├── README.md             # This file
├── .gitignore            # Git ignore rules
└── package.json          # Node.js dependencies
```

## Usage

1. Launch the application
2. Select a rocket/satellite model
3. Configure launch parameters (angle, thrust, etc.)
4. Click "Launch" to begin the simulation
5. Watch the real-time trajectory visualization


