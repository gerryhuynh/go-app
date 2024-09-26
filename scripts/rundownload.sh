# URL="https://cable.ayra.ch/empty/?id=3"
URL="https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-000.zip"
N=1
SEQUENTIAL=false

while getopts ":n:s" opt; do
  case $opt in
    n) N="$OPTARG"
    ;;
    s) SEQUENTIAL=true
    ;;
    :) echo "Option -$OPTARG requires an argument." >&2
       exit 1
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    exit 1
    ;;
  esac
done

kill $(lsof -ti:8080)
go run main.go &
sleep 1
curl "http://localhost:8080/download?url=$URL&n=$N&s=$SEQUENTIAL"
