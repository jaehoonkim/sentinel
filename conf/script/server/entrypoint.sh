#!/bin/bash

apprun() {
	exec /app/synapseee-manager -config '/app/conf/synapse-manager.yml'
}

apprun
