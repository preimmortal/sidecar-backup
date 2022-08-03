#!/bin/bash

#########################################################################
# Setup Environment
#########################################################################
echo "Setting up environment"

if [ "${PGID}" == "0" ]; then
  export USERGROUP=root
else
  export USERGROUP=backup
  addgroup -g ${PGID} ${USERGROUP}
fi

if [ "${PUID}" == "0" ]; then
  export USER=root
else
  export USER=backup
  adduser \
    --disabled-password \
    --gecos "" \
    --ingroup "${USERGROUP}" \
    --uid "${PUID}" \
    "${USER}"
fi

su ${USER} -c "/app/sidecar-backup $@"