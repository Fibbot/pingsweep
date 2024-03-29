
#!/usr/bin/env bash
#https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
#slightly modified this as I was getting an array error

package="pingsweep.go"
package_split=(${package//\// })
package_name=${package_split[0]}
package_name=$(echo "$package_name" | cut -f 1 -d '.')

platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64" "linux/386")
#edit this to specify arch to use ie:
#platforms=("darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  
    env GOOS=$GOOS GOARCH=$GOARCH go build -o binaries/$output_name $package
    
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
