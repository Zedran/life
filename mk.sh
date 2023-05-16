build_dir=./build
exe_fname=life

if [ ! -d "$build_dir" ]; then
    mkdir $build_dir
fi

failed=0

for package in ./src ./src/world ./src/config ./src/config/theme; do
    go test $package
    
    if [ $? -ne 0 ]; then
        failed=1
    fi
done

if [ $failed -eq 1 ]; then
    echo -e "\nTests failed. Aborting build.\n"
    exit 1
fi

go build -trimpath -ldflags "-s -w" -o "$build_dir/$exe_fname" ./src
