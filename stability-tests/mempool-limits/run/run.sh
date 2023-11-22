#!/bin/bash

APPDIR=/tmp/pyipad-temp
pyipad_RPC_PORT=29587

rm -rf "${APPDIR}"

pyipad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${pyipad_RPC_PORT}" --profile=6061 &
pyipad_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${pyipad_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $pyipad_PID

wait $pyipad_PID
pyipad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "pyipad exit code: $pyipad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $pyipad_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
