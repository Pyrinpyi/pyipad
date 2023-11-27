#!/bin/bash
rm -rf /tmp/pyipad-temp

pyipad --devnet --appdir=/tmp/pyipad-temp --profile=6061 --loglevel=debug &
pyipad_PID=$!
pyipad_KILLED=0
function killpyipadIfNotKilled() {
    if [ $pyipad_KILLED -eq 0 ]; then
      kill $pyipad_PID
    fi
}
trap "killpyipadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $pyipad_PID

wait $pyipad_PID
pyipad_KILLED=1
pyipad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "pyipad exit code: $pyipad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $pyipad_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
