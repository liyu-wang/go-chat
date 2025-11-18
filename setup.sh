#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
    source .env
else
    echo "Error: .env file not found"
    echo "Please create a .env file with the required environment variables."
    exit 1
fi

# Validate required variables
if [ -z "$GOOGLE_CLIENT_ID" ] || [ -z "$GOOGLE_CLIENT_SECRET" ] || [ -z "$SESSION_SECRET" ]; then
    echo "Error: Missing required environment variables"
    echo "Please ensure GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, and SESSION_SECRET are set in .env"
    exit 1
fi

export GOOGLE_CLIENT_ID
export GOOGLE_CLIENT_SECRET
export SESSION_SECRET

echo "✓ Environment variables loaded successfully"