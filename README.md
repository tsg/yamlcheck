# yamlcheck

Command line tool to check if a YAML file is valid. Based on the
[go-yaml](https://github.com/go-yaml/yaml) library.


## Usage

        $ go get github.com/tsg/yamlcheck
        $ yamlcheck file.yml
        YAML check ok

This works as long as you have `$GOPATH/bin` in your `$PATH`.

The returned status code is 0 if the check is successful and non-zero
otherwise.
