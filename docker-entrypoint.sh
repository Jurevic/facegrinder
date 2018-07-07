#!/bin/bash

echo "Applying database migrations"
facegrinder migrate

echo "Starting server"
facegrinder serve
