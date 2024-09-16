ENDPOINT=""
N=2
URL="https://cable.ayra.ch/empty/?id=3"
# URL="https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-000.zip"

while getopts ":e:n:" opt; do
  case $opt in
    e) ENDPOINT="$OPTARG"
    ;;
    n) N="$OPTARG"
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    exit 1
    ;;
  esac
done

kill $(lsof -ti:8080) 2>/dev/null
go run main.go &
sleep 1
curl "http://localhost:8080/download$ENDPOINT?url=$URL&n=$N"
