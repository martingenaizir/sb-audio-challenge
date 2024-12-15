# SpeakBuddy Audio-Challenge

![technology Go](https://img.shields.io/badge/technology-Go-blue.svg)
![storage MySQL](https://img.shields.io/badge/MySQL-black.svg)

Simple Backend System for Uploading, Converting, and Retrieving User Audio Recordings in a Practice Phrase Application.

[Challenge Doc](https://docs.google.com/document/d/19IV7EREMiYK6amIYYONmTUIDR3VqjYmKED0Ht45-IDo/edit?tab=t.0)

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
    - [Preloaded Data](#preloaded-data)
- [Usage](#usage)
    - [Storing Audio](#storing-audio)
    - [Retrieving Audio](#retrieving-audio)
- [Endpoints](#endpoints)
- [Database](#database)
- [Records history](#recording-history)
- [Corners cut](#corners-cut)

---

## Overview

This service allows users to upload and retrieve audio files linked to specific user IDs and phrase IDs.
The audio is stored in a `.wav` by default, and can be retrieved in different formats.

### Features:
- **POST**: Accepts an audio file, converts it to a specific file format (`.wav` by default) and stores it on the server.
    - Supported input formats: `wav`, `m4a` (or `mp4`), `mp3`
- **GET**: Gets a stored audio file by converting it to the requested audio format.
    - Accepted output formats: `wav`, `m4a` (or `mp4`), `mp3`

---

## Installation

### 1. Clone the Repository
```bash
git clone git@github.com:martingenaizir/sb-audio-challenge.git
cd sb-audio-challenge
```

### 2. Run with Docker

```bash
docker compose up -d 
```

The service will be available at `http://localhost:8080`.

### Preloaded Data
The database has **2 users**, which correspond to the IDs: `1` and `2`

- User ID `1` has a `level` of `10`, which allows to store and retrieve phrases at the same level or lower.
- User ID `2` has a `level` of `1`, which allows only to store and retrieve level `1` phrases.

The database has **2 phrases**, which correspond to the IDs: `1` and `2`

- Phrase ID `1` has a `level` of `1`.
- Phrase ID `2` has a `level` of `11`.


> Note 1
>
> Added the `levels` concept to simulate the correspondence between a phrase, its difficulty, and the user's current level.

> Note 2
>
> The DB is exposed so more data can be added if needed.
>
> (credentials in `./.env`)

---

## Usage

### Storing Audio

To upload an audio file, use the `POST` endpoint:

```bash
curl --request POST 'http://localhost:8080/audio/user/1/phrase/1' --form 'audio_file=@"./test_audio_file_1.m4a"'
```

### Retrieving Audio

To retrieve an audio file, use the `GET` endpoint:
```bash
curl --request GET 'http://localhost:8080/audio/user/1/phrase/1/m4a' -o './test_response_file.m4a'
```

---

## Endpoints

`POST /audio/user/:user_id/phrase/:phrase_id`
- **Description**: Upload an audio file for a specific user and phrase.
- **Request Body**:
    - A multipart form with the key `audio_file`, containing the audio file to upload.
- **Example**:
  ```bash
  curl --request POST 'http://localhost:8080/audio/user/1/phrase/1' --form 'audio_file=@"./audio.m4a"'
  ```


`GET /audio/user/:user_id/phrase/:phrase_id/:audio_format`
- **Description**: Retrieve a stored audio file, converting it to the specified format.
- **Parameters**:
    - `audio_format`: one of `mp3`, `m4a` (or `mp4`), `wav`.
- **Example**:
  ```bash
  curl --request GET 'http://localhost:8080/audio/user/1/phrase/1/m4a'
  ```

---

## Database
The application uses a MySQL v8.0 database consisting of three simple tables:
1. `users`: Contains valid `user_id` values.
2. `phrases`: Contains valid `phrase_id` values.
3. `user_phrases`: Associates audio files with `user_id` and `phrase_id` and stores the file paths.

[see Preloaded data](#preloaded-data)

---

## Recording history
By default, the app has recording `history` disabled. That is, both the database and the FS maintain a single phrase record per user.

The history option can be enabled from the `./.env` file, by setting `app.keep_recordings_history=true`.
> The GET endpoint remains unchanged.

## Corners cut

- Configs
    - A single file is used for environment variables, and it is not included in the `.gitignore`.
    - DB client pools use the default settings.
    - The ability to configure folders was not added to the FS.

- Validations
    - The error handler exposes date regardless the environment.
    - The storage modules (FS and DB) has weak validations, and in some cases none at all.
    - The domains have weak data validation.

- Tests
    - The application does not test most packages.

- Patterns
    - Although the `hexagonal` + `DDD` + `Facade` patterns are used, the package `services` was not vertically sliced.
- Others
  - They can be found with the comment "cutting corners"
