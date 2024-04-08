#!/bin/bash

#########################################################################
# Setup Environment
#########################################################################
echo "Setting up environment"
echo "Current User: ${UID}"
if [ "${UID}" == "0" ]; then
  if [ "${PGID}" == "0" ]; then
    export USERGROUP=root
  else
    echo "Creating Group"
    export USERGROUP=backup
    addgroup -g ${PGID} ${USERGROUP}
  fi
  
  if [ "${PUID}" == "0" ]; then
    export USER=root
  else
    echo "Creating User"
    export USER=backup
    adduser \
      --disabled-password \
      --gecos "" \
      --ingroup "${USERGROUP}" \
      --uid "${PUID}" \
      "${USER}"
  fi
  su ${USER} -c "sidecar-backup $@"
else
  sidecar-backup $@
fi

