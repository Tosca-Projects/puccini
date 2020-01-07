pushd "$ROOT/dist" > /dev/null
python3 -m http.server 8000 &
HTTP_SERVER_PID=$!
popd > /dev/null

sleep 1
if ! kill -0 "$HTTP_SERVER_PID" 2> /dev/null; then
	echo 'Cannot start web server'
	exit 1
fi

function the_end () {
	kill "$HTTP_SERVER_PID"
	exit $?
}

trap the_end EXIT
