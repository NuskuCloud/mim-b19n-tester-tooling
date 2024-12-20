#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("windows/amd64" )
#platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/arm")
#platforms=("linux/arm")

rm -rf ../dist

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name='mimb19n-tester-'$GOOS'-'$GOARCH

	env GOOS=$GOOS GOARCH=$GOARCH go build -o ../dist/$output_name $package

		output_name+='.exe'

	mkdir -p ../dist

	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done