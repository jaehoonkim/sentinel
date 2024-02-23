#!/bin/bash

apprun() {
	exec /app/morpheus-manager -config '/app/conf/morpheus-manager.yml'
}

apprun
