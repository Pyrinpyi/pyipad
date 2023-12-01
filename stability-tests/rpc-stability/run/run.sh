#!/bin/bash
rm -rf /tmp/pyipad-temp

pyipad --devnet --appdir=/tmp/pyipad-temp --profile=6061 --loglevel=debug &
pyipad_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $pyipad_PID

wait $pyipad_PID
pyipad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "pyipad exit code: $pyipad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $pyipad_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
