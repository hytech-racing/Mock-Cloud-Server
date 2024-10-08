# MCAP API Documentation

This document provides information on how to use the MCAP API, including available endpoints and how to make requests.

---

## Table of Contents
- [Docker Setup](#docker-setup)
- [API Endpoints](#api-endpoints)
  - [GET /api/v2/mcap](#get-apiv2mcap)
  - [POST /api/v2/mcap/upload](#post-apiv2mcapupload)
- [Error Codes](#error-codes)

---

## Docker Setup

*To be filled in later with instructions on how to set up the project using Docker.*

---

## API Endpoints

### 1. GET `/api/v2/mcap`

This endpoint allows you to retrieve signed URLs for files stored in AWS S3 buckets based on specific query parameters.

**URL:**  
http://localhost:3000/api/v2/mcap


**Method:**  
`GET`

**Query Parameters:**
- `location` (string): The location related to the event (e.g., `atlanta`).
- `date` (string): The date of the event in `MM/DD/YY` format (e.g., `9/24/24`).
- `notes` (string): Additional notes about the event (e.g., `car ran fast`). Note that spaces should be URL-encoded as `%20`.
- `event_type` (string): The type of event (e.g., `endurance`).

**Example Request:**
```bash
curl -X GET "http://localhost:3000/api/v2/mcap?location=atlanta&date=9/24/24&notes=car%20ran%20fast&event_type=endurance"
```
This returns a json of the signed URLS of the mcap files with matching parameters, sorted by date.

### 2. POST `/api/v2/mcap/upload`

This endpoint allows you to lorem ipsum dolor.

**URL:**  
http://localhost:3000/api/v2/mcap/upload


**Method:**  
`POST`

**Query Parameters:**
- `TBD` (TBD): Lorem ipsum dolor (e.g., `TBD`).

**Example Request:**
```bash
curl -X GET "http://localhost:3000/api/v2/mcap/upload"
```
This returns a lorem ipsum dolor.
