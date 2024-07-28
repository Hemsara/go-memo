# Google Calendar API Integration with Go

This guide helps you integrate the Google Calendar API with a Go application. By following these steps, you will be able to authenticate with Google and access calendar events programmatically.

## Prerequisites

Before starting, ensure you have the following:

- [Go](https://golang.org/dl/) installed on your machine (version 1.16 or higher recommended)
- A [Google Cloud Platform (GCP) account](https://cloud.google.com/)
- A project created in the [GCP Console](https://console.cloud.google.com/)
- Access to the Google Calendar API enabled in your GCP project

## Setup Instructions

### 1. Create a Google Cloud Project

1. Visit the [Google Cloud Console](https://console.cloud.google.com/).
2. Click on the project drop-down menu at the top of the page and select **New Project**.
3. Enter a project name and click **Create**.

### 2. Enable the Google Calendar API

1. In the [Google API Library](https://console.cloud.google.com/apis/library), search for "Google Calendar API".
2. Click on **Google Calendar API** and then click **Enable**.

### 3. Configure OAuth Consent Screen

1. Navigate to the [OAuth consent screen](https://console.cloud.google.com/apis/credentials/consent).
2. Choose **External** as the user type and click **Create**.
3. Fill in the necessary details, such as app name and user support email, and click **Save and Continue**.

### 4. Create OAuth 2.0 Credentials

1. Go to the [Credentials page](https://console.cloud.google.com/apis/credentials).
2. Click on **Create Credentials** and select **OAuth client ID**.
3. Choose **Web application** as the application type.
4. Add authorized redirect URIs. For development, you can use `http://localhost:8080/callback`.
5. Click **Create** and save the **client ID** and **client secret**.

### 5. Install the Google API Client Library for Go

In your Go project directory, run the following command to install the Google Calendar API client:

```bash
go get google.golang.org/api/calendar/v3
