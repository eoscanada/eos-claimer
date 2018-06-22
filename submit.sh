#!/bin/bash

gcloud container builds submit . --config cloudbuild.yaml --timeout 5m --substitutions COMMIT_SHA=`date +"%Y-%m-%d"`
