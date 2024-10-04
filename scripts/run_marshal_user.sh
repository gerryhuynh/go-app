CONTENT_TYPE="application/json"

if [ $# -eq 1 ]; then
    case $1 in
        json)
            CONTENT_TYPE="Content-Type: application/json"
            ;;
        proto)
            CONTENT_TYPE="Content-Type: application/protobuf"
            ;;
        *)
            echo "Invalid argument. Use 'json' or 'proto'."
            exit 1
            ;;
    esac
fi

kill $(lsof -ti:8080,50051)
go run main.go &
sleep 1
curl -H "$CONTENT_TYPE" http://localhost:8080/marshal-user
