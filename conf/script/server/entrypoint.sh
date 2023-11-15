#!/bin/bash

apprun() {
	exec /app/sudory-manager -config '/app/conf/sudory-manager.yml'
}

apprun
