#!/bin/bash
rm -rf /tmp/pyipad-temp

pyipad --simnet --appdir=/tmp/pyipad-temp --profile=6061 &
pyipad_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $pyipad_PID

wait $pyipad_PID
pyipad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "pyipad exit code: $pyipad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $pyipad_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
