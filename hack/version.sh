VERSION=`git describe --tags --always`

if [ $# -eq 0 ]
then
    echo "Version not supplied using $VERSION"
else
    VERSION=$1
    export GO111MODULE=on
    go get -u ./...
fi

RELEASE_DATE=`date +%Y-%m-%d,%H:%M:%S`
GoVersion=`go version |awk '{print $3}'`

sed -i.bak "/var VERSION/ s/=.*/= \"$VERSION\"/" ./pkg/version/version.go
sed -i.bak "/sonar.projectVersion/ s/=.*/=$VERSION/" sonar-project.properties
sed -i.bak "/var RELEASE_DATE/ s/=.*/= \"$RELEASE_DATE\"/" ./pkg/version/version.go
sed -i.bak "/var GoVersion/ s/=.*/= \"$GoVersion\"/" ./pkg/version/version.go

if [ $# -eq 1 ]
then
    git commit -am "Update dependendcies and bump Version"
    git tag $1

    echo ""
    echo "Tagged, bagged, and ready to push"
fi


echo version is $VERSION, $RELEASE_DATE
