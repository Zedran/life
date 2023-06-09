build_dir=./build
exe_fname=life

version=`git rev-list --abbrev-commit -1 HEAD`

if [ ! -d "$build_dir" ]; then
    mkdir $build_dir
fi

for package in ./src ./src/world ./src/config ./src/config/theme; do
    go test $package
    
    if [ $? -ne 0 ]; then
        echo -e "\nTests failed. Aborting build.\n"
        exit 1
    fi
done

go build -trimpath -ldflags "-s -w -X main.Version=$version" -o "$build_dir/$exe_fname" ./src
