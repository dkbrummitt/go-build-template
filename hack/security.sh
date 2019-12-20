
WORKING=generated/security_scan
TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

mkdir -p $WORKING
go get -u github.com/securego/gosec/cmd/gosec

#run gosec enabling tests and JSON output
gosec -tests -fmt json -out $WORKING/$TIMESTAMP-results.json ./...


echo "-----------------------------"
echo "Security Scan Complete"
echo "-----------------------------"
echo "Review the security scan results at $WORKING/$TIMESTAMP-results.json"
echo ""
