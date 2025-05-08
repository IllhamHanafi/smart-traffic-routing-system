dev:
	go mod tidy && wgo run .

debug:
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --continue --accept-multiclient .

watch:
	find . -name "*.go" | entr -r dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --continue .