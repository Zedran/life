build_dir=./build
exe_fname=life

if [ ! -d "$build_dir" ]; then
    mkdir $build_dir
fi

for package in ./src ./src/world; do
    go test -v $package
done

go build -trimpath -ldflags "-s -w" -o "$build_dir/$exe_fname" ./src
