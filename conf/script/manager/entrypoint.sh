#!/bin/bash

apprun() {
	exec /app/sentinel-manager -config '/app/conf/sentinel-manager.yml'
}

apprun
