# Ice Cream Shop - Product service

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Course Information

**Course Name:** Cloud Application Requirements and Specifications  
**Course Code:** CLCM3102    
**Student Name:** Sorawat Tanthikun

## Overview
Welcome to the Ice Cream Shop Product service! This project serves as the backend for a delightful web application that helps manage orders, inventory, and customer interactions for your favorite ice cream shop.

## Features

- **Inventory Tracking:** Keep track of available ice cream flavors.
- **API Integration:** Provides a RESTful API for communication with frontend applications.

## Table of Contents

- [Ice Cream Shop Backend](#ice-cream-shop-backend)
  - [Features](#features)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Configuration](#configuration)
    - [API Documentation](#api-documentation)
    - [License](#license)

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16+ installed
- A MySQL database

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/CLCM3102-Ice-Cream-Shop/backend-product-service.git
   
   ```
2. Go to project directory
   ```bash
   cd backend-product-service
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run proejct
   ```bash
   make run
   ```
### Configuration
- Copy the example configuration file and customize it according to your environment:

    ```bash
    cp config.example.yaml config.yaml
    ```
- Update the config.yaml file with your database and other settings.

### API Documentation
The API documentation is available [here](https://satrawo38.atlassian.net/wiki/spaces/CP/pages/4555062/API+Specification).

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

