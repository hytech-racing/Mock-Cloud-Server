# MCAP API Documentation

This document provides information on how to use the MCAP API, including available endpoints and how to make requests.

---

## Table of Contents

- [MCAP API Documentation](#mcap-api-documentation)
  - [Table of Contents](#table-of-contents)
- [Docker Installation and Setup Guide](#docker-installation-and-setup-guide)
  - [1. Download Docker Desktop](#1-download-docker-desktop)
  - [2. Build the Docker Image](#2-build-the-docker-image)
  - [3. Running the Docker Image](#3-running-the-docker-image)
  - [4. Saving and Loading the Docker Image as a `.tar` File](#4-saving-and-loading-the-docker-image-as-a-tar-file)
      - [Step 1: Save the Docker Image:](#step-1-save-the-docker-image)
      - [Step 2: Load the Docker Image](#step-2-load-the-docker-image)
      - [Step 3: Run the Loaded Image](#step-3-run-the-loaded-image)
  - [API Endpoints](#api-endpoints)
    - [1. GET `/api/v2/mcap`](#1-get-apiv2mcap)
    - [2. POST `/api/v2/mcap/upload`](#2-post-apiv2mcapupload)

---

# Docker Installation and Setup Guide

This guide will walk you through setting up Docker to build and run a Docker image for our project.

## 1. Download Docker Desktop

First, ensure that Docker Desktop is installed on your machine. You can download it from the official Docker website:

-   [Download Docker Desktop](https://www.docker.com/products/docker-desktop/)

Docker Desktop is available for macOS, Windows, and Linux distributions.

## 2. Build the Docker Image

Before building the Docker image, ensure that you have an `.env` file containing the necessary environment variables. The `.env` file should look something like this:

```.env
AWS_REGION=your-aws-region
AWS_S3_RUN_BUCKET=your-aws-s3-run-bucket
AWS_ACCESS_KEY=your-aws-access-key
AWS_SECRET_KEY=your-aws-secret-key
```

To build the Docker image using the `.env` file, run the following command:

```bash
docker build \
  --build-arg AWS_REGION=$(grep AWS_REGION .env | cut -d '=' -f2) \
  --build-arg AWS_S3_RUN_BUCKET=$(grep AWS_S3_RUN_BUCKET .env | cut -d '=' -f2) \
  --build-arg AWS_ACCESS_KEY=$(grep AWS_ACCESS_KEY .env | cut -d '=' -f2) \
  --build-arg AWS_SECRET_KEY=$(grep AWS_SECRET_KEY .env | cut -d '=' -f2) \
  -t mockserver-hytech .
```

This command will:

-   Extract environment variables from the `.env` file.
-   Build a Docker image named `mockserver-hytech`.

## 3. Running the Docker Image

Once the image is built, you can run it directly:

```bash
docker run -p 8081:8080 mockserver-hytech
```

## 4. Saving and Loading the Docker Image as a `.tar` File

You can also save the Docker image as a `.tar` file for distribution:

#### Step 1: Save the Docker Image:

```bash
docker save -o mockserver-hytech.tar mockserver-hytech
```

#### Step 2: Load the Docker Image

To load the Docker image from the `.tar` file, use:

```bash
docker load -i mockserver-hytech.tar
```

#### Step 3: Run the Loaded Image

After loading the image, you can run it using the same `docker run` command:

```bash
docker run -p 8081:8080 mockserver-hytech
```

---

## API Endpoints

### 1. GET `/api/v2/mcap`

This endpoint allows you to retrieve signed URLs for files stored in AWS S3 buckets based on specific query parameters.

**URL:**  
http://localhost:8080/api/v2/mcap/get

**Method:**  
`GET`

**Query Parameters:**

-   `location` (string): The location related to the event (e.g., `atlanta`).
-   `date` (string): The date of the event in `MM/DD/YY` format (e.g., `9/24/24`).
-   `notes` (string): Additional notes about the event (e.g., `car ran fast`). Note that spaces should be URL-encoded as `%20`.
-   `event_type` (string): The type of event (e.g., `endurance`).

**Example Request:**

```bash
curl -X GET "http://localhost:8080/api/v2/mcap/get?location=atlanta&date=9/24/24&notes=car%20ran%20fast&event_type=endurance"
```

This returns a json of the signed URLS of the mcap files with matching parameters, sorted by date.

### 2. POST `/api/v2/mcap/upload`

**URL:**  
http://localhost:8080/api/v2/mcap/upload

**Method:**  
`POST`

**Query Parameters:**

-   `TBD` (TBD): Lorem ipsum dolor (e.g., `TBD`).

**Example Request:**

```bash
curl -X GET "http://localhost:3000/api/v2/mcap/upload"
```
